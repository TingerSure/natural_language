package grammar

type TableProject struct {
	Rule  *Rule
	Index int
}

func NewTableProject(rule *Rule, index int) *TableProject {
	return &TableProject{
		Rule:  rule,
		Index: index,
	}
}
