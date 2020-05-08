package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ClassMatcher struct {
	className string
	alias     string
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
	if c.alias == "" {
		return fmt.Sprintf("class=%v", c.className)
	}
	return fmt.Sprintf("class=%v(%v)", c.className, c.alias)
}

func (c *ClassMatcher) Match(value concept.Variable) bool {
	object, ok := variable.VariableFamilyInstance.IsObjectHome(value)
	if !ok {
		return false
	}
	return object.IsClassAlias(c.className, c.alias)
}

func NewClassMatcher(className string, alias string) *ClassMatcher {
	return &ClassMatcher{
		className: className,
		alias:     alias,
	}
}
