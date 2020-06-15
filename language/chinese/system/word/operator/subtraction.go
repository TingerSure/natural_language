package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	SubtractionName string = "word.operator.subtraction"

	subtractionCharactor = "-"
)

var (
	subtractionWords []*tree.Word = []*tree.Word{tree.NewWord(subtractionCharactor)}
)

func init() {

}

type Subtraction struct {
	*adaptor.SourceAdaptor
	operator concept.Function
}

func (p *Subtraction) GetName() string {
	return SubtractionName
}

func (p *Subtraction) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(subtractionWords, sentence)
}

func (p *Subtraction) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return libs.Sandbox.Index.ConstIndex.New(p.operator)
					},
					Content: treasure,
					Types:   phrase_type.Operator,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewSubtraction(param *adaptor.SourceAdaptorParam) *Subtraction {
	instance := (&Subtraction{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	instance.operator = instance.Libs.GetLibraryPage("system", "operator").GetFunction(libs.Sandbox.Variable.String.New("SubtractionFunc"))
	return instance
}
