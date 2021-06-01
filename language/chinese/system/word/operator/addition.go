package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	AdditionName string = "word.operator.addition"

	additionCharactor = "+"
)

type Addition struct {
	*adaptor.SourceAdaptor
	operator  concept.Function
	instances []*tree.Vocabulary
}

func (p *Addition) GetName() string {
	return AdditionName
}

func (p *Addition) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
}

func (p *Addition) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Pipe {
						return p.Libs.Sandbox.Index.ConstIndex.New(p.operator)
					},
					Content: treasure,
					Types:   phrase_type.OperatorArithmeticName,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewAddition(param *adaptor.SourceAdaptorParam) *Addition {
	instance := (&Addition{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	instance.operator = instance.Libs.GetLibraryPage("system", "operator").GetFunction(instance.Libs.Sandbox.Variable.String.New("AdditionFunc"))
	instance.instances = []*tree.Vocabulary{
		tree.NewVocabulary(additionCharactor, instance),
	}
	return instance
}
