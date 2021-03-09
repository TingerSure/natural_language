package grammar

import ()

const (
	AcceptType = iota
	MoveType
	GotoType
)

type TableAction struct {
	types  int
	status int
}

func NewTableActionAccept() *TableAction {
	return &TableAction{
		types: AcceptType,
	}
}

func NewTableActionMove(status int) *TableAction {
	return &TableAction{
		types:  MoveType,
		status: status,
	}
}

func NewTableActionGoto(status int) *TableAction {
	return &TableAction{
		types:  GotoType,
		status: status,
	}
}

func (t *TableAction) Types() int {
	return t.types
}

func (t *TableAction) Status() int {
	return t.status
}
