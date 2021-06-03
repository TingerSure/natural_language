package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableBoolType = "bool"
)

type BoolSeed interface {
	ToLanguage(string, concept.Pool, *Bool) (string, concept.Exception)
	Type() string
}

type Bool struct {
	*adaptor.AdaptorVariable
	value bool
	seed  BoolSeed
}

func (o *Bool) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Bool) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Bool) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}

func (n *Bool) Value() bool {
	return n.value
}

func (n *Bool) Type() string {
	return n.seed.Type()
}

type BoolCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type BoolCreator struct {
	param *BoolCreatorParam
	Seeds map[string]func(concept.Pool, *Bool) (string, concept.Exception)
}

func (s *BoolCreator) New(value bool) *Bool {
	return &Bool{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		value: value,
		seed:  s,
	}
}

func (s *BoolCreator) ToLanguage(language string, space concept.Pool, instance *Bool) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *BoolCreator) Type() string {
	return VariableBoolType
}

func NewBoolCreator(param *BoolCreatorParam) *BoolCreator {
	return &BoolCreator{
		param: param,
		Seeds: map[string]func(concept.Pool, *Bool) (string, concept.Exception){},
	}
}
