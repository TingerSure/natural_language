package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type SystemMatcher struct {
	match func(concept.Variable) bool
}

var (
	SystemMatcherLanguageSeeds = map[string]func(string, *SystemMatcher) string{}
)

func (f *SystemMatcher) ToLanguage(language string) string {
	seed := SystemMatcherLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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
