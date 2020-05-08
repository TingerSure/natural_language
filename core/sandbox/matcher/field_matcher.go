package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type FieldMatcher struct {
	fieldName concept.String
}

var (
	FieldMatcherLanguageSeeds = map[string]func(string, *FieldMatcher) string{}
)

func (f *FieldMatcher) ToLanguage(language string) string {
	seed := FieldMatcherLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (c *FieldMatcher) ToString(prefix string) string {
	return fmt.Sprintf("field=%v", c.fieldName.ToString(prefix))
}

func (c *FieldMatcher) Match(value concept.Variable) bool {
	object, ok := variable.VariableFamilyInstance.IsObjectHome(value)
	if !ok {
		return false
	}
	return object.HasField(c.fieldName)
}

func NewFieldMatcher(fieldName concept.String) *FieldMatcher {
	return &FieldMatcher{
		fieldName: fieldName,
	}
}
