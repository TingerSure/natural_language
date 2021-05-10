package grammar

type Symbol interface {
	SymbolType() int
	Type() int
	Name() string
	Equal(Symbol) bool
}
