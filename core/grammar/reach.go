package grammar

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Reach struct {
	structs []*tree.StructRule
}

func NewReach() *Reach {
	return &Reach{}
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

func (r *Reach) Check(lake *Lake, onStruct func(*tree.StructRule)) {
	if lake.IsEmpty() {
		return
	}
	for _, twig := range r.structs {
		if lake.Len() < twig.Size() {
			continue
		}
		if twig.Match(lake.PeekAll()) {
			onStruct(twig)
		}
	}
}
