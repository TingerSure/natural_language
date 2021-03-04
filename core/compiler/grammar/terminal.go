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
