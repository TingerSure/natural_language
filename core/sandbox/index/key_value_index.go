package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type KeyValueIndexSeed interface {
	ToLanguage(string, concept.Pool, *KeyValueIndex) (string, concept.Exception)
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
}

type KeyValueIndex struct {
	value concept.Pipe
	key   concept.String
	seed  KeyValueIndexSeed
}

const (
	IndexKeyValueType = "KeyValue"
)

func (f *KeyValueIndex) Key() concept.String {
	return f.key
}

func (f *KeyValueIndex) Value() concept.Pipe {
	return f.value
}

func (f *KeyValueIndex) Type() string {
	return f.seed.Type()
}

func (f *KeyValueIndex) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (s *KeyValueIndex) ToString(prefix string) string {
	return fmt.Sprintf("%v: %v", s.key.ToString(""), s.value.ToString(prefix))
}

func (s *KeyValueIndex) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "KeyValueIndex cannot be called.")
}

func (s *KeyValueIndex) CallAnticipate(space concept.Pool, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *KeyValueIndex) Get(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return s.value.Get(space)
}

func (s *KeyValueIndex) Anticipate(space concept.Pool) concept.Variable {
	return s.value.Anticipate(space)
}

func (s *KeyValueIndex) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "KeyValueIndex cannot be changed.")
}

type KeyValueIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
}

type KeyValueIndexCreator struct {
	Seeds map[string]func(concept.Pool, *KeyValueIndex) (string, concept.Exception)
	param *KeyValueIndexCreatorParam
}

func (s *KeyValueIndexCreator) New(key concept.String, value concept.Pipe) *KeyValueIndex {
	return &KeyValueIndex{
		key:   key,
		value: value,
		seed:  s,
	}
}

func (s *KeyValueIndexCreator) ToLanguage(language string, space concept.Pool, instance *KeyValueIndex) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
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
		Seeds: map[string]func(concept.Pool, *KeyValueIndex) (string, concept.Exception){},
		param: param,
	}
}
