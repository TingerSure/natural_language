package operator

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	additionName string = "system.operator.addition"
	additionType int    = word_types.Operator
)

var (
	AdditionCharactor = "+"

	additionWords []*tree.Word = []*tree.Word{tree.NewWord(AdditionCharactor, additionType)}
)

type Addition struct {
}

func (p *Addition) GetName() string {
	return additionName
}

func (p *Addition) GetWords(firstCharacter string) []*tree.Word {
	return tree.WordsFilter(additionWords, firstCharacter)
}

func (p *Addition) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return nil
				//TODO
			}, treasure, phrase_types.Operator)
		}, p.GetName()),
	}
}

func (p *Addition) GetStructRules() []*tree.StructRule {
	return nil
}

func NewAddition() *Addition {
	return (&Addition{})
}
