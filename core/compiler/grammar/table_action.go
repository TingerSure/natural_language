package grammar

import ()

const (
	End = iota
	Move
	Polymerize
	Goto
)

type TableAction struct {
	types  int
	status int
	rule   int
}

func (t *TableAction) Types() int {
	return t.types
}

func (t *TableAction) Status() int {
	return t.status
}

func (t *TableAction) Rule() int {
	return t.rule
}
