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
	ToLanguage(string, *NewArray) string
	NewArray() *variable.Array
}

type NewArray struct {
	*adaptor.ExpressionIndex
	items []concept.Index
	seed  NewArraySeed
}

func (f *NewArray) SetItems(items []concept.Index) {
	f.items = items
}

func (f *NewArray) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewArray) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	items := []string{}
	for _, item := range a.items {
		items = append(items, item.ToString(subPrefix))
	}
	return fmt.Sprintf("[ %v ]", strings.Join(items, ", "))
}

func (a *NewArray) Anticipate(space concept.Closure) concept.Variable {
	array := a.seed.NewArray()
	for _, item := range a.items {
		array.Append(item.Anticipate(space))
	}
	return array
}

func (a *NewArray) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
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
	Seeds map[string]func(string, *NewArray) string
	param *NewArrayCreatorParam
}

func (s *NewArrayCreator) New() *NewArray {
	back := &NewArray{
		seed:  s,
		items: []concept.Index{},
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewArrayCreator) NewArray() *variable.Array {
	return s.param.ArrayCreator()
}

func (s *NewArrayCreator) ToLanguage(language string, instance *NewArray) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewArrayCreator(param *NewArrayCreatorParam) *NewArrayCreator {
	return &NewArrayCreator{
		Seeds: map[string]func(string, *NewArray) string{},
		param: param,
	}
}
