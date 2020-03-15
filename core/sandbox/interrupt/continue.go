package interrupt

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	ContinueInterruptType = "continue"
)

type Continue struct {
	tag concept.String
}

func (e *Continue) InterruptType() string {
	return ContinueInterruptType
}

func (e *Continue) Tag() concept.String {
	return e.tag
}

func NewContinueWithTag(tag concept.String) *Continue {
	return &Continue{
		tag: tag,
	}
}

func NewContinue() *Continue {
	return &Continue{
		tag: nil,
	}
}
