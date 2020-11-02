package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	SetName string = "types.set"
)

type Set struct {
	*adaptor.SourceAdaptor
}

func (p *Set) GetName() string {
	return SetName
}

func (p *Set) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: SetName,
			From: SetName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: AnyName,
					Rule:  nil,
				},
			},
		}),
	}
}

func NewSet(param *adaptor.SourceAdaptorParam) *Set {
	instance := (&Set{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
