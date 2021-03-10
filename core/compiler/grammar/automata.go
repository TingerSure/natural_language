package grammar

import ()

type Automata struct {
	table *Table
}

func NewAutomata(table *Table) *Automata {
	return &Automata{
		table: table,
	}
}

func (a *Automata) Run() (Phrase, error) {

	return nil, nil
}
