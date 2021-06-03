package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableNumberType = "number"
)

type NumberSeed interface {
	ToLanguage(string, concept.Pool, *Number) (string, concept.Exception)
	Type() string
}

type Number struct {
	*adaptor.AdaptorVariable
	value float64
	seed  NumberSeed
}

func (o *Number) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Number) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Number) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}

func (n *Number) Value() float64 {
	return n.value
}

func (n *Number) Type() string {
	return n.seed.Type()
}

type NumberCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type NumberCreator struct {
	param *NumberCreatorParam
	Seeds map[string]func(concept.Pool, *Number) (string, concept.Exception)
}

func (s *NumberCreator) New(value float64) *Number {
	return &Number{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		value: value,
		seed:  s,
	}
}

func (s *NumberCreator) ToLanguage(language string, space concept.Pool, instance *Number) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *NumberCreator) Type() string {
	return VariableNumberType
}

func NewNumberCreator(param *NumberCreatorParam) *NumberCreator {
	return &NumberCreator{
		param: param,
		Seeds: map[string]func(concept.Pool, *Number) (string, concept.Exception){},
	}
}
