package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type FunctionEndSeed interface {
	ToLanguage(string, *FunctionEnd) string
	NewEnd() *interrupt.End
	NewNull() concept.Null
}

type FunctionEnd struct {
	*adaptor.ExpressionIndex
	seed FunctionEndSeed
}

func (f *FunctionEnd) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *FunctionEnd) ToString(prefix string) string {
	return fmt.Sprintf("end")
}

func (e *FunctionEnd) Anticipate(space concept.Closure) concept.Variable {
	return e.seed.NewNull()
}

func (a *FunctionEnd) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return a.seed.NewNull(), a.seed.NewEnd()
}

type FunctionEndCreatorParam struct {
	EndCreator             func() *interrupt.End
	NullCreator            func() concept.Null
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
}

type FunctionEndCreator struct {
	Seeds        map[string]func(string, *FunctionEnd) string
	param        *FunctionEndCreatorParam
	defaultParam concept.Index
}

func (s *FunctionEndCreator) New() *FunctionEnd {
	back := &FunctionEnd{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *FunctionEndCreator) ToLanguage(language string, instance *FunctionEnd) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *FunctionEndCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *FunctionEndCreator) NewEnd() *interrupt.End {
	return s.param.EndCreator()
}

func NewFunctionEndCreator(param *FunctionEndCreatorParam) *FunctionEndCreator {
	return &FunctionEndCreator{
		Seeds: map[string]func(string, *FunctionEnd) string{},
		param: param,
	}
}
