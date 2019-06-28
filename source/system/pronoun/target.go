package pronoun

import (
	"github.com/TingerSure/natural_language/word"
)

const (
	targetPronounName string = "system.pronoun.target"
	targetType        int    = word.Pronoun
)

const (
	He  string = "他"
	She string = "她"
	It  string = "它"
	You string = "你"
	I   string = "我"
)

type Target struct {
}

func (p *Target) GetName() string {
	return targetPronounName
}

func (p *Target) GetWords(firstCharacter string) []*word.Word {
	return word.WordsFilter([]*word.Word{
		word.NewWord(He, targetType),
		word.NewWord(She, targetType),
		word.NewWord(It, targetType),
		word.NewWord(You, targetType),
		word.NewWord(I, targetType),
	}, firstCharacter)
}

func NewTarget() *Target {
	return (&Target{})
}
