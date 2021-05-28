package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type KeyKeyIndexSeed interface {
	ToLanguage(string, concept.Closure, *KeyKeyIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
}

type KeyKeyIndex struct {
	to   concept.String
	from concept.String
	seed KeyKeyIndexSeed
}

const (
	IndexKeyKeyType = "KeyKey"
)

func (f *KeyKeyIndex) From() concept.String {
	return f.from
}

func (f *KeyKeyIndex) To() concept.String {
	return f.to
}

func (f *KeyKeyIndex) Type() string {
	return f.seed.Type()
}

func (f *KeyKeyIndex) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (s *KeyKeyIndex) ToString(prefix string) string {
	return fmt.Sprintf("%v: %v", s.from.ToString(""), s.to.ToString(prefix))
}

func (s *KeyKeyIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "KeyKeyIndex cannot be called.")
}

func (s *KeyKeyIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *KeyKeyIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.to, nil
}

func (s *KeyKeyIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.to
}

func (s *KeyKeyIndex) Set(space concept.Closure, to concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "KeyKeyIndex cannot be changed.")
}

type KeyKeyIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
}

type KeyKeyIndexCreator struct {
	Seeds map[string]func(string, concept.Closure, *KeyKeyIndex) string
	param *KeyKeyIndexCreatorParam
}

func (s *KeyKeyIndexCreator) New(from concept.String, to concept.String) *KeyKeyIndex {
	return &KeyKeyIndex{
		from: from,
		to:   to,
		seed: s,
	}
}

func (s *KeyKeyIndexCreator) ToLanguage(language string, space concept.Closure, instance *KeyKeyIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *KeyKeyIndexCreator) Type() string {
	return IndexKeyKeyType
}

func (s *KeyKeyIndexCreator) NewException(key string, message string) concept.Exception {
	return s.param.ExceptionCreator(key, message)
}

func (s *KeyKeyIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func NewKeyKeyIndexCreator(param *KeyKeyIndexCreatorParam) *KeyKeyIndexCreator {
	return &KeyKeyIndexCreator{
		Seeds: map[string]func(string, concept.Closure, *KeyKeyIndex) string{},
		param: param,
	}
}
