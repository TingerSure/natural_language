package sandbox

const (
	VariableBoolType = "bool"
)

type Bool struct {
	value bool
}

func (n *Bool) Value() bool {
	return n.value
}

func (n *Bool) Type() string {
	return VariableBoolType
}

func NewBool(value bool) *Bool {
	return &Bool{
		value: value,
	}
}
