package grammar

const (
	SymbolTypeNonterminal = 1
)

type Nonterminal struct {
	types int
	name  string
}

func NewNonterminal(types int, name string) *Nonterminal {
	return &Nonterminal{
		types: types,
		name:  name,
	}
}

func (t *Nonterminal) SymbolType() int {
	return SymbolTypeNonterminal
}

func (t *Nonterminal) Type() int {
	return t.types
}

func (t *Nonterminal) Name() string {
	return t.name
}

func (t *Nonterminal) Equal(another Symbol) bool {
	return another.SymbolType() == SymbolTypeNonterminal && another.Type() == t.types
}
