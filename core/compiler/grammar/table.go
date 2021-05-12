package grammar

import (
	"errors"
	"fmt"
	"strings"
)

const (
	TableClosureMatchExist     = -1
	TableClosureMatchReplace   = 1
	TableClosureMatchUnrelated = 0
)

const (
	TableVirtualGlobalType = -1
)

type Table struct {
	rules           *RuleSet
	global          *Nonterminal
	end             *Terminal
	start           int
	actions         map[int]*TableActionGroup // map[status]map[symbol]action
	gotos           map[int]*TableActionGroup // map[status]map[symbol]goto
	nonterminalKeys *SymbolSet
	terminalKeys    *SymbolSet
	toActions       map[int][]*TableAction
	closures        map[int]*TableClosure
	firsts          map[Symbol]*SymbolSet
	startProjects   map[int][]*TableProject
	counter         *Count
}

func NewTable() *Table {
	return &Table{
		rules:           NewRuleSet(),
		actions:         map[int]*TableActionGroup{},
		gotos:           map[int]*TableActionGroup{},
		nonterminalKeys: NewSymbolSet(),
		terminalKeys:    NewSymbolSet(),
		closures:        map[int]*TableClosure{},
		toActions:       map[int][]*TableAction{},
		firsts:          map[Symbol]*SymbolSet{},
		startProjects:   map[int][]*TableProject{},
		counter:         NewCount(0),
	}
}

func (g *Table) GetFirsts() map[Symbol]*SymbolSet {
	return g.firsts
}

func (g *Table) GetClosures() map[int]*TableClosure {
	return g.closures
}

func (g *Table) GetNonterminalKeys() *SymbolSet {
	return g.nonterminalKeys
}

func (g *Table) GetTerminalKeys() *SymbolSet {
	return g.terminalKeys
}

func (g *Table) GetAction(status int, symbol int) *TableAction {
	return g.actions[status].GetAction(symbol)
}

func (g *Table) GetGoto(status int, symbol int) *TableAction {
	return g.gotos[status].GetAction(symbol)
}

func (g *Table) GetStart() int {
	return g.start
}

func (g *Table) SetRules(rules []*Rule) {
	g.rules.Add(rules...)
}

func (g *Table) SetGlobal(global *Nonterminal) {
	g.global = global
}

func (g *Table) SetEnd(end *Terminal) {
	g.end = end
}

func (g *Table) Build() error {
	err := g.init()
	if err != nil {
		return err
	}
	g.makeFirsts()
	g.makeProjects()
	g.makeClosures()
	return nil
}

func (g *Table) makeClosures() {
	g.start = g.makeClosureStep(map[*TableProject]*SymbolSet{
		g.startProjects[g.global.Type()][0]: NewSymbolSet(g.end),
	}).Id()
}

func (g *Table) makeClosureStep(cursors map[*TableProject]*SymbolSet) *TableClosure {
	closure := NewTableClosure(g.counter.Next())
	for cursor, lookaheads := range cursors {
		closure.AddProject(cursor, lookaheads)
	}
	g.equivalenceClosure(cursors, closure)

	oldClosure, direction := g.matchClosure(closure)
	if direction == TableClosureMatchExist {
		//exist
		return oldClosure
	}
	if direction == TableClosureMatchReplace {
		//replace
		g.moveClosure(oldClosure, closure)
	}

	g.closures[closure.Id()] = closure
	g.toActions[closure.Id()] = []*TableAction{}

	g.actions[closure.Id()] = NewTableActionGroup()
	g.gotos[closure.Id()] = NewTableActionGroup()

	endProjects := closure.GetProjectsByNextChild(nil)
	for endProject, lookaheads := range endProjects {
		lookaheads.Iterate(func(lookahead Symbol) bool {
			if endProject.Rule.GetResult() == g.global {
				//accept
				g.actions[closure.Id()].SetAction(lookahead.Type(), NewTableActionAccept())
			} else {
				// polymerize
				g.actions[closure.Id()].SetAction(lookahead.Type(), NewTableActionPolymerize(endProject.Rule))
			}
			return false
		})
	}

	closure.NextChildren().Iterate(func(child Symbol) bool {
		projects := closure.GetProjectsByNextChild(child)
		nextProjects := g.nextProjects(projects)
		nextClosure := g.makeClosureStep(nextProjects)
		if child.SymbolType() == SymbolTypeTerminal {
			//move
			action := NewTableActionMove(nextClosure.Id())
			g.actions[closure.Id()].SetAction(child.Type(), action)
			g.toActions[nextClosure.Id()] = append(g.toActions[nextClosure.Id()], action)
		} else {
			// goto
			action := NewTableActionGoto(nextClosure.Id())
			g.gotos[closure.Id()].SetAction(child.Type(), action)
			g.toActions[nextClosure.Id()] = append(g.toActions[nextClosure.Id()], action)
		}
		return false
	})
	return closure
}

