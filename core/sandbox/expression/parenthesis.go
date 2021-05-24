package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type ParenthesisSeed interface {
	ToLanguage(string, *Parenthesis) string
}

type Parenthesis struct {
	*adaptor.ExpressionIndex
	target concept.Index
	seed   ParenthesisSeed
}

func (f *Parenthesis) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Parenthesis) ToString(prefix string) string {
	return fmt.Sprintf("(%v)", a.target.ToString(prefix))
}

func (a *Parenthesis) Anticipate(space concept.Closure) concept.Variable {
	return a.target.Anticipate(space)
}

func (a *Parenthesis) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return a.target.Get(space)
}

func (a *Parenthesis) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return a.target.Set(space, value)
}

func (a *Parenthesis) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return a.target.Call(space, param)
}

type ParenthesisCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
}

type ParenthesisCreator struct {
	Seeds map[string]func(string, *Parenthesis) string
	param *ParenthesisCreatorParam
}

func (s *ParenthesisCreator) New(target concept.Index) *Parenthesis {
	back := &Parenthesis{
		target: target,
		seed:   s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *ParenthesisCreator) ToLanguage(language string, instance *Parenthesis) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewParenthesisCreator(param *ParenthesisCreatorParam) *ParenthesisCreator {
	return &ParenthesisCreator{
		Seeds: map[string]func(string, *Parenthesis) string{},
		param: param,
	}
}
