package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type SearchIndex struct {
	items []concept.Matcher
}

const (
	IndexSearchType = "Search"
)

func (f *SearchIndex) Type() string {
	return IndexSearchType
}

var (
	SearchIndexLanguageSeeds = map[string]func(string, *SearchIndex) string{}
)

func (f *SearchIndex) ToLanguage(language string) string {
	seed := SearchIndexLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (s *SearchIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *SearchIndex) ToString(prefix string) string {
	var subprefix = fmt.Sprintf("%v\t", prefix)
	subs := []string{}
	for _, item := range s.items {
		subs = append(subs, item.ToString(subprefix))
	}
	return fmt.Sprintf("search<%v>", strings.Join(subs, ", "))
}

func (s *SearchIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	var selected concept.Variable = nil
	space.IterateHistory(func(_ concept.String, value concept.Variable) bool {
		for _, item := range s.items {
			if !item.Match(value) {
				return false
			}
		}
		selected = value
		return true
	})
	if nl_interface.IsNil(selected) {
		selected = variable.NewNull()
	}
	return selected, nil
}

func (s *SearchIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException(variable.NewString("read only"), variable.NewString("Search index cannot be changed."))
}

func NewSearchIndex(items []concept.Matcher) *SearchIndex {
	return &SearchIndex{
		items: items,
	}
}
