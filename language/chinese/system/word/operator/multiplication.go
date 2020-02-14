package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
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
	libs     *tree.LibraryManager
	operator concept.Function
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
						return index.NewConstIndex(p.operator)
					},
					Content: treasure,
					Types:   phrase_type.Operator,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewMultiplication(libs *tree.LibraryManager) *Multiplication {
	return (&Multiplication{
		libs:     libs,
		operator: libs.GetLibraryPage("system", "operator").GetFunction("MultiplicationFunc"),
	})
}
