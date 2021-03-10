package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Grammar struct {
	rules    []*Rule
	global   *Nonterminal
	accept   *Terminal
	table    *Table
	automata *Automata
}

func NewGrammar() *Grammar {
	table := NewTable()
	return &Grammar{
		table:    table,
		automata: NewAutomata(table),
	}
}

func (g *Grammar) AddRule(rule *Rule) {
	g.rules = append(g.rules, rule)
}

func (g *Grammar) SetGlobal(global *Nonterminal) {
	g.SetGlobal(global)
}

func (g *Grammar) Build() error {
	g.table.SetRules(g.rules)
	g.table.SetGlobal(g.global)
	g.table.SetAccept(g.accept)
	return g.table.Build()
}

func (g *Grammar) Read([]*lexer.Token) (tree.Page, error) {
	return nil, nil
}
