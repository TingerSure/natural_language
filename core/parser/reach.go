package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Reach struct {
	structs []*tree.StructRule
	types   *Types
}

func NewReach(types *Types) *Reach {
	return &Reach{
		types: types,
	}
}

func (g *Reach) GetRulesByLastType(given string) []*tree.StructRule {
	back := []*tree.StructRule{}
	for _, rule := range g.structs {
		wanteds := rule.Types()
		if g.types.Match(wanteds[len(wanteds)-1], given) {
			back = append(back, rule)
		}
	}
	return back
}

func (g *Reach) AddRule(rule *tree.StructRule) {
	if rule == nil {
		return
	}
	g.structs = append(g.structs, rule)
}

func (g *Reach) RemoveRule(need func(rule *tree.StructRule) bool) {
	for index := 0; index < len(g.structs); index++ {
		rule := g.structs[index]
		if need(rule) {
			g.structs = append(g.structs[:index], g.structs[index+1:]...)
		}
	}
}
