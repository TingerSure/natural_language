package interrupt

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	ReturnInterruptType = "return"
)

type ReturnSeed interface {
	Type() string
}

type Return struct {
	seed ReturnSeed
}

func (e *Return) InterruptType() string {
	return e.seed.Type()
}

func (e *Return) IterateLines(func(concept.Line) bool) bool {
	return false
}

func (e *Return) AddLine(concept.Line) concept.Interrupt {
	return e
}

type ReturnCreatorParam struct {
}

type ReturnCreator struct {
	param *ReturnCreatorParam
}

func (s *ReturnCreator) New() *Return {
	return &Return{
		seed: s,
	}
}

func (s *ReturnCreator) Type() string {
	return ReturnInterruptType
}

func NewReturnCreator(param *ReturnCreatorParam) *ReturnCreator {
	return &ReturnCreator{
		param: param,
	}
}
