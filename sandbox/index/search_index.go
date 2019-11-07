package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"strings"
)

type SearchIndex struct {
	items []concept.Matcher
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
	space.IterateHistory(func(key string, value concept.Variable) bool {
		for _, item := range s.items {
			if !item.Match(value) {
				return false
			}
		}
		selected = value
		return true
	})
	return selected, nil
}

func (s *SearchIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException("read only", "Search index cannot be changed.")
}

func NewSearchIndex(items []concept.Matcher) *SearchIndex {
	return &SearchIndex{
		items: items,
	}
}