package verb

import (
	"github.com/TingerSure/natural_language/word"
)

const (
	Is = "是"
)

const (
	setName string = "system.verb.set"
	setType int    = word.Verb
)

type Set struct {
}

func (s *Set) GetName() string {
	return setName
}

func (s *Set) GetWords(firstCharacter string) []*word.Word {
	return word.WordsFilter([]*word.Word{
		word.NewWord(Is, setType),
	}, firstCharacter)
}

func NewSet() *Set {
	return (&Set{})
}
