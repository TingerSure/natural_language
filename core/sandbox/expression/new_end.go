package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type NewEndSeed interface {
	ToLanguage(string, *NewEnd) string
	NewNull() concept.Null
	NewEnd() *interrupt.End
}

type NewEnd struct {
	*adaptor.ExpressionIndex
	seed NewEndSeed
}

func (f *NewEnd) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewEnd) ToString(prefix string) string {
	return "end"
}

func (a *NewEnd) Anticipate(space concept.Closure) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewEnd) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, a.seed.NewEnd()
}

type NewEndCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	EndCreator             func() *interrupt.End
}

type NewEndCreator struct {
	Seeds map[string]func(string, *NewEnd) string
	param *NewEndCreatorParam
}

func (s *NewEndCreator) New() *NewEnd {
	back := &NewEnd{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewEndCreator) NewEnd() *interrupt.End {
	return s.param.EndCreator()
}

func (s *NewEndCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *NewEndCreator) ToLanguage(language string, instance *NewEnd) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewEndCreator(param *NewEndCreatorParam) *NewEndCreator {
	return &NewEndCreator{
		Seeds: map[string]func(string, *NewEnd) string{},
		param: param,
	}
}
