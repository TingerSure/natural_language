package interrupt

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	ExceptionInterruptType = "exception"
)

type Exception struct {
	name    concept.String
	message concept.String
	stacks  []concept.ExceptionStack
}

var (
	ExceptionLanguageSeeds = map[string]func(string, *Exception) string{}
)

func (f *Exception) ToLanguage(language string) string {
	seed := ExceptionLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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
	newOne := NewException(e.name.Clone(), e.message.Clone())
	e.IterateStacks(func(stack concept.ExceptionStack) bool {
		newOne.AddStack(stack)
		return false
	})
	return newOne
}

func (e *Exception) InterruptType() string {
	return ExceptionInterruptType
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

func NewException(name concept.String, message concept.String) *Exception {
	return &Exception{
		name:    name,
		message: message,
		stacks:  make([]concept.ExceptionStack, 0),
	}
}
