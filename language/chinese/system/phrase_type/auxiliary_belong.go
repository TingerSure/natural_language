package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	AuxiliaryBelongName string = "types.auxiliary.belong"
)

type AuxiliaryBelong struct {
	*adaptor.SourceAdaptor
}

func (p *AuxiliaryBelong) GetName() string {
	return AuxiliaryBelongName
}

func (p *AuxiliaryBelong) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    AuxiliaryBelongName,
			From:    AuxiliaryBelongName,
			Parents: nil,
		}),
	}
}

func NewAuxiliaryBelong(param *adaptor.SourceAdaptorParam) *AuxiliaryBelong {
	instance := (&AuxiliaryBelong{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
