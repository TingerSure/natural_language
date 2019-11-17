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
	AdditionName      string = "word.operator.addition"
	additionType      int    = word_types.Operator
	additionCharactor        = "+"
)

var (
	additionWords []*tree.Word = []*tree.Word{tree.NewWord(additionCharactor, additionType)}
)

type Addition struct {
	adaptor.SourceAdaptor
}

func (p *Addition) GetName() string {
	return AdditionName
}

func (p *Addition) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(additionWords, sentence)
}

func (p *Addition) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
				Index: func() concept.Index {
					return index.NewConstIndex(operator.AdditionFunc)
				},
				Content: treasure,
				Types:   phrase_types.Operator,
				From:    p.GetName(),
			})
		}, p.GetName()),
	}
}

func NewAddition() *Addition {
	return (&Addition{})
}
