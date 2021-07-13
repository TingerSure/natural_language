package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type ParenthesisSeed interface {
	ToLanguage(string, concept.Pool, *Parenthesis) (string, concept.Exception)
}

type Parenthesis struct {
	*adaptor.ExpressionIndex
	target concept.Pipe
	seed   ParenthesisSeed
}

func (f *Parenthesis) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Parenthesis) ToString(prefix string) string {
	return fmt.Sprintf("(%v)", a.target.ToString(prefix))
}

func (a *Parenthesis) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return a.target.Get(space)
}

func (a *Parenthesis) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return a.target.Set(space, value)
}

func (a *Parenthesis) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	return a.target.Call(space, param)
}

type ParenthesisCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
}

type ParenthesisCreator struct {
	Seeds map[string]func(concept.Pool, *Parenthesis) (string, concept.Exception)
	param *ParenthesisCreatorParam
}

func (s *ParenthesisCreator) New(target concept.Pipe) *Parenthesis {
	back := &Parenthesis{
		target: target,
		seed:   s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *ParenthesisCreator) ToLanguage(language string, space concept.Pool, instance *Parenthesis) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewParenthesisCreator(param *ParenthesisCreatorParam) *ParenthesisCreator {
	return &ParenthesisCreator{
		Seeds: map[string]func(concept.Pool, *Parenthesis) (string, concept.Exception){},
		param: param,
	}
}
