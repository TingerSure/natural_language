package grammar

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
	return t.Rule.GetChild(t.Index)
}