func (g *Table) moveClosure(from, to *TableClosure) {
	for _, action := range g.toActions[from.Id()] {
		action.SetStatus(to.Id())
	}
	delete(g.closures, from.Id())
	delete(g.actions, from.Id())
	delete(g.gotos, from.Id())
	delete(g.toActions, from.Id())
}

func (g *Table) matchClosure(target *TableClosure) (*TableClosure, int) {
	for _, closure := range g.closures {
		if closure.Include(target) {
			return closure, TableClosureMatchExist
		}
		if target.Include(closure) {
			return closure, TableClosureMatchReplace
		}
	}
	return nil, TableClosureMatchUnrelated
}

func (g *Table) nextProjects(nows map[*TableProject]*SymbolSet) map[*TableProject]*SymbolSet {
	nexts := map[*TableProject]*SymbolSet{}
	for project, lookaheads := range nows {
		nexts[project.Next] = lookaheads
	}
	return nexts
}

func (g *Table) equivalenceClosure(cursors map[*TableProject]*SymbolSet, closure *TableClosure) {
	for cursor, lookaheads := range cursors {
		symbol := cursor.GetNextChild()
		if symbol == nil || symbol.SymbolType() == SymbolTypeTerminal {
			continue
		}
		equivalences := g.startProjects[symbol.Type()]
		for _, equivalence := range equivalences {
			equivalenceLookaheads := lookaheads
			if cursor.Next != nil && !cursor.Next.IsEnd() {
				next := cursor.Next.GetNextChild()
				if next.SymbolType() == SymbolTypeTerminal {
					equivalenceLookaheads = NewSymbolSet(next)

				} else {
					equivalenceLookaheads = g.firsts[next]
				}
			}
			successLookaheads := closure.AddProject(equivalence, equivalenceLookaheads)
			if successLookaheads.Size() > 0 {
				// add success
				g.equivalenceClosure(
					map[*TableProject]*SymbolSet{
						equivalence: successLookaheads,
					},
					closure,
				)
			}
			// else project[lookaheads] exist
		}
	}
}

func (g *Table) makeProjects() {
	g.rules.Iterate(func(rule *Rule) bool {
		startProject := NewTableProject(rule, 0)
		g.startProjects[rule.GetResult().Type()] = append(g.startProjects[rule.GetResult().Type()], startProject)
		last := startProject
		for childIndex := 1; childIndex <= rule.Size(); childIndex++ {
			project := NewTableProject(rule, childIndex)
			last.Next = project
			last = project
		}
		return false
	})
}

func (g *Table) makeFirsts() {
	for next := true; next; {
		next = false
		g.rules.Iterate(func(rule *Rule) bool {
			result := rule.GetResult()
			if g.firsts[result] == nil {
				g.firsts[result] = NewSymbolSet()
			}
			if rule.Size() == 0 {
				return false
			}
			child := rule.GetChild(0)
			if child.SymbolType() == SymbolTypeTerminal {
				if !g.firsts[result].Has(child) {
					g.firsts[result].Add(child)
					next = true
				}
			} else {
				if g.firsts[child] == nil {
					return false
				}
				g.firsts[child].Iterate(func(first Symbol) bool {
					if !g.firsts[result].Has(first) {
						g.firsts[result].Add(first)
						next = true
					}
					return false
				})
			}
			return false
		})
	}
}

