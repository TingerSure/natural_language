package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type SearchIndexSeed interface {
	ToLanguage(string, *SearchIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type SearchIndex struct {
	items []concept.Matcher
	seed  SearchIndexSeed
}

const (
	IndexSearchType = "Search"
)

func (f *SearchIndex) Type() string {
	return f.seed.Type()
}

func (f *SearchIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
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

func (s *SearchIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *SearchIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	funcs := s.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (s *SearchIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := s.Get(space)
	return value
}

func (s *SearchIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	var selected concept.Variable = s.seed.NewNull()
	space.IterateHistory(func(_ concept.String, value concept.Variable) bool {
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
	return s.seed.NewException("read only", "Search index cannot be changed.")
}

type SearchIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type SearchIndexCreator struct {
	Seeds map[string]func(string, *SearchIndex) string
	param *SearchIndexCreatorParam
}

func (s *SearchIndexCreator) New(items []concept.Matcher) *SearchIndex {
	return &SearchIndex{
		items: items,
		seed:  s,
	}
}
func (s *SearchIndexCreator) ToLanguage(language string, instance *SearchIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *SearchIndexCreator) Type() string {
	return IndexSearchType
}

func (s *SearchIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *SearchIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *SearchIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewSearchIndexCreator(param *SearchIndexCreatorParam) *SearchIndexCreator {
	return &SearchIndexCreator{
		Seeds: map[string]func(string, *SearchIndex) string{},
		param: param,
	}
}
