package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type AssignmentSeed interface {
	ToLanguage(string, concept.Pool, *Assignment) (string, concept.Exception)
}

type Assignment struct {
	*adaptor.ExpressionIndex
	from concept.Pipe
	to   concept.Pipe
	seed AssignmentSeed
}

func (f *Assignment) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Assignment) ToString(prefix string) string {
	return fmt.Sprintf("%v = %v", a.to.ToString(prefix), a.from.ToString(prefix))
}

func (e *Assignment) Anticipate(space concept.Pool) concept.Variable {
	return e.from.Anticipate(space)
}

func (a *Assignment) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	preFrom, suspend := a.from.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return preFrom, a.to.Set(space, preFrom)
}

type AssignmentCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
}

type AssignmentCreator struct {
	Seeds map[string]func(concept.Pool, *Assignment) (string, concept.Exception)
	param *AssignmentCreatorParam
}

func (s *AssignmentCreator) New(from concept.Pipe, to concept.Pipe) *Assignment {
	back := &Assignment{
		from: from,
		to:   to,
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *AssignmentCreator) ToLanguage(language string, space concept.Pool, instance *Assignment) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewAssignmentCreator(param *AssignmentCreatorParam) *AssignmentCreator {
	return &AssignmentCreator{
		Seeds: map[string]func(concept.Pool, *Assignment) (string, concept.Exception){},
		param: param,
	}
}
