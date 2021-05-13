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
	types    int
	status   int
	rule     *Rule
	projects map[*TableProject]*SymbolSet
}

func NewTableActionPolymerize(rule *Rule, projects map[*TableProject]*SymbolSet) *TableAction {
	return &TableAction{
		types:    ActionPolymerizeType,
		rule:     rule,
		projects: projects,
	}
}

func NewTableActionAccept(projects map[*TableProject]*SymbolSet) *TableAction {
	return &TableAction{
		types:    ActionAcceptType,
		projects: projects,
	}
}

func NewTableActionMove(status int, projects map[*TableProject]*SymbolSet) *TableAction {
	return &TableAction{
		types:    ActionMoveType,
		status:   status,
		projects: projects,
	}
}

func NewTableActionGoto(status int, projects map[*TableProject]*SymbolSet) *TableAction {
	return &TableAction{
		types:    ActionGotoType,
		status:   status,
		projects: projects,
	}
}

func (t *TableAction) Type() int {
	return t.types
}

func (t *TableAction) Status() int {
	return t.status
}

func (t *TableAction) SetStatus(status int) {
	t.status = status
}

func (t *TableAction) Rule() *Rule {
	return t.rule
}

func (t *TableAction) Projects() map[*TableProject]*SymbolSet {
	return t.projects
}

func (t *TableAction) ToString() string {
	if t == nil {
		return ""
	}
	if t.Type() == ActionAcceptType {
		return "accept"
	}
	if t.Type() == ActionMoveType {
		return fmt.Sprintf("move %v", t.Status())
	}
	if t.Type() == ActionPolymerizeType {
		return t.rule.ToString()
	}
	if t.Type() == ActionGotoType {
		return fmt.Sprintf("goto %v", t.Status())
	}
	return "unknown"
}
