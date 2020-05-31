package variable

const (
	VariableNullType = "null"
)

type NullSeed interface {
	ToLanguage(string, *Null) string
	Type() string
}

type Null struct {
	seed NullSeed
}

func (f *Null) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Null) ToString(prefix string) string {
	return "null"
}

func (n *Null) Type() string {
	return n.seed.Type()
}

type NullCreator struct {
	Seeds        map[string]func(string, *Null) string
	onlyInstance *Null
}

func (s *NullCreator) New() *Null {
	return s.onlyInstance
}

func (s *NullCreator) ToLanguage(language string, instance *Null) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *NullCreator) Type() string {
	return VariableNullType
}

func NewNullCreator() *NullCreator {
	instance := &NullCreator{
		Seeds: map[string]func(string, *Null) string{},
	}

	instance.onlyInstance = &Null{
		seed: instance,
	}

	return instance
}
