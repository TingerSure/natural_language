package interrupt

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	BreakInterruptType = "break"
)

type BreakSeed interface {
	Type() string
}

type Break struct {
	tag  concept.String
	seed BreakSeed
}

func (e *Break) InterruptType() string {
	return e.seed.Type()
}

func (e *Break) Tag() concept.String {
	return e.tag
}

type BreakCreatorParam struct {
}

type BreakCreator struct {
	param *BreakCreatorParam
}

func (s *BreakCreator) New(tag concept.String) *Break {
	return &Break{
		tag:  tag,
		seed: s,
	}
}

func (s *BreakCreator) Type() string {
	return BreakInterruptType
}

func NewBreakCreator(param *BreakCreatorParam) *BreakCreator {
	return &BreakCreator{
		param: param,
	}
}
