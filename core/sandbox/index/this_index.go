package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	thisIndexKey = "this"
)

type ThisIndexSeed interface {
	ToLanguage(string, *ThisIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewString(string) concept.String
}

type ThisIndex struct {
	seed ThisIndexSeed
}

const (
	IndexThisType = "This"
)

func (f *ThisIndex) Type() string {
	return f.seed.Type()
}

func (f *ThisIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ThisIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ThisIndex) ToString(prefix string) string {
	return thisIndexKey
}

func (s *ThisIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := space.PeekBubble(s.seed.NewString(thisIndexKey))
	return value
}

func (s *ThisIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(s.seed.NewString(thisIndexKey))
}

func (s *ThisIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "This cannot be changed.")

}

type ThisIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	StringCreator    func(string) concept.String
}

type ThisIndexCreator struct {
	Seeds map[string]func(string, *ThisIndex) string
	param *ThisIndexCreatorParam
}

func (s *ThisIndexCreator) New() *ThisIndex {
	return &ThisIndex{
		seed: s,
	}
}

func (s *ThisIndexCreator) ToLanguage(language string, instance *ThisIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ThisIndexCreator) Type() string {
	return IndexThisType
}

func (s *ThisIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ThisIndexCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func NewThisIndexCreator(param *ThisIndexCreatorParam) *ThisIndexCreator {
	return &ThisIndexCreator{
		Seeds: map[string]func(string, *ThisIndex) string{},
		param: param,
	}
}
