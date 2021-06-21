package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	selfIndexKey = "self"
)

type SelfIndexSeed interface {
	ToLanguage(string, concept.Pool, *SelfIndex) (string, concept.Exception)
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewString(string) concept.String
	NewNull() concept.Null
}

type SelfIndex struct {
	line concept.Line
	seed SelfIndexSeed
}

const (
	IndexSelfType = "Self"
)

func (f *SelfIndex) SetLine(line concept.Line) {
	f.line = line
}

func (f *SelfIndex) Type() string {
	return f.seed.Type()
}

func (f *SelfIndex) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (s *SelfIndex) ToString(prefix string) string {
	return selfIndexKey
}

func (s *SelfIndex) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *SelfIndex) CallAnticipate(space concept.Pool, param concept.Param) concept.Param {
	funcs := s.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (s *SelfIndex) Anticipate(space concept.Pool) concept.Variable {
	value, _ := space.PeekBubble(s.seed.NewString(selfIndexKey))
	return value
}

func (s *SelfIndex) Get(space concept.Pool) (concept.Variable, concept.Interrupt) {
	value, suspend := space.GetBubble(s.seed.NewString(selfIndexKey))
	if !nl_interface.IsNil(suspend) {
		return nil, suspend.AddLine(s.line)
	}
	return value, nil
}

func (s *SelfIndex) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Self cannot be changed.")
}

type SelfIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	StringCreator    func(string) concept.String
	NullCreator      func() concept.Null
}

type SelfIndexCreator struct {
	Seeds map[string]func(concept.Pool, *SelfIndex) (string, concept.Exception)
	param *SelfIndexCreatorParam
}

func (s *SelfIndexCreator) New() *SelfIndex {
	return &SelfIndex{
		seed: s,
	}
}

func (s *SelfIndexCreator) ToLanguage(language string, space concept.Pool, instance *SelfIndex) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *SelfIndexCreator) Type() string {
	return IndexSelfType
}

func (s *SelfIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *SelfIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *SelfIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *SelfIndexCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func NewSelfIndexCreator(param *SelfIndexCreatorParam) *SelfIndexCreator {
	return &SelfIndexCreator{
		Seeds: map[string]func(concept.Pool, *SelfIndex) (string, concept.Exception){},
		param: param,
	}
}
