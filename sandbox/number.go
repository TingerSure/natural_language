package sandbox

const (
	VariableNumberType = "number"
)

type Number struct {
	value float64
}

func (n *Number) Value() float64 {
	return n.value
}

func (n *Number) Type() string {
	return VariableNumberType
}

func NewNumber(value float64) *Number {
	return &Number{
		value: value,
	}
}
