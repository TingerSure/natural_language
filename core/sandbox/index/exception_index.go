package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ExceptionIndexSeed interface {
	ToLanguage(string, concept.Pool, *ExceptionIndex) (string, concept.Exception)
	Type() string
	NewParam() concept.Param
	NewNull() concept.Null
}

type ExceptionIndex struct {
	value concept.Exception
	seed  ExceptionIndexSeed
}

const (
	IndexExceptionType = "Exception"
)

func (f *ExceptionIndex) Type() string {
	return f.seed.Type()
}

func (f *ExceptionIndex) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (s *ExceptionIndex) ToString(prefix string) string {
	return s.value.ToString(prefix)
}

func (s *ExceptionIndex) Value() concept.Exception {
	return s.value
}

func (s *ExceptionIndex) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.value
}

func (s *ExceptionIndex) Get(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return nil, s.value
}

func (s *ExceptionIndex) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return s.value
}

type ExceptionIndexCreatorParam struct {
	ParamCreator func() concept.Param
	NullCreator  func() concept.Null
}

type ExceptionIndexCreator struct {
	Seeds map[string]func(concept.Pool, *ExceptionIndex) (string, concept.Exception)
	param *ExceptionIndexCreatorParam
}

func (s *ExceptionIndexCreator) New(value concept.Exception) *ExceptionIndex {
	return &ExceptionIndex{
		value: value,
		seed:  s,
	}
}
func (s *ExceptionIndexCreator) ToLanguage(language string, space concept.Pool, instance *ExceptionIndex) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *ExceptionIndexCreator) Type() string {
	return IndexExceptionType
}

func (s *ExceptionIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *ExceptionIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewExceptionIndexCreator(param *ExceptionIndexCreatorParam) *ExceptionIndexCreator {
	return &ExceptionIndexCreator{
		Seeds: map[string]func(concept.Pool, *ExceptionIndex) (string, concept.Exception){},
		param: param,
	}
}
