package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type PublicIndexSeed interface {
	ToLanguage(string, concept.Closure, *PublicIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type PublicIndex struct {
	originator concept.Index
	name       string
	seed       PublicIndexSeed
}

const (
	IndexPublicType = "Public"
)

func (f *PublicIndex) Name() string {
	return f.name
}

func (f *PublicIndex) Originator() concept.Index {
	return f.originator
}

func (f *PublicIndex) Type() string {
	return f.seed.Type()
}

func (f *PublicIndex) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (s *PublicIndex) ToString(prefix string) string {
	return fmt.Sprintf("public %v = %v", s.name, s.originator.ToString(prefix))
}

func (s *PublicIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "PublicIndex cannot be called.")

}

func (s *PublicIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *PublicIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.originator.Get(space)
}

func (s *PublicIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.originator.Anticipate(space)
}

func (s *PublicIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "PublicIndex cannot be changed.")
}

type PublicIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type PublicIndexCreator struct {
	Seeds map[string]func(string, concept.Closure, *PublicIndex) string
	param *PublicIndexCreatorParam
}

func (s *PublicIndexCreator) New(name string, originator concept.Index) *PublicIndex {
	return &PublicIndex{
		name:       name,
		originator: originator,
		seed:       s,
	}
}

func (s *PublicIndexCreator) ToLanguage(language string, space concept.Closure, instance *PublicIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *PublicIndexCreator) Type() string {
	return IndexPublicType
}

func (s *PublicIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *PublicIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *PublicIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewPublicIndexCreator(param *PublicIndexCreatorParam) *PublicIndexCreator {
	return &PublicIndexCreator{
		Seeds: map[string]func(string, concept.Closure, *PublicIndex) string{},
		param: param,
	}
}
