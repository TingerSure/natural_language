package grammar

import ()

const (
	SymbolTypeTerminal = 0
)

type Terminal struct {
	tokenTypes int
}

func (t *Terminal) SymbolType() int {
	return SymbolTypeTerminal
}

func (t *Terminal) Type() int {
	return t.tokenTypes
}

func (t *Terminal) Equal(another Symbol) bool {
	return another.SymbolType() == SymbolTypeTerminal && another.Type() == t.tokenTypes
}
