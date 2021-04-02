package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ImportIndexSeed interface {
	ToLanguage(string, *ImportIndex) string
	Type() string
	NewException(string, string) concept.Exception
}

type ImportIndex struct {
	page concept.Index
	path string
	name string
	seed ImportIndexSeed
}

const (
	IndexImportType = "Import"
)

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
	return fmt.Sprintf("import %v \"%v\";", s.name, s.path)
}

func (s *ImportIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, nil
}

func (s *ImportIndex) Anticipate(space concept.Closure) concept.Variable {
	return nil
}

func (s *ImportIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Import cannot be changed.")
}

type ImportIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
}

type ImportIndexCreator struct {
	Seeds map[string]func(string, *ImportIndex) string
	param *ImportIndexCreatorParam
}

func (s *ImportIndexCreator) New(name string, path string, page concept.Index) *ImportIndex {
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

func NewImportIndexCreator(param *ImportIndexCreatorParam) *ImportIndexCreator {
	return &ImportIndexCreator{
		Seeds: map[string]func(string, *ImportIndex) string{},
		param: param,
	}
}
