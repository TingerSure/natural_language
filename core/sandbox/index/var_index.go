package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type VarIndexSeed interface {
	ToLanguage(string, *VarIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type VarIndex struct {
	originator concept.Index
	name       string
	seed       VarIndexSeed
}

const (
	IndexVarType = "Var"
)

func (f *VarIndex) Name() string {
	return f.name
}

func (f *VarIndex) Originator() concept.Index {
	return f.originator
}

func (f *VarIndex) Type() string {
	return f.seed.Type()
}

func (f *VarIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *VarIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *VarIndex) ToString(prefix string) string {
	return fmt.Sprintf("var %v %v", s.name, s.originator.ToString(prefix))
}

func (s *VarIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "VarIndex cannot be called.")
}

func (s *VarIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *VarIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.originator.Get(space)
}

func (s *VarIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.originator.Anticipate(space)
}

func (s *VarIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "VarIndex cannot be changed.")
}

type VarIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type VarIndexCreator struct {
	Seeds map[string]func(string, *VarIndex) string
	param *VarIndexCreatorParam
}

func (s *VarIndexCreator) New(name string, originator concept.Index) *VarIndex {
	return &VarIndex{
		name:       name,
		originator: originator,
		seed:       s,
	}
}

func (s *VarIndexCreator) ToLanguage(language string, instance *VarIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *VarIndexCreator) Type() string {
	return IndexVarType
}

func (s *VarIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *VarIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *VarIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewVarIndexCreator(param *VarIndexCreatorParam) *VarIndexCreator {
	return &VarIndexCreator{
		Seeds: map[string]func(string, *VarIndex) string{},
		param: param,
	}
}
