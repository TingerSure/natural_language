package interrupt

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	BreakInterruptType = "break"
)

type Break struct {
	tag concept.String
}

func (e *Break) InterruptType() string {
	return BreakInterruptType
}

func (e *Break) Tag() concept.String {
	return e.tag
}

func NewBreak() *Break {
	return &Break{
		tag: nil,
	}
}

func NewBreakWithTag(tag concept.String) *Break {
	return &Break{
		tag: tag,
	}
}
