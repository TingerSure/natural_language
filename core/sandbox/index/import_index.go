package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ImportIndexSeed interface {
	ToLanguage(string, *ImportIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type ImportIndex struct {
	page concept.Variable
	path string
	name string
	seed ImportIndexSeed
}

const (
	IndexImportType = "Import"
)

func (f *ImportIndex) Page() concept.Variable {
	return f.page
}

func (f *ImportIndex) Name() string {
	return f.name
}

func (f *ImportIndex) Path() string {
	return f.path
}

func (f *ImportIndex) Type() string {
	return f.seed.Type()
}

func (f *ImportIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ImportIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ImportIndex) ToString(prefix string) string {
	return fmt.Sprintf("import %v \"%v\"", s.name, s.path)
}

func (s *ImportIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	return nil, s.seed.NewException("runtime error", "Import cannot be called.")

}

func (s *ImportIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	return s.seed.NewParam()
}

func (s *ImportIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.page, nil
}

func (s *ImportIndex) Anticipate(space concept.Closure) concept.Variable {
	return s.page
}

func (s *ImportIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("runtime error", "Import cannot be changed.")
}

type ImportIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
}

type ImportIndexCreator struct {
	Seeds map[string]func(string, *ImportIndex) string
	param *ImportIndexCreatorParam
}

func (s *ImportIndexCreator) New(name string, path string, page concept.Variable) *ImportIndex {
	return &ImportIndex{
		name: name,
		path: path,
		page: page,
		seed: s,
	}
}

func (s *ImportIndexCreator) ToLanguage(language string, instance *ImportIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ImportIndexCreator) Type() string {
	return IndexImportType
}

func (s *ImportIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ImportIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *ImportIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewImportIndexCreator(param *ImportIndexCreatorParam) *ImportIndexCreator {
	return &ImportIndexCreator{
		Seeds: map[string]func(string, *ImportIndex) string{},
		param: param,
	}
}