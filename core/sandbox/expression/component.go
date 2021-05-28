package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type ComponentSeed interface {
	ToLanguage(string, concept.Closure, *Component) string
}

type Component struct {
	*adaptor.ExpressionIndex
	field  concept.String
	object concept.Index
	seed   ComponentSeed
}

func (f *Component) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Component) ToString(prefix string) string {
	return fmt.Sprintf("%v.%v", a.object.ToString(prefix), a.field.Value())
}

func (a *Component) Anticipate(space concept.Closure) concept.Variable {
	value, _ := a.object.Anticipate(space).GetField(a.field)
	return value
}

func (a *Component) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return object.GetField(a.field)
}

func (a *Component) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	return object.SetField(a.field, value)
}

func (a *Component) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
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
	Seeds map[string]func(string, concept.Closure, *Component) string
	param *ComponentCreatorParam
}

func (s *ComponentCreator) New(object concept.Index, field concept.String) *Component {
	back := &Component{
		field:  field,
		object: object,
		seed:   s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *ComponentCreator) ToLanguage(language string, space concept.Closure, instance *Component) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewComponentCreator(param *ComponentCreatorParam) *ComponentCreator {
	return &ComponentCreator{
		Seeds: map[string]func(string, concept.Closure, *Component) string{},
		param: param,
	}
}
