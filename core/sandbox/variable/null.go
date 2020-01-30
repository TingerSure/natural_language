package variable

const (
	VariableNullType = "null"
)

var (
	NullOnlyInstance = &Null{}
)

type Null struct {
}

func (a *Null) ToString(prefix string) string {
	return "null"
}

func (n *Null) Type() string {
	return VariableNullType
}

func NewNull() *Null {
	return NullOnlyInstance
}
