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

type Result struct {
	*adaptor.SourceAdaptor
	ResultIndex concept.Index
	instances   []*tree.Vocabulary
}

func (p *Result) GetName() string {
	return ResultName
}

func (p *Result) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
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
	instance.instances = []*tree.Vocabulary{
		tree.NewVocabulary(ResultCharactor, instance),
	}
	return instance
}
