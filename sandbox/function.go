package sandbox

const (
	VariableFunctionType = "function"
)

type Function struct {
	flow       []Expression
	paramNames []string
	parent     *Closure
}

func (f *Function) AddParamName(paramName string) {
	f.paramNames = append(f.paramNames, paramName)
}

func (f *Function) AddStep(step Expression) {
	f.flow = append(f.flow, step)
}

func (f *Function) Exec(params map[string]Variable) (map[string]Variable, error) {
	space := NewClosure(f.parent)
	for _, name := range f.paramNames {
		space.InitLocal(name)
	}
	for name, value := range params {
		err := space.SetLocal(name, value)
		if err != nil {
			return nil, err
		}
	}
	for _, step := range f.flow {
		keep, err := step.Exec(space)
		if err != nil {
			return nil, err
		}
		if !keep {
			break
		}
	}
	returns := space.Return()
	space.Clear()
	return returns, nil
}

func (s *Function) Type() string {
	return VariableFunctionType
}

func NewFunction(parent *Closure) *Function {
	return &Function{
		parent: parent,
	}
}
