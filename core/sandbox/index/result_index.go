package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type ResaultIndex struct {
	items []concept.Matcher
}

const (
	IndexResaultType = "Resault"
)

func (f *ResaultIndex) Type() string {
	return IndexResaultType
}

var (
	ResaultIndexLanguageSeeds = map[string]func(string, *ResaultIndex) string{}
)

func (f *ResaultIndex) ToLanguage(language string) string {
	seed := ResaultIndexLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (s *ResaultIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ResaultIndex) ToString(prefix string) string {
	var subprefix = fmt.Sprintf("%v\t", prefix)
	subs := []string{}
	for _, item := range s.items {
		subs = append(subs, item.ToString(subprefix))
	}
	return fmt.Sprintf("result<%v>", strings.Join(subs, ", "))
}

func (s *ResaultIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	var selected concept.Variable = nil
	space.IterateExtempore(func(line concept.Index, value concept.Variable) bool {
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

func (s *ResaultIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException(variable.NewString("read only"), variable.NewString("Resault index cannot be changed."))
}

func NewResaultIndex(items []concept.Matcher) *ResaultIndex {
	return &ResaultIndex{
		items: items,
	}
}
