package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ClassMatcher struct {
	class concept.Class
}

var (
	ClassMatcherLanguageSeeds = map[string]func(string, *ClassMatcher) string{}
)

func (f *ClassMatcher) ToLanguage(language string) string {
	seed := ClassMatcherLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (c *ClassMatcher) ToString(string) string {
	return fmt.Sprintf("class=%v", c.class.ToString(""))
}

func (c *ClassMatcher) Match(value concept.Variable) bool {
	return value.GetClass() == c.class
}

func NewClassMatcher(class concept.Class) *ClassMatcher {
	return &ClassMatcher{
		class: class,
	}
}
