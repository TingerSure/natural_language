package grammar

import (
	"fmt"
	"strings"
)

type TableProject struct {
	Rule  *Rule
	Index int
	Next  *TableProject
}

func NewTableProject(rule *Rule, index int) *TableProject {
	return &TableProject{
		Rule:  rule,
		Index: index,
		Next:  nil,
	}
}

func (t *TableProject) IsStart() bool {
	return t.Index == 0
}

func (t *TableProject) IsEnd() bool {
	return t.Index == t.Rule.Size()
}

func (t *TableProject) GetNextChild() Symbol {
	if t.IsEnd() {
		return nil
	}
	return t.Rule.GetChild(t.Index)
}

func (t *TableProject) ToString() string {
	names := []string{}
	for index := 0; index < t.Rule.Size(); index++ {
		if index == t.Index {
			names = append(names, "•")
		}
		names = append(names, t.Rule.GetChild(index).Name())
	}
	if t.IsEnd() {
		names = append(names, "•")
	}
	return fmt.Sprintf("%v -> %v", t.Rule.GetResult().Name(), strings.Join(names, " "))
}
