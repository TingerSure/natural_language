package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type NewStringSeed interface {
	ToLanguage(string, concept.Pool, *NewString) (string, concept.Exception)
	NewString(string) concept.String
}

type NewString struct {
	*adaptor.ExpressionIndex
	value string
	seed  NewStringSeed
}

func (f *NewString) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewString) ToString(prefix string) string {
	return fmt.Sprintf("\"%v\"", a.value)
}

func (a *NewString) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return a.seed.NewString(a.value), nil
}

type NewStringCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	StringCreator          func(string) concept.String
}

type NewStringCreator struct {
	Seeds map[string]func(concept.Pool, *NewString) (string, concept.Exception)
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

func (s *NewStringCreator) ToLanguage(language string, space concept.Pool, instance *NewString) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewStringCreator(param *NewStringCreatorParam) *NewStringCreator {
	return &NewStringCreator{
		Seeds: map[string]func(concept.Pool, *NewString) (string, concept.Exception){},
		param: param,
	}
}
