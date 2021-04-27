package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ExportIndexSeed interface {
	ToLanguage(string, *ExportIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type ExportIndex struct {
	originator concept.Index
	name       string
	seed       ExportIndexSeed
}

const (
	IndexExportType = "Export"
)

func (f *ExportIndex) Name() string {
	return f.name
}

func (f *ExportIndex) Originator() concept.Index {
	return f.originator
}

func (f *ExportIndex) Type() string {
	return f.seed.Type()
}

func (f *ExportIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ExportIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ExportIndex) ToString(prefix string) string {
	return fmt.Sprintf("export %v %v", s.name, s.originator.ToString(prefix))
}

func (s *ExportIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "ExportIndex cannot be called.")

}

func (s *ExportIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *ExportIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.originator.Get(space)
}

func (s *ExportIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.originator.Anticipate(space)
}

func (s *ExportIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "ExportIndex cannot be changed.")
}

type ExportIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type ExportIndexCreator struct {
	Seeds map[string]func(string, *ExportIndex) string
	param *ExportIndexCreatorParam
}

func (s *ExportIndexCreator) New(name string, originator concept.Index) *ExportIndex {
	return &ExportIndex{
		name:       name,
		originator: originator,
		seed:       s,
	}
}

func (s *ExportIndexCreator) ToLanguage(language string, instance *ExportIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ExportIndexCreator) Type() string {
	return IndexExportType
}

func (s *ExportIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ExportIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *ExportIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewExportIndexCreator(param *ExportIndexCreatorParam) *ExportIndexCreator {
	return &ExportIndexCreator{
		Seeds: map[string]func(string, *ExportIndex) string{},
		param: param,
	}
}
