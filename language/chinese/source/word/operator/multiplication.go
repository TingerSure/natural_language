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
	MultiplicationName string = "word.operator.multiplication"

	multiplicationCharactor = "*"
)

var (
	multiplicationWords []*tree.Word = []*tree.Word{tree.NewWord(multiplicationCharactor)}
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
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(operator.MultiplicationFunc)
					},
					Content: treasure,
					Types:   phrase_type.Operator,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewMultiplication() *Multiplication {
	return (&Multiplication{})
}
