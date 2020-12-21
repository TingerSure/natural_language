package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	PronounPersonalName string = "types.pronoun.personal"
)

type PronounPersonal struct {
	*adaptor.SourceAdaptor
}

func (p *PronounPersonal) GetName() string {
	return PronounPersonalName
}

func (p *PronounPersonal) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: PronounPersonalName,
			From: PronounPersonalName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: EntityName,
				},
			},
		}),
	}
}

func NewPronounPersonal(param *adaptor.SourceAdaptorParam) *PronounPersonal {
	instance := (&PronounPersonal{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
