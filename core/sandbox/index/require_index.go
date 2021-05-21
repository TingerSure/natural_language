package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type RequireIndexSeed interface {
	ToLanguage(string, *RequireIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type RequireIndex struct {
	originator concept.Index
	name       string
	seed       RequireIndexSeed
}

const (
	IndexRequireType = "Require"
)

func (f *RequireIndex) Name() string {
	return f.name
}

func (f *RequireIndex) Originator() concept.Index {
	return f.originator
}

func (f *RequireIndex) Type() string {
	return f.seed.Type()
}

func (f *RequireIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *RequireIndex) ToString(prefix string) string {
	return fmt.Sprintf("require %v = %v", s.name, s.originator.ToString(prefix))
}

func (s *RequireIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "RequireIndex cannot be called.")

}

func (s *RequireIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *RequireIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.originator.Get(space)
}

func (s *RequireIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.originator.Anticipate(space)
}

func (s *RequireIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "RequireIndex cannot be changed.")
}

type RequireIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type RequireIndexCreator struct {
	Seeds map[string]func(string, *RequireIndex) string
	param *RequireIndexCreatorParam
}

func (s *RequireIndexCreator) New(name string, originator concept.Index) *RequireIndex {
	return &RequireIndex{
		name:       name,
		originator: originator,
		seed:       s,
	}
}

func (s *RequireIndexCreator) ToLanguage(language string, instance *RequireIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *RequireIndexCreator) Type() string {
	return IndexRequireType
}

func (s *RequireIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *RequireIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *RequireIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewRequireIndexCreator(param *RequireIndexCreatorParam) *RequireIndexCreator {
	return &RequireIndexCreator{
		Seeds: map[string]func(string, *RequireIndex) string{},
		param: param,
	}
}
