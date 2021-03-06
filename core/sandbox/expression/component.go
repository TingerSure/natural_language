package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type ComponentSeed interface {
	ToLanguage(string, concept.Pool, *Component) (string, concept.Exception)
}

type Component struct {
	*adaptor.ExpressionIndex
	field     concept.String
	object    concept.Pipe
	fieldLine concept.Line
	seed      ComponentSeed
}

func (f *Component) SetFieldLine(fieldLine concept.Line) {
	f.fieldLine = fieldLine
}

func (f *Component) Object() concept.Pipe {
	return f.object
}

func (f *Component) Field() concept.String {
	return f.field
}

func (f *Component) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Component) ToString(prefix string) string {
	return fmt.Sprintf("%v.%v", a.object.ToString(prefix), a.field.Value())
}

func (a *Component) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	value, suspend := object.GetField(a.field)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend.AddLine(a.fieldLine)
	}
	return value, nil
}

func (a *Component) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	return object.SetField(a.field, value)
}

func (a *Component) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend.(concept.Exception)
	}
	return object.Call(a.field, param)
}

type ComponentCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
}

type ComponentCreator struct {
	Seeds map[string]func(concept.Pool, *Component) (string, concept.Exception)
	param *ComponentCreatorParam
}

func (s *ComponentCreator) New(object concept.Pipe, field concept.String) *Component {
	back := &Component{
		field:  field,
		object: object,
		seed:   s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *ComponentCreator) ToLanguage(language string, space concept.Pool, instance *Component) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewComponentCreator(param *ComponentCreatorParam) *ComponentCreator {
	return &ComponentCreator{
		Seeds: map[string]func(concept.Pool, *Component) (string, concept.Exception){},
		param: param,
	}
}
