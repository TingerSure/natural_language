package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type ResaultIndexSeed interface {
	ToLanguage(string, *ResaultIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewNull() concept.Null
}

type ResaultIndex struct {
	items []concept.Matcher
	seed  ResaultIndexSeed
}

const (
	IndexResaultType = "Resault"
)

func (f *ResaultIndex) Type() string {
	return f.seed.Type()
}

func (f *ResaultIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
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
		selected = s.seed.NewNull()
	}
	return selected, nil
}

func (s *ResaultIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Resault index cannot be changed.")
}

type ResaultIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	NullCreator      func() concept.Null
}

type ResaultIndexCreator struct {
	Seeds map[string]func(string, *ResaultIndex) string
	param *ResaultIndexCreatorParam
}

func (s *ResaultIndexCreator) New(items []concept.Matcher) *ResaultIndex {
	return &ResaultIndex{
		items: items,
		seed:  s,
	}
}
func (s *ResaultIndexCreator) ToLanguage(language string, instance *ResaultIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ResaultIndexCreator) Type() string {
	return IndexResaultType
}

func (s *ResaultIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ResaultIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewResaultIndexCreator(param *ResaultIndexCreatorParam) *ResaultIndexCreator {
	return &ResaultIndexCreator{
		Seeds: map[string]func(string, *ResaultIndex) string{},
		param: param,
	}
}
