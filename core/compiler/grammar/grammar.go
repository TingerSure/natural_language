package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Grammar struct {
	table *Table
}

func NewGrammar() *Grammar {
	return &Grammar{
		table: NewTable(),
	}
}

func (g *Grammar) AddRule(rule *Rule) {
	g.table.AddRule(rule)
}

func (g *Grammar) SetGlobal(global *Nonterminal) {
	g.SetGlobal(global)
}

func (g *Grammar) Build() error {
	return g.table.Build()
}

func (g *Grammar) Read([]*lexer.Token) (tree.Page, error) {
	return nil, nil
}
