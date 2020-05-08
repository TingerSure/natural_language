package variable

const (
	VariableNullType = "null"
)

var (
	NullOnlyInstance = &Null{}
)

type Null struct {
}

var (
	NullLanguageSeeds = map[string]func(string, *Null) string{}
)

func (f *Null) ToLanguage(language string) string {
	seed := NullLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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
