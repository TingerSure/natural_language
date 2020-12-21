package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	PronounInterrogativeName string = "types.pronoun.interrogative"
)

type PronounInterrogative struct {
	*adaptor.SourceAdaptor
}

func (p *PronounInterrogative) GetName() string {
	return PronounInterrogativeName
}

func (p *PronounInterrogative) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: PronounInterrogativeName,
			From: PronounInterrogativeName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: AnyName,
				},
			},
		}),
	}
}

func NewPronounInterrogative(param *adaptor.SourceAdaptorParam) *PronounInterrogative {
	instance := (&PronounInterrogative{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
