package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ProvideIndexSeed interface {
	ToLanguage(string, *ProvideIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type ProvideIndex struct {
	originator concept.Index
	name       string
	seed       ProvideIndexSeed
}

const (
	IndexProvideType = "Provide"
)

func (f *ProvideIndex) Name() string {
	return f.name
}

func (f *ProvideIndex) Originator() concept.Index {
	return f.originator
}

func (f *ProvideIndex) Type() string {
	return f.seed.Type()
}

func (f *ProvideIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ProvideIndex) ToString(prefix string) string {
	return fmt.Sprintf("provide %v = %v", s.name, s.originator.ToString(prefix))
}

func (s *ProvideIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "ProvideIndex cannot be called.")

}

func (s *ProvideIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *ProvideIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.originator.Get(space)
}

func (s *ProvideIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.originator.Anticipate(space)
}

func (s *ProvideIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "ProvideIndex cannot be changed.")
}

type ProvideIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type ProvideIndexCreator struct {
	Seeds map[string]func(string, *ProvideIndex) string
	param *ProvideIndexCreatorParam
}

func (s *ProvideIndexCreator) New(name string, originator concept.Index) *ProvideIndex {
	return &ProvideIndex{
		name:       name,
		originator: originator,
		seed:       s,
	}
}

func (s *ProvideIndexCreator) ToLanguage(language string, instance *ProvideIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ProvideIndexCreator) Type() string {
	return IndexProvideType
}

func (s *ProvideIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ProvideIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *ProvideIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewProvideIndexCreator(param *ProvideIndexCreatorParam) *ProvideIndexCreator {
	return &ProvideIndexCreator{
		Seeds: map[string]func(string, *ProvideIndex) string{},
		param: param,
	}
}
