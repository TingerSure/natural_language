package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/source/phrase_type"

	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/library/operator"
)

const (
	AdditionName string = "word.operator.addition"

	additionCharactor = "+"
)

var (
	additionWords []*tree.Word = []*tree.Word{tree.NewWord(additionCharactor)}
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
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(operator.AdditionFunc)
					},
					Content: treasure,
					Types:   phrase_type.Operator,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewAddition() *Addition {
	return (&Addition{})
}
