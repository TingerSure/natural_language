package sandbox

const (
	VariableFunctionType = "function"
)

type Function struct {
}

func (s *Function) Type() string {
	return VariableFunctionType
}

func NewFunction() *Function {
	return &Function{}
}
