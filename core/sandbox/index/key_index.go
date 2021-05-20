package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type KeyIndexSeed interface {
	ToLanguage(string, *KeyIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type KeyIndex struct {
	key  concept.String
	seed KeyIndexSeed
}

const (
	IndexKeyType = "Key"
)

func (f *KeyIndex) Key() concept.String {
	return f.key
}

func (f *KeyIndex) Type() string {
	return f.seed.Type()
}

func (f *KeyIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *KeyIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *KeyIndex) ToString(prefix string) string {
	return fmt.Sprintf("%v", s.key.ToString(""))
}

func (s *KeyIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "KeyIndex cannot be called.")
}

func (s *KeyIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *KeyIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.seed.NewNull(), nil
}

func (s *KeyIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.seed.NewNull()
}

func (s *KeyIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "KeyIndex cannot be changed.")
}

type KeyIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type KeyIndexCreator struct {
	Seeds map[string]func(string, *KeyIndex) string
	param *KeyIndexCreatorParam
}

func (s *KeyIndexCreator) New(key concept.String) *KeyIndex {
	return &KeyIndex{
		key:  key,
		seed: s,
	}
}

func (s *KeyIndexCreator) ToLanguage(language string, instance *KeyIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *KeyIndexCreator) Type() string {
	return IndexKeyType
}

func (s *KeyIndexCreator) NewException(key string, message string) concept.Exception {
	return s.param.ExceptionCreator(key, message)
}

func (s *KeyIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *KeyIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewKeyIndexCreator(param *KeyIndexCreatorParam) *KeyIndexCreator {
	return &KeyIndexCreator{
		Seeds: map[string]func(string, *KeyIndex) string{},
		param: param,
	}
}