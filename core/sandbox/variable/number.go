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
	ToLanguage(string, concept.Closure, *Number) string
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

func (f *Number) ToLanguage(language string, space concept.Closure) string {
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
	Seeds map[string]func(string, concept.Closure, *Number) string
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

func (s *NumberCreator) ToLanguage(language string, space concept.Closure, instance *Number) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *NumberCreator) Type() string {
	return VariableNumberType
}

func NewNumberCreator(param *NumberCreatorParam) *NumberCreator {
	return &NumberCreator{
		param: param,
		Seeds: map[string]func(string, concept.Closure, *Number) string{},
	}
}
