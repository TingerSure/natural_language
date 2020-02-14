package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
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
	libs     *tree.LibraryManager
	operator concept.Function
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

func NewAddition(libs *tree.LibraryManager) *Addition {
	return (&Addition{
		libs:     libs,
		operator: libs.GetLibraryPage("system", "operator").GetFunction("AdditionFunc"),
	})
}
