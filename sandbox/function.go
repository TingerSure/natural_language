package sandbox

const (
	VariableFunctionType = "function"
)

type Function struct {
	flow       []Expression
	paramNames []string
}

func (f *Function) AddParamName(paramName string) {
	f.paramNames = append(f.paramNames, paramName)
}

func (f *Function) AddStep(step Expression) {
	f.flow = append(f.flow, step)
}

func (f *Function) Exec(params map[string]Variable) (*Closure, error) {
	space := NewClosure()
	for _, name := range f.paramNames {
		space.InitLocal(name)
	}
	for name, value := range params {
		err := space.SetLocal(name, value)
		if err != nil {
			return space, err
		}
	}
	for _, step := range f.flow {
		err := step.Exec(space)
		if err != nil {
			return space, err
		}
	}
	space.ClearCaches()
	return space, nil
}

func (s *Function) Type() string {
	return VariableFunctionType
}

func NewFunction() *Function {
	return &Function{}
}
