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
	eof             *Terminal
	start           int
	actions         map[int]*TableActionGroup // map[status]map[symbol]action
	gotos           map[int]*TableActionGroup // map[status]map[symbol]goto
	nonterminalKeys *SymbolSet
	terminalKeys    *SymbolSet
	toActions       map[int][]*TableAction
	closures        map[int]*TableClosure
	firsts          map[Symbol]*SymbolSet
	startProjects   map[int][]*TableProject
	moves           map[int]int
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
		moves:           map[int]int{},
		counter:         NewCount(1),
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

func (g *Table) GetExpect(status int) []Symbol {
	expectations := []Symbol{}
	g.terminalKeys.Iterate(func(key Symbol) bool {
		if g.GetAction(status, key.Type()) != nil {
			expectations = append(expectations, key)
		}
		return false
	})
	return expectations
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

func (g *Table) SetEof(eof *Terminal) {
	g.eof = eof
}

func (g *Table) Build() error {
	err := g.init()
	if err != nil {
		return err
	}
	g.makeFirsts()
	g.makeProjects()
	return g.makeClosures()
}

func (g *Table) Clear() {
	g.rules = nil
	g.global = nil
	g.eof = nil
	g.toActions = nil
	g.closures = nil
	g.firsts = nil
	g.startProjects = nil
	g.moves = nil
	g.counter = nil
}

func (g *Table) makeClosures() error {
	closure, err := g.makeClosureStep(map[*TableProject]*SymbolSet{
		g.startProjects[g.global.Type()][0]: NewSymbolSet(g.eof),
	})
	if err != nil {
		return err
	}
	g.start = closure.Id()
	return nil
}

func (g *Table) makeClosureStep(cursors map[*TableProject]*SymbolSet) (closure *TableClosure, err error) {
	closure = NewTableClosure(g.counter.Next())
	for cursor, lookaheads := range cursors {
		closure.AddProject(cursor, lookaheads)
	}
	g.equivalenceClosure(cursors, closure)

	oldClosures, direction := g.matchClosure(closure)
	if direction == TableClosureMatchExist {
		//exist
		closure = oldClosures[0]
		return
	}
	g.closures[closure.Id()] = closure
	g.toActions[closure.Id()] = []*TableAction{}
	g.actions[closure.Id()] = NewTableActionGroup()
	g.gotos[closure.Id()] = NewTableActionGroup()
	if direction == TableClosureMatchReplace {
		//replace
		g.moveClosure(oldClosures, closure)
	}

	for endProject, lookaheads := range closure.GetProjectsByNextChild(nil) {
		if lookaheads.Iterate(func(lookahead Symbol) bool {
			if endProject.Rule.GetResult() == g.global {
				// accept
				accept := NewTableActionAccept(map[*TableProject]*SymbolSet{
					endProject: NewSymbolSet(lookahead),
				})
				exist := g.actions[closure.Id()].GetAction(lookahead.Type())
				if exist != nil {
					err = g.actionConfictError(exist, accept, closure, lookahead)
					return true
				}
				g.actions[closure.Id()].SetAction(lookahead.Type(), accept)
			} else {
				// polymerize
				polymerize := NewTableActionPolymerize(endProject.Rule, map[*TableProject]*SymbolSet{
					endProject: NewSymbolSet(lookahead),
				})
				exist := g.actions[closure.Id()].GetAction(lookahead.Type())
				if exist != nil {
					err = g.actionConfictError(exist, polymerize, closure, lookahead)
					return true
				}
				g.actions[closure.Id()].SetAction(lookahead.Type(), polymerize)
			}
			return false
		}) {
			return
		}
	}

	closure.NextChildren().Iterate(func(child Symbol) bool {
		var nextClosure *TableClosure
		projects := closure.GetProjectsByNextChild(child)
		nextClosure, err = g.makeClosureStep(g.nextProjects(projects))
		if err != nil {
			return true
		}
		if g.actions[closure.Id()] == nil {
			return true
		}
		nextId := g.getClosureIdMoved(nextClosure.Id())
		if child.SymbolType() == SymbolTypeTerminal {
			//move
			move := NewTableActionMove(nextId, projects)
			exist := g.actions[closure.Id()].GetAction(child.Type())
			if exist != nil {
				err = g.actionConfictError(exist, move, closure, child)
				return true
			}
			g.actions[closure.Id()].SetAction(child.Type(), move)
			g.toActions[nextId] = append(g.toActions[nextId], move)
		} else {
			// goto
			action := NewTableActionGoto(nextId, projects)
			exist := g.gotos[closure.Id()].GetAction(child.Type())
			if exist != nil {
				err = g.actionConfictError(exist, action, closure, child)
				return true
			}
			g.gotos[closure.Id()].SetAction(child.Type(), action)
			g.toActions[nextId] = append(g.toActions[nextId], action)
		}
		return false
	})
	return
}

func (g *Table) getClosureIdMoved(id int) int {
	for {
		to := g.moves[id]
		if to == 0 {
			return id
		}
		id = to
	}
}

func (g *Table) actionConfictError(left, right *TableAction, closure *TableClosure, lookahead Symbol) error {
	leftNames, rightNames := []string{}, []string{}
	for project, _ := range left.Projects() {
		leftNames = append(leftNames, project.ToString())
	}
	for project, _ := range right.Projects() {
		rightNames = append(rightNames, project.ToString())
	}
	return fmt.Errorf("Rule conflict between (%v) and (%v), status: %v, lookahead: %v.", strings.Join(leftNames, " , "), strings.Join(rightNames, " , "), closure.Id(), lookahead.Name())
}

func (g *Table) moveClosure(froms []*TableClosure, to *TableClosure) {
	for _, from := range froms {
		g.moves[from.Id()] = to.Id()
		for _, action := range g.toActions[from.Id()] {
			action.SetStatus(to.Id())
			g.toActions[to.Id()] = append(g.toActions[to.Id()], action)
		}
		delete(g.closures, from.Id())
		delete(g.actions, from.Id())
		delete(g.gotos, from.Id())
		delete(g.toActions, from.Id())
	}
}

func (g *Table) matchClosure(target *TableClosure) ([]*TableClosure, int) {
	for _, closure := range g.closures {
		if closure.Include(target) {
			return []*TableClosure{closure}, TableClosureMatchExist
		}
	}
	replaces := []*TableClosure{}
	for _, closure := range g.closures {
		if target.Include(closure) {
			replaces = append(replaces, closure)
		}
	}
	if len(replaces) > 0 {
		return replaces, TableClosureMatchReplace
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
			err = fmt.Errorf("Illegal nonterminal type, must be greater than 0. nonterminal: %v, type : %v", key.Name(), key.Type())
			return true
		}
		if !g.rules.HasByResult(key) {
			err = fmt.Errorf("No rule can polymerize nonterminal: %v", key.Name())
			return true
		}
		return false
	}) {
		return err
	}
	if g.terminalKeys.Iterate(func(key Symbol) bool {
		if key.Type() < 0 {
			err = fmt.Errorf("Illegal terminal type, must be greater than 0. terminal: %v, type : %v", key.Name(), key.Type())
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
	g.terminalKeys.Add(g.eof)
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
	values := []string{}
	titles := []string{}
	brs := []string{}
	titles = append(titles, "status")
	brs = append(brs, ":--:")
	terminals := []Symbol{}
	nonterminals := []Symbol{}
	g.GetTerminalKeys().Iterate(func(key Symbol) bool {
		terminals = append(terminals, key)
		titles = append(titles, key.Name())
		brs = append(brs, ":--:")
		return false
	})
	g.GetNonterminalKeys().Iterate(func(key Symbol) bool {
		nonterminals = append(nonterminals, key)
		titles = append(titles, key.Name())
		brs = append(brs, ":--:")
		return false
	})
	values = append(values, strings.Join(titles, "|"), strings.Join(brs, "|"))
	for status, _ := range g.actions {
		value := []string{}
		value = append(value, fmt.Sprintf("%v", status))
		for _, terminal := range terminals {
			value = append(value, g.GetAction(status, terminal.Type()).ToString())
		}
		for _, nonterminal := range nonterminals {
			value = append(value, g.GetGoto(status, nonterminal.Type()).ToString())
		}
		values = append(values, strings.Join(value, "|"))
	}
	return strings.Join(values, "\n")
}
