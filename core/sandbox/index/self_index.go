package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	selfIndexKey = "self"
)

type SelfIndexSeed interface {
	ToLanguage(string, *SelfIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewString(string) concept.String
}

type SelfIndex struct {
	seed SelfIndexSeed
}

const (
	IndexSelfType = "Self"
)

func (f *SelfIndex) Type() string {
	return f.seed.Type()
}

func (f *SelfIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *SelfIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *SelfIndex) ToString(prefix string) string {
	return selfIndexKey
}

func (s *SelfIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := space.PeekBubble(s.seed.NewString(selfIndexKey))
	return value
}

func (s *SelfIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(s.seed.NewString(selfIndexKey))
}

func (s *SelfIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Self cannot be changed.")
}

type SelfIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	StringCreator    func(string) concept.String
}

type SelfIndexCreator struct {
	Seeds map[string]func(string, *SelfIndex) string
	param *SelfIndexCreatorParam
}

func (s *SelfIndexCreator) New() *SelfIndex {
	return &SelfIndex{
		seed: s,
	}
}

func (s *SelfIndexCreator) ToLanguage(language string, instance *SelfIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *SelfIndexCreator) Type() string {
	return IndexSelfType
}

func (s *SelfIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *SelfIndexCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func NewSelfIndexCreator(param *SelfIndexCreatorParam) *SelfIndexCreator {
	return &SelfIndexCreator{
		Seeds: map[string]func(string, *SelfIndex) string{},
		param: param,
	}
}
