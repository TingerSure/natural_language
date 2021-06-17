package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ConstIndexSeed interface {
	ToLanguage(string, concept.Pool, *ConstIndex) (string, concept.Exception)
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type ConstIndex struct {
	value concept.Variable
	line  concept.Line
	seed  ConstIndexSeed
}

const (
	IndexConstType = "Const"
)

func (f *ConstIndex) SetLine(line concept.Line) {
	f.line = line
}

func (f *ConstIndex) Type() string {
	return f.seed.Type()
}

func (f *ConstIndex) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (s *ConstIndex) ToString(prefix string) string {
	return s.value.ToString(prefix)
}

func (s *ConstIndex) Value() concept.Variable {
	return s.value
}

func (s *ConstIndex) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	if !s.value.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString(""))).AddExceptionLine(s.line)
	}
	return s.value.(concept.Function).Exec(param, nil)
}

func (s *ConstIndex) CallAnticipate(space concept.Pool, param concept.Param) concept.Param {
	if !s.value.IsFunction() {
		return s.seed.NewParam()
	}
	return s.value.(concept.Function).Anticipate(param, nil)
}

func (s *ConstIndex) Get(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return s.value, nil
}

func (s *ConstIndex) Anticipate(space concept.Pool) concept.Variable {
	return s.value
}

func (s *ConstIndex) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Constants cannot be changed.").AddExceptionLine(s.line)
}

type ConstIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type ConstIndexCreator struct {
	Seeds map[string]func(concept.Pool, *ConstIndex) (string, concept.Exception)
	param *ConstIndexCreatorParam
}

func (s *ConstIndexCreator) New(value concept.Variable) *ConstIndex {
	return &ConstIndex{
		value: value,
		seed:  s,
	}
}
func (s *ConstIndexCreator) ToLanguage(language string, space concept.Pool, instance *ConstIndex) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *ConstIndexCreator) Type() string {
	return IndexConstType
}

func (s *ConstIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ConstIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *ConstIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewConstIndexCreator(param *ConstIndexCreatorParam) *ConstIndexCreator {
	return &ConstIndexCreator{
		Seeds: map[string]func(concept.Pool, *ConstIndex) (string, concept.Exception){},
		param: param,
	}
}
