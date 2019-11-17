package operator

import (
	"github.com/TingerSure/natural_language/library/operator"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	MultiplicationName      string = "word.operator.multiplication"
	multiplicationType      int    = word_types.Operator
	multiplicationCharactor        = "*"
)

var (
	multiplicationWords []*tree.Word = []*tree.Word{tree.NewWord(multiplicationCharactor, multiplicationType)}
)

type Multiplication struct {
	adaptor.SourceAdaptor
}

func (p *Multiplication) GetName() string {
	return MultiplicationName
}

func (p *Multiplication) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(multiplicationWords, sentence)
}

func (p *Multiplication) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return index.NewConstIndex(operator.MultiplicationFunc)
			}, treasure, phrase_types.Operator, p.GetName())
		}, p.GetName()),
	}
}

func NewMultiplication() *Multiplication {
	return (&Multiplication{})
}
