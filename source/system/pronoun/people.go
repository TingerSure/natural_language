package pronoun

import (
	"github.com/TingerSure/natural_language/word"
)

const (
	peopleProNounName string = "system.pronoun.people"
	peopleType        int    = word.Pronoun
)

const (
	He  string = "他"
	She string = "她"
	It  string = "它"
	You string = "你"
	I   string = "我"
)

type People struct {
}

func (p *People) GetName() string {
	return peopleProNounName
}

func (p *People) GetWords(firstCharacter string) []*word.Word {
	return []*word.Word{
		word.NewWord(He, peopleType),
		word.NewWord(She, peopleType),
		word.NewWord(It, peopleType),
		word.NewWord(You, peopleType),
		word.NewWord(I, peopleType),
	}
}

func NewPeople() *People {
	return (&People{})
}
