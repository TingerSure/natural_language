package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type RequireIndexSeed interface {
	ToLanguage(string, concept.Pool, *RequireIndex) (string, concept.Exception)
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type RequireIndex struct {
	originator concept.Pipe
	name       string
	seed       RequireIndexSeed
}

const (
	IndexRequireType = "Require"
)

func (f *RequireIndex) Name() string {
	return f.name
}

func (f *RequireIndex) Originator() concept.Pipe {
	return f.originator
}

func (f *RequireIndex) Type() string {
	return f.seed.Type()
}

func (f *RequireIndex) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (s *RequireIndex) ToString(prefix string) string {
	return fmt.Sprintf("require %v = %v", s.name, s.originator.ToString(prefix))
}

func (s *RequireIndex) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "RequireIndex cannot be called.")

}

func (s *RequireIndex) CallAnticipate(space concept.Pool, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *RequireIndex) Get(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return s.originator.Get(space)
}

func (s *RequireIndex) Anticipate(space concept.Pool) concept.Variable {
	return s.originator.Anticipate(space)
}

func (s *RequireIndex) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "RequireIndex cannot be changed.")
}

type RequireIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type RequireIndexCreator struct {
	Seeds map[string]func(concept.Pool, *RequireIndex) (string, concept.Exception)
	param *RequireIndexCreatorParam
}

func (s *RequireIndexCreator) New(name string, originator concept.Pipe) *RequireIndex {
	return &RequireIndex{
		name:       name,
		originator: originator,
		seed:       s,
	}
}

func (s *RequireIndexCreator) ToLanguage(language string, space concept.Pool, instance *RequireIndex) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
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
		Seeds: map[string]func(concept.Pool, *RequireIndex) (string, concept.Exception){},
		param: param,
	}
}
