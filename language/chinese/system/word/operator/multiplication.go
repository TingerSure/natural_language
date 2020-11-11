package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	MultiplicationName string = "word.operator.multiplication"

	multiplicationCharactor = "*"
)

type Multiplication struct {
	*adaptor.SourceAdaptor
	operator  concept.Function
	instances []*tree.Vocabulary
}

func (p *Multiplication) GetName() string {
	return MultiplicationName
}

func (p *Multiplication) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
}

func (p *Multiplication) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Index {
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

func NewMultiplication(param *adaptor.SourceAdaptorParam) *Multiplication {
	instance := (&Multiplication{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	instance.operator = instance.Libs.GetLibraryPage("system", "operator").GetFunction(instance.Libs.Sandbox.Variable.String.New("MultiplicationFunc"))
	instance.instances = []*tree.Vocabulary{
		tree.NewVocabulary(multiplicationCharactor, instance),
	}
	return instance
}
