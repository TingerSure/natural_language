package grammar

import (
	"errors"
	// "github.com/TingerSure/natural_language/core/compiler/lexer"
)

type Table struct {
	typesCount int
	rules      []*Rule
	global     *Nonterminal
	actions    map[int]map[int]*TableAction // map[status]map[symbol]action
	gotos      map[int]map[int]*TableAction // map[status]map[symbol]goto
	status     []int
	phrases    []Phrase
}

func NewTable() *Table {
	return &Table{
		typesCount: 0,
		actions:    map[int]map[int]*TableAction{},
		gotos:      map[int]map[int]*TableAction{},
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
	// projects := g.makeProjects(g.rules)
	// closures := g.makeClosures(projects)
	// TODO
	return nil
}

func (g *Table) makeClosures(projects [][]*TableProject) []*TableClosure {
	// TODO
	return nil
}

func (g *Table) makeProjects(rules []*Rule) [][]*TableProject {
	projects := [][]*TableProject{}
	for ruleIndex, rule := range rules {
		projects = append(projects, []*TableProject{})
		for childIndex := 0; childIndex < rule.Size(); childIndex++ {
			projects[ruleIndex] = append(projects[ruleIndex], NewTableProject(rule, childIndex))
		}
	}
	return projects
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
