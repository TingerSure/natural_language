package interrupt

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	ContinueInterruptType = "continue"
)

type ContinueSeed interface {
	Type() string
}

type Continue struct {
	tag  concept.String
	seed ContinueSeed
}

func (e *Continue) InterruptType() string {
	return e.seed.Type()
}

func (e *Continue) Tag() concept.String {
	return e.tag
}

type ContinueCreatorParam struct {
}

type ContinueCreator struct {
	param *ContinueCreatorParam
}

func (s *ContinueCreator) New(tag concept.String) *Continue {
	return &Continue{
		tag:  tag,
		seed: s,
	}
}

func (s *ContinueCreator) Type() string {
	return ContinueInterruptType
}

func NewContinueCreator(param *ContinueCreatorParam) *ContinueCreator {
	return &ContinueCreator{
		param: param,
	}
}
