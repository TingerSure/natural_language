package grammar

type TableClosure struct {
	projects []*TableProject
}

func NewTableClosure() *TableClosure {
	return &TableClosure{}
}
