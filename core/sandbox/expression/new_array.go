package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type NewArraySeed interface {
	ToLanguage(string, concept.Pool, *NewArray) (string, concept.Exception)
	NewArray() *variable.Array
}

type NewArray struct {
	*adaptor.ExpressionIndex
	items []concept.Pipe
	seed  NewArraySeed
}

func (f *NewArray) SetItems(items []concept.Pipe) {
	f.items = items
}

func (f *NewArray) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewArray) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	items := []string{}
	for _, item := range a.items {
		items = append(items, item.ToString(subPrefix))
	}
	return fmt.Sprintf("[ %v ]", strings.Join(items, ", "))
}

func (a *NewArray) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	array := a.seed.NewArray()
	for _, item := range a.items {
		value, suspend := item.Get(space)
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
		}
		array.Append(value)
	}
	return array, nil
}

type NewArrayCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ArrayCreator           func() *variable.Array
}

type NewArrayCreator struct {
	Seeds map[string]func(concept.Pool, *NewArray) (string, concept.Exception)
	param *NewArrayCreatorParam
}

func (s *NewArrayCreator) New() *NewArray {
	back := &NewArray{
		seed:  s,
		items: []concept.Pipe{},
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewArrayCreator) NewArray() *variable.Array {
	return s.param.ArrayCreator()
}

func (s *NewArrayCreator) ToLanguage(language string, space concept.Pool, instance *NewArray) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewArrayCreator(param *NewArrayCreatorParam) *NewArrayCreator {
	return &NewArrayCreator{
		Seeds: map[string]func(concept.Pool, *NewArray) (string, concept.Exception){},
		param: param,
	}
}
