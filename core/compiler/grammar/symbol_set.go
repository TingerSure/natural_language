package grammar

type SymbolSet struct {
	values map[Symbol]bool
}

func NewSymbolSet(values ...Symbol) *SymbolSet {
	set := &SymbolSet{
		values: map[Symbol]bool{},
	}

	for _, value := range values {
		set.Add(value)
	}
	return set
}

func (s *SymbolSet) Iterate(on func(Symbol) bool) bool {
	for value, _ := range s.values {
		if on(value) {
			return true
		}
	}
	return false
}

func (s *SymbolSet) Size() int {
	return len(s.values)
}

func (s *SymbolSet) Add(symbol Symbol) {
	s.values[symbol] = true
}

func (s *SymbolSet) Remove(symbol Symbol) {
	delete(s.values, symbol)
}

func (s *SymbolSet) Has(symbol Symbol) bool {
	return s.values[symbol]
}
