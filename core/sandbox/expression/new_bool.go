package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type NewBoolSeed interface {
	ToLanguage(string, concept.Pool, *NewBool) string
	NewBool(bool) concept.Bool
}

type NewBool struct {
	*adaptor.ExpressionIndex
	value bool
	seed  NewBoolSeed
}

func (f *NewBool) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewBool) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}

func (a *NewBool) Anticipate(space concept.Pool) concept.Variable {
	return a.seed.NewBool(a.value)
}

func (a *NewBool) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return a.seed.NewBool(a.value), nil
}

type NewBoolCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	BoolCreator            func(bool) concept.Bool
}

type NewBoolCreator struct {
	Seeds map[string]func(string, concept.Pool, *NewBool) string
	param *NewBoolCreatorParam
}

func (s *NewBoolCreator) New(value bool) *NewBool {
	back := &NewBool{
		seed:  s,
		value: value,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewBoolCreator) NewBool(value bool) concept.Bool {
	return s.param.BoolCreator(value)
}

func (s *NewBoolCreator) ToLanguage(language string, space concept.Pool, instance *NewBool) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewNewBoolCreator(param *NewBoolCreatorParam) *NewBoolCreator {
	return &NewBoolCreator{
		Seeds: map[string]func(string, concept.Pool, *NewBool) string{},
		param: param,
	}
}
