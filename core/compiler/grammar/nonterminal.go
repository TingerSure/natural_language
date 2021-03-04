package grammar

import ()

const (
	SymbolTypeNonterminal = 1
)

type Nonterminal struct {
	types int
}

func (t *Nonterminal) SymbolType() int {
	return SymbolTypeNonterminal
}

func (t *Nonterminal) Type() int {
	return t.types
}
