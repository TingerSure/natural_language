package interrupt

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	ExceptionInterruptType = "exception"
)

type ExceptionSeed interface {
	ToLanguage(string, *Exception) string
	Type() string
	New(concept.String, concept.String) *Exception
}

type Exception struct {
	name    concept.String
	message concept.String
	stacks  []concept.ExceptionStack
	seed    ExceptionSeed
}

func (f *Exception) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (e *Exception) IterateStacks(listener func(concept.ExceptionStack) bool) bool {
	for _, stack := range e.stacks {
		if listener(stack) {
			return true
		}
	}
	return false
}

func (e *Exception) AddStack(stack concept.ExceptionStack) concept.Exception {
	e.stacks = append(e.stacks, stack)
	return e
}

func (e *Exception) Copy() concept.Exception {
	newOne := e.seed.New(e.name.Clone(), e.message.Clone())
	e.IterateStacks(func(stack concept.ExceptionStack) bool {
		newOne.AddStack(stack)
		return false
	})
	return newOne
}

func (e *Exception) InterruptType() string {
	return e.seed.Type()
}

func (e *Exception) Name() concept.String {
	return e.name
}

func (e *Exception) Message() concept.String {
	return e.message
}

func (e *Exception) ToString(prefix string) string {
	return fmt.Sprintf("[%v] %v", e.name.ToString(prefix), e.message.ToString(prefix))
}

type ExceptionCreatorParam struct {
	StringCreator func(value string) concept.String
}

type ExceptionCreator struct {
	Seeds map[string]func(string, *Exception) string
	param *ExceptionCreatorParam
}

func (s *ExceptionCreator) NewOriginal(name string, message string) *Exception {
	return &Exception{
		name:    s.param.StringCreator(name),
		message: s.param.StringCreator(message),
		stacks:  make([]concept.ExceptionStack, 0),
		seed:    s,
	}
}

func (s *ExceptionCreator) New(name concept.String, message concept.String) *Exception {
	return &Exception{
		name:    name,
		message: message,
		stacks:  make([]concept.ExceptionStack, 0),
		seed:    s,
	}
}

func (s *ExceptionCreator) ToLanguage(language string, instance *Exception) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ExceptionCreator) Type() string {
	return ExceptionInterruptType
}

func NewExceptionCreator(param *ExceptionCreatorParam) *ExceptionCreator {
	return &ExceptionCreator{
		Seeds: map[string]func(string, *Exception) string{},
		param: param,
	}
}
