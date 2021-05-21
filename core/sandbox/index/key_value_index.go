package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type KeyValueIndexSeed interface {
	ToLanguage(string, *KeyValueIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
}

type KeyValueIndex struct {
	value concept.Index
	key   concept.String
	seed  KeyValueIndexSeed
}

const (
	IndexKeyValueType = "KeyValue"
)

func (f *KeyValueIndex) Key() concept.String {
	return f.key
}

func (f *KeyValueIndex) Value() concept.Index {
	return f.value
}

func (f *KeyValueIndex) Type() string {
	return f.seed.Type()
}

func (f *KeyValueIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *KeyValueIndex) ToString(prefix string) string {
	return fmt.Sprintf("%v: %v", s.key.ToString(""), s.value.ToString(prefix))
}

func (s *KeyValueIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "KeyValueIndex cannot be called.")
}

func (s *KeyValueIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *KeyValueIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.value.Get(space)
}

func (s *KeyValueIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.value.Anticipate(space)
}

func (s *KeyValueIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "KeyValueIndex cannot be changed.")
}

type KeyValueIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
}

type KeyValueIndexCreator struct {
	Seeds map[string]func(string, *KeyValueIndex) string
	param *KeyValueIndexCreatorParam
}

func (s *KeyValueIndexCreator) New(key concept.String, value concept.Index) *KeyValueIndex {
	return &KeyValueIndex{
		key:   key,
		value: value,
		seed:  s,
	}
}

func (s *KeyValueIndexCreator) ToLanguage(language string, instance *KeyValueIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *KeyValueIndexCreator) Type() string {
	return IndexKeyValueType
}

func (s *KeyValueIndexCreator) NewException(key string, message string) concept.Exception {
	return s.param.ExceptionCreator(key, message)
}

func (s *KeyValueIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func NewKeyValueIndexCreator(param *KeyValueIndexCreatorParam) *KeyValueIndexCreator {
	return &KeyValueIndexCreator{
		Seeds: map[string]func(string, *KeyValueIndex) string{},
		param: param,
	}
}
