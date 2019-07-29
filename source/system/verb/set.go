package verb

import (
	"github.com/TingerSure/natural_language/tree"
)

const (
	Is = "是"
)

const (
	setName string = "system.verb.set"
	setType int    = tree.Verb
)

type Set struct {
}

func (s *Set) GetName() string {
	return setName
}

func (s *Set) GetWords(firstCharacter string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(Is, setType),
	}, firstCharacter)
}
func (p *Set) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{}
}

func (p *Set) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{}
}
func NewSet() *Set {
	return (&Set{})
}
