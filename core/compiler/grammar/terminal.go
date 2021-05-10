package grammar

const (
	SymbolTypeTerminal = 0
)

type Terminal struct {
	tokenTypes int
	name       string
}

func NewTerminal(tokenTypes int, name string) *Terminal {
	return &Terminal{
		tokenTypes: tokenTypes,
		name:       name,
	}
}

func (t *Terminal) SymbolType() int {
	return SymbolTypeTerminal
}

func (t *Terminal) Name() string {
	return t.name
}

func (t *Terminal) Type() int {
	return t.tokenTypes
}

func (t *Terminal) Equal(another Symbol) bool {
	return another.SymbolType() == SymbolTypeTerminal && another.Type() == t.tokenTypes
}
