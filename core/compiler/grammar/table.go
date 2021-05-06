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

type Table struct {
	rules           []*Rule
	global          *Nonterminal
	end             *Terminal
	start           int
	actions         map[int]*TableActionGroup // map[status]map[symbol]action
	gotos           map[int]*TableActionGroup // map[status]map[symbol]goto
	nonterminalKeys map[Symbol]bool
	terminalKeys    map[Symbol]bool
	toActions       map[int][]*TableAction
	closures        map[int]*TableClosure
	firsts          map[Symbol]map[Symbol]bool
	startProjects   map[int][]*TableProject
	counter         *Count
}

func NewTable() *Table {
	return &Table{
		actions:         map[int]*TableActionGroup{},
		gotos:           map[int]*TableActionGroup{},
		nonterminalKeys: map[Symbol]bool{},
		terminalKeys:    map[Symbol]bool{},
		closures:        map[int]*TableClosure{},
		toActions:       map[int][]*TableAction{},
		firsts:          map[Symbol]map[Symbol]bool{},
		startProjects:   map[int][]*TableProject{},
		counter:         NewCount(0),
	}
}

func (g *Table) GetFirsts() map[Symbol]map[Symbol]bool {
	return g.firsts
}

func (g *Table) GetClosures() map[int]*TableClosure {
	return g.closures
}

func (g *Table) GetNonterminalKeys() map[Symbol]bool {
	return g.nonterminalKeys
}

func (g *Table) GetTerminalKeys() map[Symbol]bool {
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
	g.rules = rules
}

func (g *Table) SetGlobal(global *Nonterminal) {
	g.global = global
}

func (g *Table) SetEnd(end *Terminal) {
	g.end = end
}

func (g *Table) Build() error {
	err := g.check()
	if err != nil {
		return err
	}
	g.augmentGlobal()
	g.makeFirsts()
	g.makeProjects()
	g.makeClosures()
	return nil
}

func (g *Table) makeClosures() {
	g.start = g.makeClosureStep(map[*TableProject]map[Symbol]bool{
		g.startProjects[g.global.Type()][0]: map[Symbol]bool{
			g.end: true,
		},
	}).Id()
}

func (g *Table) makeClosureStep(cursors map[*TableProject]map[Symbol]bool) *TableClosure {
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
		for lookahead, _ := range lookaheads {
			if endProject.Rule.GetResult() == g.global {
				//accept
				g.actions[closure.Id()].SetAction(lookahead.Type(), NewTableActionAccept())
			} else {
				// polymerize
				g.actions[closure.Id()].SetAction(lookahead.Type(), NewTableActionPolymerize(endProject.Rule))
			}
		}
	}

	nextChildren := closure.NextChildren()
	for child, _ := range nextChildren {
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
	}
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

func (g *Table) nextProjects(nows map[*TableProject]map[Symbol]bool) map[*TableProject]map[Symbol]bool {
	nexts := map[*TableProject]map[Symbol]bool{}
	for project, lookaheads := range nows {
		nexts[project.Next] = lookaheads
	}
	return nexts
}

func (g *Table) equivalenceClosure(cursors map[*TableProject]map[Symbol]bool, closure *TableClosure) {
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
					equivalenceLookaheads = map[Symbol]bool{
						next: true,
					}
				} else {
					equivalenceLookaheads = g.firsts[next]
				}
			}
			successLookaheads := closure.AddProject(equivalence, equivalenceLookaheads)
			if len(successLookaheads) > 0 {
				// add success
				g.equivalenceClosure(
					map[*TableProject]map[Symbol]bool{
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
	for _, rule := range g.rules {
		startProject := NewTableProject(rule, 0)
		g.startProjects[rule.GetResult().Type()] = append(g.startProjects[rule.GetResult().Type()], startProject)
		last := startProject
		for childIndex := 1; childIndex <= rule.Size(); childIndex++ {
			project := NewTableProject(rule, childIndex)
			last.Next = project
			last = project
		}
	}
}

func (g *Table) makeFirsts() {
	for next := true; next; {
		next = false
		for _, rule := range g.rules {
			result := rule.GetResult()
			if g.firsts[result] == nil {
				g.firsts[result] = map[Symbol]bool{}
			}
			child := rule.GetChild(0)
			if child.SymbolType() == SymbolTypeTerminal {
				if !g.firsts[result][child] {
					g.firsts[result][child] = true
					next = true
				}
			} else {
				if g.firsts[child] == nil {
					continue
				}
				for first, _ := range g.firsts[child] {
					if !g.firsts[result][first] {
						g.firsts[result][first] = true
						next = true
					}
				}
			}
		}
	}
}

func (g *Table) augmentGlobal() {
	virtualGlobal := NewNonterminal(-1, fmt.Sprintf("-%v", g.global.Name()))
	virtualRule := NewRule(virtualGlobal, g.global)
	g.rules = append(g.rules, virtualRule)
	g.global = virtualGlobal
}

func (g *Table) check() error {
	err := g.checkGlobal()
	if err != nil {
		return err
	}
	g.checkRules()
	// TODO check more
	return nil
}

func (g *Table) checkRules() {
	for _, rule := range g.rules {
		g.nonterminalKeys[rule.GetResult()] = true
		for index := 0; index < rule.Size(); index++ {
			symbol := rule.GetChild(index)
			if symbol.SymbolType() == SymbolTypeNonterminal {
				g.nonterminalKeys[symbol] = true
			} else {
				g.terminalKeys[symbol] = true
			}
		}
	}
	g.terminalKeys[g.end] = true
}

func (g *Table) checkGlobal() error {
	if g.global == nil {
		return errors.New("Global missed.")
	}
	resultCount := 0
	childCount := 0
	for _, rule := range g.rules {
		if g.global.Equal(rule.GetResult()) {
			resultCount++
		}
		for index := 0; index < rule.Size(); index++ {
			if g.global.Equal(rule.GetChild(index)) {
				childCount++
			}
		}
	}
	if resultCount < 1 {
		return errors.New("Rule missed which result to global")
	}
	return nil
}

func (g *Table) ToString() string {
	nonterminals := []Symbol{}
	for key, _ := range g.GetNonterminalKeys() {
		nonterminals = append(nonterminals, key)
	}
	terminals := []Symbol{}
	for key, _ := range g.GetTerminalKeys() {
		terminals = append(terminals, key)
	}
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
