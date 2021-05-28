package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type BubbleIndexSeed interface {
	ToLanguage(string, concept.Closure, *BubbleIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type BubbleIndex struct {
	key  concept.String
	seed BubbleIndexSeed
}

const (
	IndexBubbleType = "Bubble"
)

func (f *BubbleIndex) Type() string {
	return f.seed.Type()
}

func (f *BubbleIndex) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)

}

func (s *BubbleIndex) ToString(prefix string) string {
	return s.key.Value()
}

func (s *BubbleIndex) Key() concept.String {
	return s.key
}

func (s *BubbleIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *BubbleIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	funcs := s.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (s *BubbleIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(s.key)
}

func (s *BubbleIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := space.PeekBubble(s.key)
	return value
}

func (s *BubbleIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return space.SetBubble(s.key, value)
}

type BubbleIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type BubbleIndexCreator struct {
	Seeds map[string]func(string, concept.Closure, *BubbleIndex) string
	param *BubbleIndexCreatorParam
}

func (s *BubbleIndexCreator) New(key concept.String) *BubbleIndex {
	return &BubbleIndex{
		key:  key,
		seed: s,
	}
}

func (s *BubbleIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *BubbleIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *BubbleIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *BubbleIndexCreator) ToLanguage(language string, space concept.Closure, instance *BubbleIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *BubbleIndexCreator) Type() string {
	return IndexBubbleType
}

func NewBubbleIndexCreator(param *BubbleIndexCreatorParam) *BubbleIndexCreator {
	return &BubbleIndexCreator{
		Seeds: map[string]func(string, concept.Closure, *BubbleIndex) string{},
		param: param,
	}
}
