package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type NewNumberSeed interface {
	ToLanguage(string, concept.Closure, *NewNumber) string
	NewNumber(float64) concept.Number
}

type NewNumber struct {
	*adaptor.ExpressionIndex
	value float64
	seed  NewNumberSeed
}

func (f *NewNumber) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewNumber) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}

func (a *NewNumber) Anticipate(space concept.Closure) concept.Variable {
	return a.seed.NewNumber(a.value)
}

func (a *NewNumber) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return a.seed.NewNumber(a.value), nil
}

type NewNumberCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NumberCreator          func(float64) concept.Number
}

type NewNumberCreator struct {
	Seeds map[string]func(string, concept.Closure, *NewNumber) string
	param *NewNumberCreatorParam
}

func (s *NewNumberCreator) New(value float64) *NewNumber {
	back := &NewNumber{
		seed:  s,
		value: value,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewNumberCreator) NewNumber(value float64) concept.Number {
	return s.param.NumberCreator(value)
}

func (s *NewNumberCreator) ToLanguage(language string, space concept.Closure, instance *NewNumber) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewNewNumberCreator(param *NewNumberCreatorParam) *NewNumberCreator {
	return &NewNumberCreator{
		Seeds: map[string]func(string, concept.Closure, *NewNumber) string{},
		param: param,
	}
}
