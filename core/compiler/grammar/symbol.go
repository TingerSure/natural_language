package grammar

import ()

type Symbol interface {
	SymbolType() int
	Type() int
	Equal(Symbol) bool
}
