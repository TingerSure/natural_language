package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type MethodMatcher struct {
	methodName concept.String
}

var (
	MethodMatcherLanguageSeeds = map[string]func(string, *MethodMatcher) string{}
)

func (f *MethodMatcher) ToLanguage(language string) string {
	seed := MethodMatcherLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (c *MethodMatcher) ToString(prefix string) string {
	return fmt.Sprintf("method=%v", c.methodName.ToString(prefix))
}

func (c *MethodMatcher) Match(value concept.Variable) bool {
	method, exception := value.GetField(c.methodName)
	return nl_interface.IsNil(exception) && !nl_interface.IsNil(method) && method.IsFunction()
}

func NewMethodMatcher(methodName concept.String) *MethodMatcher {
	return &MethodMatcher{
		methodName: methodName,
	}
}
