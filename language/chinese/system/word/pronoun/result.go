package pronoun

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	ResultName string = "word.pronoun.result"

	ResultCharactor string = "结果"
)

var (
	resultPronounWords []*tree.Word = []*tree.Word{
		tree.NewWord(ResultCharactor),
	}
)

type Result struct {
	*adaptor.SourceAdaptor
	ResultIndex concept.Index
}

func (p *Result) GetName() string {
	return ResultName
}

func (p *Result) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(resultPronounWords, sentence)
}

func (p *Result) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return p.ResultIndex
					},
					Content: treasure,
					Types:   phrase_type.Any,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewResult(param *adaptor.SourceAdaptorParam) *Result {
	instance := (&Result{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	page := instance.Libs.GetLibraryPage("system", "pronoun")
	instance.ResultIndex = page.GetIndex(instance.Libs.Sandbox.Variable.String.New("Result"))
	return instance
}
