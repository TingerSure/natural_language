package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	NounName string = "types.noun"
)

type Noun struct {
	*adaptor.SourceAdaptor
}

func (p *Noun) GetName() string {
	return NounName
}

func (p *Noun) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: NounName,
			From: NounName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: EntityName,
				},
			},
		}),
	}
}

func NewNoun(param *adaptor.SourceAdaptorParam) *Noun {
	instance := (&Noun{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
