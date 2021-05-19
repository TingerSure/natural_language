package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

type Grammar struct {
	rules    []*Rule
	global   *Nonterminal
	eof      *Terminal
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

func (g *Grammar) GetTable() *Table {
	return g.table
}

func (g *Grammar) AddRule(rule *Rule) {
	g.rules = append(g.rules, rule)
}

func (g *Grammar) SetGlobal(global *Nonterminal) {
	g.global = global
}

func (g *Grammar) SetEof(eof *Terminal) {
	g.eof = eof
}

func (g *Grammar) Build() (err error) {
	g.table.SetRules(g.rules)
	g.table.SetGlobal(g.global)
	g.table.SetEof(g.eof)
	err = g.table.Build()
	g.table.Clear()
	return
}

func (g *Grammar) Read(tokens *lexer.TokenList) (Phrase, error) {
	return g.automata.Run(tokens)
}
