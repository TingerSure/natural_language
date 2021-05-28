package variable

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableObjectType = "object"
)

type ObjectSeed interface {
	ToLanguage(string, concept.Closure, *Object) string
	Type() string
}

type Object struct {
	*adaptor.AdaptorVariable
	seed ObjectSeed
}

func (o *Object) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Object) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (o *Object) Type() string {
	return o.seed.Type()
}

type ObjectCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type ObjectCreator struct {
	Seeds map[string]func(string, concept.Closure, *Object) string
	param *ObjectCreatorParam
}

func (s *ObjectCreator) New() *Object {
	return &Object{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		seed: s,
	}
}

func (s *ObjectCreator) ToLanguage(language string, space concept.Closure, instance *Object) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *ObjectCreator) Type() string {
	return VariableObjectType
}

func NewObjectCreator(param *ObjectCreatorParam) *ObjectCreator {
	return &ObjectCreator{
		Seeds: map[string]func(string, concept.Closure, *Object) string{},
		param: param,
	}
}
