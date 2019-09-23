package sandbox

const (
	VariableFunctionType = "function"
)

type Function struct {
	body       *CodeBlock
	paramNames []string
	parent     *Closure
}

func (f *Function) AddParamName(paramName string) {
	f.paramNames = append(f.paramNames, paramName)
}

func (f *Function) AddStep(step Expression) {
	f.body.AddStep(step)
}

func (f *Function) Exec(params map[string]Variable) (map[string]Variable, error) {

	space, _, err := f.body.Exec(f.parent, false, func(space *Closure) error {
		for _, name := range f.paramNames {
			space.InitLocal(name)
		}
		for name, value := range params {
			err := space.SetLocal(name, value)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return space.Return(), nil
}

func (s *Function) Type() string {
	return VariableFunctionType
}

func NewFunction(parent *Closure) *Function {
	return &Function{
		parent: parent,
		body:   NewCodeBlock(),
	}
}
