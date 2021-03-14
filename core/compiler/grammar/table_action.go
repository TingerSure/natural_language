package grammar

import (
	"fmt"
)

const (
	ActionAcceptType = iota
	ActionMoveType
	ActionGotoType
	ActionPolymerizeType
)

type TableAction struct {
	types  int
	status int
	rule   *Rule
}

func NewTableActionPolymerize(rule *Rule) *TableAction {
	return &TableAction{
		types: ActionPolymerizeType,
		rule:  rule,
	}
}

func NewTableActionAccept() *TableAction {
	return &TableAction{
		types: ActionAcceptType,
	}
}

func NewTableActionMove(status int) *TableAction {
	return &TableAction{
		types:  ActionMoveType,
		status: status,
	}
}

func NewTableActionGoto(status int) *TableAction {
	return &TableAction{
		types:  ActionGotoType,
		status: status,
	}
}

func (t *TableAction) Type() int {
	return t.types
}

func (t *TableAction) Status() int {
	return t.status
}

func (t *TableAction) Rule() *Rule {
	return t.rule
}

func (t *TableAction) ToString() string {
	if t == nil {
		return ""
	}
	if t.Type() == ActionAcceptType {
		return "Accept"
	}
	if t.Type() == ActionMoveType {
		return fmt.Sprintf("move %v", t.Status())
	}
	if t.Type() == ActionPolymerizeType {
		return fmt.Sprintf("%v", t.rule.ToString())
	}
	if t.Type() == ActionGotoType {
		return fmt.Sprintf("goto %v", t.Status())
	}
	return "unknown"
}
