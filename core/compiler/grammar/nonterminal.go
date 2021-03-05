package grammar

import ()

const (
	SymbolTypeNonterminal = 1
)

type Nonterminal struct {
	types int
}

func NewNonterminal(types int) *Nonterminal {
	return &Nonterminal{
		types: types,
	}
}

func (t *Nonterminal) SymbolType() int {
	return SymbolTypeNonterminal
}

func (t *Nonterminal) Type() int {
	return t.types
}

func (t *Nonterminal) Equal(another Symbol) bool {
	return another.SymbolType() == SymbolTypeNonterminal && another.Type() == t.types
}
