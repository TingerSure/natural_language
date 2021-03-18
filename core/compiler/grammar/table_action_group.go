package grammar

type TableActionGroup struct {
	actions map[int]*TableAction
}

func NewTableActionGroup() *TableActionGroup {
	return &TableActionGroup{
		actions: map[int]*TableAction{},
	}
}

func (t *TableActionGroup) GetAction(symbol int) *TableAction {
	return t.actions[symbol]
}

func (t *TableActionGroup) SetAction(symbol int, action *TableAction) {
	t.actions[symbol] = action
}
