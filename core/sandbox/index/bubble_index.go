package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type BubbleIndexSeed interface {
	ToLanguage(string, *BubbleIndex) string
	Type() string
}

type BubbleIndex struct {
	key  concept.String
	seed BubbleIndexSeed
}

const (
	IndexBubbleType = "Bubble"
)

func (f *BubbleIndex) Type() string {
	return f.seed.Type()
}

func (f *BubbleIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)

}

func (s *BubbleIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *BubbleIndex) ToString(prefix string) string {
	return s.key.ToString(prefix)
}

func (s *BubbleIndex) Key() concept.String {
	return s.key
}

func (s *BubbleIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(s.key)
}

func (s *BubbleIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := space.PeekBubble(s.key)
	return value
}

func (s *BubbleIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return space.SetBubble(s.key, value)
}

type BubbleIndexCreatorParam struct {
}

type BubbleIndexCreator struct {
	Seeds map[string]func(string, *BubbleIndex) string
	param *BubbleIndexCreatorParam
}

func (s *BubbleIndexCreator) New(key concept.String) *BubbleIndex {
	return &BubbleIndex{
		key:  key,
		seed: s,
	}
}

func (s *BubbleIndexCreator) ToLanguage(language string, instance *BubbleIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *BubbleIndexCreator) Type() string {
	return IndexBubbleType
}

func NewBubbleIndexCreator(param *BubbleIndexCreatorParam) *BubbleIndexCreator {
	return &BubbleIndexCreator{
		Seeds: map[string]func(string, *BubbleIndex) string{},
		param: param,
	}
}
