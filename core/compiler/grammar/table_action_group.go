package grammar

const (
	GroupActionType = iota
	GroupPolymerizeType
)

type TableActionGroup struct {
	types      int
	actions    map[int]*TableAction
	polymerize *TableAction
}

func NewTableActionGroup() *TableActionGroup {
	return &TableActionGroup{
		actions: map[int]*TableAction{},
		types:   GroupActionType,
	}
}

func NewTableActionGroupPolymerize(polymerizeRule *Rule) *TableActionGroup {
	return &TableActionGroup{
		polymerize: NewTableActionPolymerize(polymerizeRule),
		types:      GroupPolymerizeType,
	}
}

func (t *TableActionGroup) Type() int {
	return t.types
}

func (t *TableActionGroup) GetAction(symbol int) *TableAction {
	if t.types == GroupPolymerizeType {
		return t.polymerize
	}
	return t.actions[symbol]
}

func (t *TableActionGroup) SetAction(symbol int, action *TableAction) {
	t.actions[symbol] = action
}
