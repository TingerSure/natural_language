package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Grammar struct {
	rules []*Rule
}

func NewGrammar() *Grammar {
	return &Grammar{}
}

func (g *Grammar) AddRule(rule *Rule) {
	g.rules = append(g.rules, rule)
}

func (g *Grammar) Build() {
	// TODO create goto/action table
}

func (g *Grammar) Read([]*lexer.Token) (tree.Page, error) {
	return nil, nil
}
