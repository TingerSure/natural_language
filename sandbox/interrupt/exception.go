package interrupt

import (
	"fmt"
)

const (
	ExceptionInterruptType = "exception"
)

type Exception struct {
	name    string
	message string
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
	}
}
