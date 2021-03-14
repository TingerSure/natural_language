package grammar

import (
	"errors"
	"fmt"
	"strings"
)

type Table struct {
	rules           []*Rule
	global          *Nonterminal
	accept          *Terminal
	start           int
	actions         map[int]*TableActionGroup // map[status]map[symbol]action
	gotos           map[int]*TableActionGroup // map[status]map[symbol]goto
	nonterminalKeys map[Symbol]bool
	terminalKeys    map[Symbol]bool
	statusKeys      []int
}

func NewTable() *Table {
	return &Table{
		actions:         map[int]*TableActionGroup{},
		gotos:           map[int]*TableActionGroup{},
		nonterminalKeys: map[Symbol]bool{},
		terminalKeys:    map[Symbol]bool{},
	}
}

func (g *Table) GetStatusKeys() []int {
	return g.statusKeys
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

func (g *Table) SetAccept(accept *Terminal) {
	g.accept = accept
}

func (g *Table) Build() error {
	err := g.check()
	if err != nil {
		return err
	}
	g.augmentGlobal()
	g.makeClosures(g.makeProjects(g.rules))
	return nil
}

func (g *Table) makeClosures(startProjects map[int][]*TableProject) {
	g.start = g.makeClosureStep(
		startProjects[g.global.Type()][0],
		NewCount(0),
		startProjects,
		map[*TableProject]*TableClosure{},
	).Id()
}

func (g *Table) makeClosureStep(
	cursor *TableProject,
	counter *Count,
	startProjects map[int][]*TableProject,
	closureMaps map[*TableProject]*TableClosure,
) *TableClosure {
	closure := closureMaps[cursor]
	if closure != nil {
		// existed
		return closure
	}
	closure = NewTableClosure(counter.Next())
	closure.AddProject(cursor)
	closureMaps[cursor] = closure
	statusIndex := closure.Id()
	g.statusKeys = append(g.statusKeys, statusIndex)
	if cursor.Index == cursor.Rule.Size() && cursor.Rule.GetResult() != g.global {
		// polymerize
		group := NewTableActionGroupPolymerize(cursor.Rule)
		g.actions[statusIndex] = group
		g.gotos[statusIndex] = group
		return closure
	}
	g.actions[statusIndex] = NewTableActionGroup()
	g.gotos[statusIndex] = NewTableActionGroup()
	if cursor.Index == cursor.Rule.Size() && cursor.Rule.GetResult() == g.global {
		// accept
		g.actions[statusIndex].SetAction(g.accept.Type(), NewTableActionAccept())
		return closure
	}
	g.equivalenceClosure(cursor, closure, startProjects)
	for project, _ := range closure.GetProjects() {
		next := project.Next
		nextClosure := g.makeClosureStep(next, counter, startProjects, closureMaps)
		symbol := project.GetNextChild()
		if symbol.SymbolType() == SymbolTypeTerminal {
			// move
			g.actions[statusIndex].SetAction(symbol.Type(), NewTableActionMove(nextClosure.Id()))
		} else {
			// goto
			g.gotos[statusIndex].SetAction(symbol.Type(), NewTableActionGoto(nextClosure.Id()))
		}
	}
	return closure
}

func (g *Table) equivalenceClosure(
	cursor *TableProject,
	closure *TableClosure,
	startProjects map[int][]*TableProject,
) {
	symbol := cursor.GetNextChild()
	if symbol.SymbolType() == SymbolTypeTerminal {
		return
	}
	equivalences := startProjects[symbol.Type()]
	for _, equivalence := range equivalences {
		if closure.AddProject(equivalence) {
			// add success
			g.equivalenceClosure(equivalence, closure, startProjects)
		}
		// project exist
	}
}

func (g *Table) makeProjects(rules []*Rule) map[int][]*TableProject {
	startProjects := map[int][]*TableProject{}
	for _, rule := range rules {
		startProject := NewTableProject(rule, 0)
		startProjects[rule.GetResult().Type()] = append(startProjects[rule.GetResult().Type()], startProject)
		last := startProject
		for childIndex := 1; childIndex <= rule.Size(); childIndex++ {
			project := NewTableProject(rule, childIndex)
			last.Next = project
			last = project
		}
	}
	return startProjects
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
		titles = append(titles, fmt.Sprintf("%v", terminal.Name()))
		brs = append(brs, ":--:")
	}
	for _, nonterminal := range nonterminals {
		titles = append(titles, fmt.Sprintf("%v", nonterminal.Name()))
		brs = append(brs, ":--:")
	}

	values = append(values, strings.Join(titles, "|"))
	values = append(values, strings.Join(brs, "|"))

	for _, status := range g.GetStatusKeys() {
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
