package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type AppendSeed interface {
	ToLanguage(string, concept.Pool, *Append) string
	NewException(string, string) concept.Exception
}

type Append struct {
	*adaptor.ExpressionIndex
	array concept.Pipe
	item  concept.Pipe
	seed  AppendSeed
}

func (f *Append) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Append) ToString(prefix string) string {
	return fmt.Sprintf("%v <- %v", a.array.ToString(prefix), a.item.ToString(prefix))
}

func (e *Append) Anticipate(space concept.Pool) concept.Variable {
	return e.array.Anticipate(space)
}

func (a *Append) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	item, suspend := a.item.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	arrayPre, suspend := a.array.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	array, yes := variable.VariableFamilyInstance.IsArray(arrayPre)
	if !yes {
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("%v is not an array", a.array.ToString("")))
	}
	array.Append(item)
	return array, nil
}

type AppendCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ExceptionCreator       func(string, string) concept.Exception
}

type AppendCreator struct {
	Seeds map[string]func(string, concept.Pool, *Append) string
	param *AppendCreatorParam
}

func (s *AppendCreator) New(array concept.Pipe, item concept.Pipe) *Append {
	back := &Append{
		array: array,
		item:  item,
		seed:  s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *AppendCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *AppendCreator) ToLanguage(language string, space concept.Pool, instance *Append) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewAppendCreator(param *AppendCreatorParam) *AppendCreator {
	return &AppendCreator{
		Seeds: map[string]func(string, concept.Pool, *Append) string{},
		param: param,
	}
}
