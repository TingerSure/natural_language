package variable

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariableNullType = "null"
)

type NullSeed interface {
	ToLanguage(string, concept.Pool, *Null) (string, concept.Exception)
	Type() string
	GetNullPointerException() concept.Exception
}

type Null struct {
	seed NullSeed
}

func (o *Null) IsFunction() bool {
	return false
}

func (o *Null) IsNull() bool {
	return true
}

func (o *Null) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	return o.seed.GetNullPointerException().Copy()
}

func (o *Null) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	return nil, o.seed.GetNullPointerException().Copy()
}

func (o *Null) KeyField(specimen concept.String) concept.String {
	return specimen
}

func (o *Null) SizeField() int {
	return 0
}

func (o *Null) Iterate(on func(concept.String, concept.Variable) bool) bool {
	return false
}

func (m *Null) HasField(specimen concept.String) bool {
	return false
}

func (o *Null) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return nil, o.seed.GetNullPointerException().Copy()
}

func (f *Null) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Null) ToString(prefix string) string {
	return "null"
}

func (n *Null) Type() string {
	return n.seed.Type()
}

type NullCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
}

type NullCreator struct {
	Seeds                map[string]func(concept.Pool, *Null) (string, concept.Exception)
	param                *NullCreatorParam
	onlyInstance         *Null
	nullPointerException concept.Exception
}

func (s *NullCreator) New() *Null {
	return s.onlyInstance
}

func (s *NullCreator) ToLanguage(language string, space concept.Pool, instance *Null) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *NullCreator) Type() string {
	return VariableNullType
}

func (s *NullCreator) GetNullPointerException() concept.Exception {
	if nl_interface.IsNil(s.nullPointerException) {
		s.nullPointerException = s.param.ExceptionCreator("runtime error", "null pointer exception.")
	}
	return s.nullPointerException
}

func NewNullCreator(param *NullCreatorParam) *NullCreator {
	instance := &NullCreator{
		param: param,
		Seeds: map[string]func(concept.Pool, *Null) (string, concept.Exception){},
	}

	instance.onlyInstance = &Null{
		seed: instance,
	}

	return instance
}
