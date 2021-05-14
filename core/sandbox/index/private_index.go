package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type PrivateIndexSeed interface {
	ToLanguage(string, *PrivateIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type PrivateIndex struct {
	originator concept.Index
	name       string
	seed       PrivateIndexSeed
}

const (
	IndexPrivateType = "Private"
)

func (f *PrivateIndex) Name() string {
	return f.name
}

func (f *PrivateIndex) Originator() concept.Index {
	return f.originator
}

func (f *PrivateIndex) Type() string {
	return f.seed.Type()
}

func (f *PrivateIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *PrivateIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *PrivateIndex) ToString(prefix string) string {
	return fmt.Sprintf("private %v = %v", s.name, s.originator.ToString(prefix))
}

func (s *PrivateIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "PrivateIndex cannot be called.")
}

func (s *PrivateIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *PrivateIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.originator.Get(space)
}

func (s *PrivateIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.originator.Anticipate(space)
}

func (s *PrivateIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "PrivateIndex cannot be changed.")
}

type PrivateIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type PrivateIndexCreator struct {
	Seeds map[string]func(string, *PrivateIndex) string
	param *PrivateIndexCreatorParam
}

func (s *PrivateIndexCreator) New(name string, originator concept.Index) *PrivateIndex {
	return &PrivateIndex{
		name:       name,
		originator: originator,
		seed:       s,
	}
}

func (s *PrivateIndexCreator) ToLanguage(language string, instance *PrivateIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *PrivateIndexCreator) Type() string {
	return IndexPrivateType
}

func (s *PrivateIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *PrivateIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *PrivateIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewPrivateIndexCreator(param *PrivateIndexCreatorParam) *PrivateIndexCreator {
	return &PrivateIndexCreator{
		Seeds: map[string]func(string, *PrivateIndex) string{},
		param: param,
	}
}
