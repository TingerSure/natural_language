package matcher

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
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
	field, exception := value.GetField(c.fieldName)
	return nl_interface.IsNil(exception) && !nl_interface.IsNil(field) && !field.IsFunction()
}

func NewFieldMatcher(fieldName concept.String) *FieldMatcher {
	return &FieldMatcher{
		fieldName: fieldName,
	}
}
