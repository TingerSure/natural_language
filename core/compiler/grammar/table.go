package grammar

import (
	"errors"
)

type Table struct {
	typesCount int
	rules      []*Rule
	global     *Nonterminal
	accept     *Terminal
	actions    map[int]*TableActionGroup // map[status]map[symbol]action
	gotos      map[int]*TableActionGroup // map[status]map[symbol]goto
}

func NewTable() *Table {
	return &Table{
		typesCount: 0,
		actions:    map[int]*TableActionGroup{},
		gotos:      map[int]*TableActionGroup{},
	}
}

func (g *Table) NextTypes() int {
	g.typesCount++
	return g.typesCount
}

func (g *Table) AddRule(rule *Rule) {
	g.rules = append(g.rules, rule)
}

func (g *Table) SetGlobal(global *Nonterminal) {
	g.global = global
}

func (g *Table) Build() error {
	err := g.format()
	if err != nil {
		return err
	}
	g.makeClosures(g.makeProjects(g.rules))
	return nil
}

func (g *Table) makeClosures(
	startProjects map[int][]*TableProject,
) {
	acceptClouse := g.makeClosureStep(
		startProjects[g.global.Type()][0],
		NewCount(0),
		startProjects,
		[]*TableClosure{},
		map[*TableProject]*TableClosure{},
	)
	// accept
	g.actions[acceptClouse.Id()].SetAction(g.accept.Type(), NewTableActionAccept())
}

func (g *Table) makeClosureStep(
	cursor *TableProject,
	counter *Count,
	startProjects map[int][]*TableProject,
	closures []*TableClosure,
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
	if cursor.Index == cursor.Rule.Size() {
		// polymerize
		group := NewTableActionPolymerize(cursor.Rule)
		g.actions[statusIndex] = group
		g.gotos[statusIndex] = group
		return closure
	}
	group := NewTableActionGroup()
	g.actions[statusIndex] = group
	g.gotos[statusIndex] = group
	g.equivalenceClosure(cursor, closure, startProjects)
	for project, _ := range closure.GetProjects() {
		next := project.Next
		nextClosure := g.makeClosureStep(next, counter, startProjects, closures, closureMaps)
		symbol := project.GetNextChild()
		if symbol.SymbolType() == SymbolTypeTerminal {
			// move
			g.actions[statusIndex].SetAction(symbol.Type(), NewTableActionMove(nextClosure.Id()))
		}
		// goto
		g.gotos[statusIndex].SetAction(symbol.Type(), NewTableActionGoto(nextClosure.Id()))
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

func (g *Table) format() error {
	err := g.formatGlobal()
	if err != nil {
		return err
	}
	// TODO format more
	return nil
}

func (g *Table) formatGlobal() error {
	if g.global == nil {
		return errors.New("Global missed.")
	}
	resultCount := 0
	fromCount := 0
	for _, rule := range g.rules {
		if g.global.Equal(rule.GetResult()) {
			resultCount++
		}
		for index := 0; index < rule.Size(); index++ {
			if g.global.Equal(rule.GetChild(index)) {
				fromCount++
			}
		}
	}
	if resultCount < 1 {
		return errors.New("Rule missed which result to global")
	}
	if resultCount > 1 || fromCount > 0 {
		virtualGlobal := NewNonterminal(g.NextTypes())
		virtualRule := NewRule(virtualGlobal, g.global)
		g.rules = append(g.rules, virtualRule)
		g.global = virtualGlobal
	}
	return nil
}
