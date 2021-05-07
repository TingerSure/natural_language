package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type NewStringSeed interface {
	ToLanguage(string, *NewString) string
	NewString(string) concept.String
}

type NewString struct {
	*adaptor.ExpressionIndex
	value string
	seed  NewStringSeed
}

func (f *NewString) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewString) ToString(prefix string) string {
	return fmt.Sprintf("\"%v\"", a.value)
}

func (a *NewString) Anticipate(space concept.Closure) concept.Variable {
	return a.seed.NewString(a.value)
}

func (a *NewString) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return a.seed.NewString(a.value), nil
}

type NewStringCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	StringCreator          func(string) concept.String
}

type NewStringCreator struct {
	Seeds map[string]func(string, *NewString) string
	param *NewStringCreatorParam
}

func (s *NewStringCreator) New(value string) *NewString {
	back := &NewString{
		seed:  s,
		value: value,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewStringCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func (s *NewStringCreator) ToLanguage(language string, instance *NewString) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewStringCreator(param *NewStringCreatorParam) *NewStringCreator {
	return &NewStringCreator{
		Seeds: map[string]func(string, *NewString) string{},
		param: param,
	}
}
