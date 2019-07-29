package auxiliary

import (
	"github.com/TingerSure/natural_language/tree"
)

const (
	belongAuxiliaryName string = "system.auxiliary.belong"
	belongType          int    = tree.AuxiliaryBelong
)

const (
	BelongTo string = "的"
)

type Belong struct {
}

func (p *Belong) GetName() string {
	return belongAuxiliaryName
}

func (p *Belong) GetWords(firstCharacter string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(BelongTo, belongType),
	}, firstCharacter)
}

func (p *Belong) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{}
}

func (p *Belong) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{}
}

func NewBelong() *Belong {
	return (&Belong{})
}
