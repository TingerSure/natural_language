package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	DivisionName string = "word.operator.division"

	divisionCharactor = "/"
)

var (
	divisionWords []*tree.Word = []*tree.Word{tree.NewWord(divisionCharactor)}
)

type Division struct {
	adaptor.SourceAdaptor
	libs     *tree.LibraryManager
	operator concept.Function
}

func (p *Division) GetName() string {
	return DivisionName
}

func (p *Division) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(divisionWords, sentence)
}

func (p *Division) GetVocabularyRules() []*tree.VocabularyRule {
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

func NewDivision(libs *tree.LibraryManager) *Division {
	return (&Division{
		libs:     libs,
		operator: libs.GetLibraryPage("system", "operator").GetFunction("DivisionFunc"),
	})
}