func (g *Table) augmentGlobal() {
	virtualGlobal := NewNonterminal(TableVirtualGlobalType, fmt.Sprintf("-%v", g.global.Name()))
	virtualRule := NewRule(virtualGlobal, g.global)
	g.rules.Add(virtualRule)
	g.global = virtualGlobal
}

func (g *Table) init() error {
	g.initKeys()
	err := g.checkKeys()
	if err != nil {
		return err
	}
	err = g.checkGlobal()
	if err != nil {
		return err
	}
	g.augmentGlobal()
	return nil
}

func (g *Table) checkKeys() error {
	var err error
	if g.nonterminalKeys.Iterate(func(key Symbol) bool {
		if key.Type() < 0 {
			err = errors.New(fmt.Sprintf("Illegal nonterminal type, must be greater than 0. nonterminal: %v, type : %v", key.Name(), key.Type()))
			return true
		}
		if !g.rules.HasByResult(key) {
			err = errors.New(fmt.Sprintf("No rule can polymerize nonterminal: %v", key.Name()))
			return true
		}
		return false
	}) {
		return err
	}
	if g.terminalKeys.Iterate(func(key Symbol) bool {
		if key.Type() < 0 {
			err = errors.New(fmt.Sprintf("Illegal terminal type, must be greater than 0. terminal: %v, type : %v", key.Name(), key.Type()))
			return true
		}
		return false
	}) {
		return err
	}
	return nil
}

func (g *Table) initKeys() {
	g.rules.Iterate(func(rule *Rule) bool {
		g.nonterminalKeys.Add(rule.GetResult())
		for index := 0; index < rule.Size(); index++ {
			symbol := rule.GetChild(index)
			if symbol.SymbolType() == SymbolTypeNonterminal {
				g.nonterminalKeys.Add(symbol)
			} else {
				g.terminalKeys.Add(symbol)
			}
		}
		return false
	})
	g.terminalKeys.Add(g.end)
}

func (g *Table) checkGlobal() error {
	if g.global == nil {
		return errors.New("Global missed.")
	}
	if !g.rules.HasByResult(g.global) {
		return errors.New("Rule missed which result to global.")
	}
	return nil
}

func (g *Table) ToString() string {
	nonterminals := []Symbol{}
	g.GetNonterminalKeys().Iterate(func(key Symbol) bool {
		nonterminals = append(nonterminals, key)
		return false
	})
	terminals := []Symbol{}
	g.GetTerminalKeys().Iterate(func(key Symbol) bool {
		terminals = append(terminals, key)
		return false
	})
	values := []string{}
	titles := []string{}
	brs := []string{}
	titles = append(titles, "status")
	brs = append(brs, ":--:")
	for _, terminal := range terminals {
		titles = append(titles, terminal.Name())
		brs = append(brs, ":--:")
	}
	for _, nonterminal := range nonterminals {
		titles = append(titles, nonterminal.Name())
		brs = append(brs, ":--:")
	}

	values = append(values, strings.Join(titles, "|"))
	values = append(values, strings.Join(brs, "|"))

	for status, _ := range g.closures {
		value := []string{}
		value = append(value, fmt.Sprintf("%v", status))
		isPolymerize := false
		for _, terminal := range terminals {
			action := g.GetAction(status, terminal.Type())
			value = append(value, action.ToString())
			if action != nil && action.Type() == ActionPolymerizeType {
				isPolymerize = true
				break
			}
		}
		if !isPolymerize {
			for _, nonterminal := range nonterminals {
				value = append(value, g.GetGoto(status, nonterminal.Type()).ToString())
			}
		}
		values = append(values, strings.Join(value, "|"))
	}
	return strings.Join(values, "\n")
}
