package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type AssignmentSeed interface {
	ToLanguage(string, *Assignment) string
}

type Assignment struct {
	*adaptor.ExpressionIndex
	from concept.Index
	to   concept.Index
	seed AssignmentSeed
}

func (f *Assignment) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Assignment) ToString(prefix string) string {
	return fmt.Sprintf("%v = %v", a.to.ToString(prefix), a.from.ToString(prefix))
}

func (a *Assignment) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preFrom, suspend := a.from.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return preFrom, a.to.Set(space, preFrom)
}

type AssignmentCreatorParam struct {
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
}

type AssignmentCreator struct {
	Seeds map[string]func(string, *Assignment) string
	param *AssignmentCreatorParam
}

func (s *AssignmentCreator) New(from concept.Index, to concept.Index) *Assignment {
	back := &Assignment{
		from: from,
		to:   to,
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *AssignmentCreator) ToLanguage(language string, instance *Assignment) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewAssignmentCreator(param *AssignmentCreatorParam) *AssignmentCreator {
	return &AssignmentCreator{
		Seeds: map[string]func(string, *Assignment) string{},
		param: param,
	}
}
