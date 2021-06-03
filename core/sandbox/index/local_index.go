package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type LocalIndexSeed interface {
	ToLanguage(string, concept.Pool, *LocalIndex) (string, concept.Exception)
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type LocalIndex struct {
	key  concept.String
	seed LocalIndexSeed
}

const (
	IndexLocalType = "Local"
)

func (f *LocalIndex) Type() string {
	return f.seed.Type()
}

func (f *LocalIndex) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (s *LocalIndex) ToString(prefix string) string {
	return s.key.ToString(prefix)
}

func (s *LocalIndex) Key() concept.String {
	return s.key
}

func (s *LocalIndex) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *LocalIndex) CallAnticipate(space concept.Pool, param concept.Param) concept.Param {
	funcs := s.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (s *LocalIndex) Anticipate(space concept.Pool) concept.Variable {
	value, _ := space.PeekLocal(s.key)
	return value
}

func (s *LocalIndex) Get(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return space.GetLocal(s.key)
}

func (s *LocalIndex) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return space.SetLocal(s.key, value)
}

type LocalIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type LocalIndexCreator struct {
	Seeds map[string]func(concept.Pool, *LocalIndex) (string, concept.Exception)
	param *LocalIndexCreatorParam
}

func (s *LocalIndexCreator) New(key concept.String) *LocalIndex {
	return &LocalIndex{
		key:  key,
		seed: s,
	}
}

func (s *LocalIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *LocalIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *LocalIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *LocalIndexCreator) ToLanguage(language string, space concept.Pool, instance *LocalIndex) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *LocalIndexCreator) Type() string {
	return IndexLocalType
}

func NewLocalIndexCreator(param *LocalIndexCreatorParam) *LocalIndexCreator {
	return &LocalIndexCreator{
		Seeds: map[string]func(concept.Pool, *LocalIndex) (string, concept.Exception){},
		param: param,
	}
}
