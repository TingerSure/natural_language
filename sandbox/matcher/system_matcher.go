package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type SystemMatcher struct {
	match func(concept.Variable) bool
}

func (c *SystemMatcher) ToString(string) string {
	return fmt.Sprintf("system_matcher")
}

func (c *SystemMatcher) Match(value concept.Variable) bool {
	return c.match(value)
}

func NewSystemMatcher(match func(concept.Variable) bool) *SystemMatcher {
	return &SystemMatcher{
		match: match,
	}
}
