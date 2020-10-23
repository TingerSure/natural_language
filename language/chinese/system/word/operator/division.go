package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	DivisionName string = "word.operator.division"

	divisionCharactor = "/"
)

type Division struct {
	*adaptor.SourceAdaptor
	operator  concept.Function
	instances []*tree.Vocabulary
}

func (p *Division) GetName() string {
	return DivisionName
}

func (p *Division) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
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
						return p.Libs.Sandbox.Index.ConstIndex.New(p.operator)
					},
					Content: treasure,
					Types:   phrase_type.OperatorArithmetic,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewDivision(param *adaptor.SourceAdaptorParam) *Division {
	instance := (&Division{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	instance.operator = instance.Libs.GetLibraryPage("system", "operator").GetFunction(instance.Libs.Sandbox.Variable.String.New("DivisionFunc"))
	instance.instances = []*tree.Vocabulary{
		tree.NewVocabulary(divisionCharactor, instance),
	}
	return instance
}
