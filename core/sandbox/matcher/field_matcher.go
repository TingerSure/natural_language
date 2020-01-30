package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type FieldMatcher struct {
	fieldName string
}

func (c *FieldMatcher) ToString(string) string {
	return fmt.Sprintf("field=%v", c.fieldName)
}

func (c *FieldMatcher) Match(value concept.Variable) bool {
	object, ok := variable.VariableFamilyInstance.IsObjectHome(value)
	if !ok {
		return false
	}
	return object.HasField(c.fieldName)
}

func NewFieldMatcher(fieldName string) *FieldMatcher {
	return &FieldMatcher{
		fieldName: fieldName,
	}
}
