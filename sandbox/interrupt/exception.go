package interrupt

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

const (
	ExceptionInterruptType = "exception"
)

type Exception struct {
	name    string
	message string
	stacks  []concept.ExceptionStack
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
	newOne := NewException(e.name, e.message)
	e.IterateStacks(func(stack concept.ExceptionStack) bool {
		newOne.AddStack(stack)
		return false
	})
	return newOne
}

func (e *Exception) InterruptType() string {
	return ExceptionInterruptType
}

func (e *Exception) Name() string {
	return e.name
}

func (e *Exception) Message() string {
	return e.message
}

func (e *Exception) ToString(prefix string) string {
	return fmt.Sprintf("[%v] %v", e.name, e.message)
}

func NewException(name string, message string) *Exception {
	return &Exception{
		name:    name,
		message: message,
		stacks:  make([]concept.ExceptionStack, 0),
	}
}
