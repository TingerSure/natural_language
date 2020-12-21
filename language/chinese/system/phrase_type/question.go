package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	QuestionName string = "types.question"
)

type Question struct {
	*adaptor.SourceAdaptor
}

func (p *Question) GetName() string {
	return QuestionName
}

func (p *Question) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: QuestionName,
			From: QuestionName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: AnyName,
				},
			},
		}),
	}
}

func NewQuestion(param *adaptor.SourceAdaptorParam) *Question {
	instance := (&Question{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
