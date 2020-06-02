package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ConstIndexSeed interface {
	ToLanguage(string, *ConstIndex) string
	Type() string
	NewException(string, string) concept.Exception
}

type ConstIndex struct {
	value concept.Variable
	seed  ConstIndexSeed
}

const (
	IndexConstType = "Const"
)

func (f *ConstIndex) Type() string {
	return f.seed.Type()
}

func (f *ConstIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ConstIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ConstIndex) ToString(prefix string) string {
	return s.value.ToString(prefix)
}

func (s *ConstIndex) Value() concept.Variable {
	return s.value
}

func (s *ConstIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.value, nil
}

func (s *ConstIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Constants cannot be changed.")
}

type ConstIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
}

type ConstIndexCreator struct {
	Seeds map[string]func(string, *ConstIndex) string
	param *ConstIndexCreatorParam
}

func (s *ConstIndexCreator) New(value concept.Variable) *ConstIndex {
	return &ConstIndex{
		value: value,
		seed:  s,
	}
}
func (s *ConstIndexCreator) ToLanguage(language string, instance *ConstIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ConstIndexCreator) Type() string {
	return IndexConstType
}

func (s *ConstIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewConstIndexCreator(param *ConstIndexCreatorParam) *ConstIndexCreator {
	return &ConstIndexCreator{
		Seeds: map[string]func(string, *ConstIndex) string{},
		param: param,
	}
}
