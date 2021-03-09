package grammar

const (
	ActionType = iota
	PolymerizeType
)

type TableActionGroup struct {
	types          int
	actions        map[int]*TableAction
	polymerizeRule *Rule
}

func NewTableActionGroup() *TableActionGroup {
	return &TableActionGroup{
		actions: map[int]*TableAction{},
		types:   ActionType,
	}
}

func NewTableActionPolymerize(polymerizeRule *Rule) *TableActionGroup {
	return &TableActionGroup{
		polymerizeRule: polymerizeRule,
		types:          PolymerizeType,
	}
}

func (t *TableActionGroup) Type() int {
	return t.types
}

func (t *TableActionGroup) GetAction(symbol int) *TableAction {
	return t.actions[symbol]
}

func (t *TableActionGroup) SetAction(symbol int, action *TableAction) {
	t.actions[symbol] = action
}

func (t *TableActionGroup) GetPolymerizeRule() *Rule {
	return t.polymerizeRule
}
