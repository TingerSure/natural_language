package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type LocalIndexSeed interface {
	ToLanguage(string, *LocalIndex) string
	Type() string
}

type LocalIndex struct {
	key  concept.String
	seed LocalIndexSeed
}

const (
	IndexLocalType = "Local"
)

func (f *LocalIndex) Type() string {
	return f.seed.Type()
}

func (f *LocalIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *LocalIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *LocalIndex) ToString(prefix string) string {
	return s.key.ToString(prefix)
}

func (s *LocalIndex) Key() concept.String {
	return s.key
}

func (s *LocalIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := space.PeekLocal(s.key)
	return value
}

func (s *LocalIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetLocal(s.key)
}

func (s *LocalIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return space.SetLocal(s.key, value)
}

type LocalIndexCreatorParam struct {
}

type LocalIndexCreator struct {
	Seeds map[string]func(string, *LocalIndex) string
	param *LocalIndexCreatorParam
}

func (s *LocalIndexCreator) New(key concept.String) *LocalIndex {
	return &LocalIndex{
		key:  key,
		seed: s,
	}
}

func (s *LocalIndexCreator) ToLanguage(language string, instance *LocalIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *LocalIndexCreator) Type() string {
	return IndexLocalType
}

func NewLocalIndexCreator(param *LocalIndexCreatorParam) *LocalIndexCreator {
	return &LocalIndexCreator{
		Seeds: map[string]func(string, *LocalIndex) string{},
		param: param,
	}
}
