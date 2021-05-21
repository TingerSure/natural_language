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
	NewParam() concept.Param
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

func (s *ResaultIndex) ToString(prefix string) string {
	var subprefix = fmt.Sprintf("%v\t", prefix)
	subs := []string{}
	for _, item := range s.items {
		subs = append(subs, item.ToString(subprefix))
	}
	return fmt.Sprintf("result<%v>", strings.Join(subs, ", "))
}

func (s *ResaultIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *ResaultIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	funcs := s.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (s *ResaultIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := s.Get(space)
	return value
}

func (s *ResaultIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	var selected concept.Variable = s.seed.NewNull()
	space.IterateExtempore(func(line concept.Index, value concept.Variable) bool {
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

func (s *ResaultIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Resault index cannot be changed.")
}

type ResaultIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
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

func (s *ResaultIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
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
