package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	BoolName string = "types.bool"
)

type Bool struct {
	*adaptor.SourceAdaptor
}

func (p *Bool) GetName() string {
	return BoolName
}

func (p *Bool) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: BoolName,
			From: BoolName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: AnyName,
				},
			},
		}),
	}
}

func NewBool(param *adaptor.SourceAdaptorParam) *Bool {
	instance := (&Bool{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
