package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Reach struct {
	structs []*tree.StructRule
}

func NewReach() *Reach {
	return &Reach{}
}

func (g *Reach) GetRulesByLastType(given *tree.PhraseType) []*tree.StructRule {
	back := []*tree.StructRule{}
	for _, rule := range g.structs {
		types := rule.Types()
		if types[len(types)-1].Match(given) {
			back = append(back, rule)
		}
	}
	return back
}

func (g *Reach) AddRule(rules []*tree.StructRule) {
	if rules == nil {
		return
	}
	g.structs = append(g.structs, rules...)
}

func (g *Reach) RemoveRule(need func(rule *tree.StructRule) bool) {
	for index := 0; index < len(g.structs); index++ {
		rule := g.structs[index]
		if need(rule) {
			g.structs = append(g.structs[:index], g.structs[index+1:]...)
		}
	}
}
