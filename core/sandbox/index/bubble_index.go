package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type BubbleIndexSeed interface {
	ToLanguage(string, concept.Pool, *BubbleIndex) (string, concept.Exception)
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type BubbleIndex struct {
	key  concept.String
	line concept.Line
	seed BubbleIndexSeed
}

const (
	IndexBubbleType = "Bubble"
)

func (f *BubbleIndex) SetLine(line concept.Line) {
	f.line = line
}

func (f *BubbleIndex) Type() string {
	return f.seed.Type()
}

func (f *BubbleIndex) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)

}

func (s *BubbleIndex) ToString(prefix string) string {
	return s.key.Value()
}

func (s *BubbleIndex) Key() concept.String {
	return s.key
}

func (s *BubbleIndex) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *BubbleIndex) Get(space concept.Pool) (concept.Variable, concept.Interrupt) {
	value, suspend := space.GetBubble(s.key)
	if !nl_interface.IsNil(suspend) {
		suspend.AddLine(s.line)
	}
	return value, suspend
}

func (s *BubbleIndex) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return space.SetBubble(s.key, value)
}

type BubbleIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type BubbleIndexCreator struct {
	Seeds map[string]func(concept.Pool, *BubbleIndex) (string, concept.Exception)
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

func (s *BubbleIndexCreator) ToLanguage(language string, space concept.Pool, instance *BubbleIndex) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *BubbleIndexCreator) Type() string {
	return IndexBubbleType
}

func NewBubbleIndexCreator(param *BubbleIndexCreatorParam) *BubbleIndexCreator {
	return &BubbleIndexCreator{
		Seeds: map[string]func(concept.Pool, *BubbleIndex) (string, concept.Exception){},
		param: param,
	}
}
