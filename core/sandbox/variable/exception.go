package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	ExceptionInterruptType = "exception"
	VariableExceptionType  = "exception"
)

type ExceptionSeed interface {
	ToLanguage(string, concept.Pool, *Exception) (string, concept.Exception)
	Type() string
	InterruptType() string
	New(concept.String, concept.String) concept.Exception
}

type Exception struct {
	*adaptor.AdaptorVariable
	name    concept.String
	message concept.String
	lines   []concept.Line
	seed    ExceptionSeed
}

func (o *Exception) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Exception) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (e *Exception) IterateLines(listener func(concept.Line) bool) bool {
	for _, line := range e.lines {
		if listener(line) {
			return true
		}
	}
	return false
}

func (e *Exception) AddLine(line concept.Line) concept.Interrupt {
	return e.AddExceptionLine(line)
}

func (e *Exception) AddExceptionLine(line concept.Line) concept.Exception {
	e.lines = append(e.lines, line)
	return e
}

func (e *Exception) Copy() concept.Exception {
	newOne := e.seed.New(e.name.Clone(), e.message.Clone())
	e.IterateLines(func(line concept.Line) bool {
		newOne.AddLine(line)
		return false
	})
	return newOne
}

func (e *Exception) InterruptType() string {
	return e.seed.InterruptType()
}

func (o *Exception) Type() string {
	return o.seed.Type()
}

func (e *Exception) Name() concept.String {
	return e.name
}

func (e *Exception) Message() concept.String {
	return e.message
}

func (e *Exception) ToString(prefix string) string {
	return fmt.Sprintf("[%v] %v", e.name.Value(), e.message.Value())
}

func (e *Exception) Error() string {
	return e.ToString("")
}

type ExceptionCreatorParam struct {
	StringCreator func(value string) concept.String
	NullCreator   func() concept.Null
}

type ExceptionCreator struct {
	Seeds map[string]func(concept.Pool, *Exception) (string, concept.Exception)
	param *ExceptionCreatorParam
}

func (s *ExceptionCreator) NewOriginal(name string, message string) concept.Exception {
	return &Exception{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.NewOriginal,
		}),
		name:    s.param.StringCreator(name),
		message: s.param.StringCreator(message),
		lines:   make([]concept.Line, 0),
		seed:    s,
	}
}

func (s *ExceptionCreator) New(name concept.String, message concept.String) concept.Exception {
	return &Exception{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.NewOriginal,
		}),
		name:    name,
		message: message,
		lines:   make([]concept.Line, 0),
		seed:    s,
	}
}

func (s *ExceptionCreator) ToLanguage(language string, space concept.Pool, instance *Exception) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *ExceptionCreator) InterruptType() string {
	return ExceptionInterruptType
}

func (s *ExceptionCreator) Type() string {
	return VariableExceptionType
}

func NewExceptionCreator(param *ExceptionCreatorParam) *ExceptionCreator {
	return &ExceptionCreator{
		Seeds: map[string]func(concept.Pool, *Exception) (string, concept.Exception){},
		param: param,
	}
}
