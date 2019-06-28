package auxiliary

import (
	"github.com/TingerSure/natural_language/word"
)

const (
	belongAuxiliaryName string = "system.auxiliary.belong"
	belongType          int    = word.AuxiliaryBelong
)

const (
	BelongTo string = "的"
)

type Belong struct {
}

func (p *Belong) GetName() string {
	return belongAuxiliaryName
}

func (p *Belong) GetWords(firstCharacter string) []*word.Word {
	return word.WordsFilter([]*word.Word{
		word.NewWord(BelongTo, belongType),
	}, firstCharacter)
}

func NewBelong() *Belong {
	return (&Belong{})
}
