package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	AnyName string = "types.any"
)

type Any struct {
	*adaptor.SourceAdaptor
}

func (p *Any) GetName() string {
	return AnyName
}

func (p *Any) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    AnyName,
			From:    AnyName,
			Parents: nil,
		}),
	}
}

func NewAny(param *adaptor.SourceAdaptorParam) *Any {
	instance := (&Any{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
