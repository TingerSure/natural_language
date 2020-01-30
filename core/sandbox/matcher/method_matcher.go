package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type MethodMatcher struct {
	methodName string
}

func (c *MethodMatcher) ToString(string) string {
	return fmt.Sprintf("method=%v", c.methodName)
}

func (c *MethodMatcher) Match(value concept.Variable) bool {
	object, ok := variable.VariableFamilyInstance.IsObjectHome(value)
	if !ok {
		return false
	}
	return object.HasMethod(c.methodName)
}

func NewMethodMatcher(methodName string) *MethodMatcher {
	return &MethodMatcher{
		methodName: methodName,
	}
}
