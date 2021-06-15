package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type DefineSeed interface {
	ToLanguage(string, concept.Pool, *Define) (string, concept.Exception)
	NewNull() concept.Null
	NewException(string, string) concept.Exception
}

type Define struct {
	*adaptor.ExpressionIndex
	defaultValue concept.Pipe
	key          concept.String
	line         concept.Line
	seed         DefineSeed
}

func (f *Define) SetLine(line concept.Line) {
	f.line = line
}

func (f *Define) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Define) ToString(prefix string) string {
	if a.defaultValue == nil {
		return fmt.Sprintf("var %v", a.key.Value())
	}
	return fmt.Sprintf("var %v = %v", a.key.Value(), a.defaultValue.ToString(prefix))
}

func (e *Define) Anticipate(space concept.Pool) concept.Variable {
	return e.defaultValue.Anticipate(space)
}

func (a *Define) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	if space.HasLocal(a.key) {
		return nil, a.seed.NewException("semantic error", fmt.Sprintf("Duplicate local definition : %v", a.key.Value())).AddLine(a.line)
	}
	var defaultValue concept.Variable
	var suspend concept.Interrupt
	if a.defaultValue != nil {
		defaultValue, suspend = a.defaultValue.Get(space)
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
		}
	} else {
		defaultValue = a.seed.NewNull()
	}

	space.InitLocal(a.key, defaultValue)
	return defaultValue, nil
}

type DefineCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	ExceptionCreator       func(string, string) concept.Exception
}

type DefineCreator struct {
	Seeds map[string]func(concept.Pool, *Define) (string, concept.Exception)
	param *DefineCreatorParam
}

func (s *DefineCreator) New(key concept.String, defaultValue concept.Pipe) *Define {
	back := &Define{
		defaultValue: defaultValue,
		key:          key,
		seed:         s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *DefineCreator) ToLanguage(language string, space concept.Pool, instance *Define) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *DefineCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *DefineCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewDefineCreator(param *DefineCreatorParam) *DefineCreator {
	return &DefineCreator{
		Seeds: map[string]func(concept.Pool, *Define) (string, concept.Exception){},
		param: param,
	}
}
