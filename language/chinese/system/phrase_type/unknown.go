package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	UnknownName string = "types.unknown"
)

type Unknown struct {
	*adaptor.SourceAdaptor
}

func (p *Unknown) GetName() string {
	return UnknownName
}

func (p *Unknown) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    UnknownName,
			From:    UnknownName,
			Parents: nil,
		}),
	}
}

func NewUnknown(param *adaptor.SourceAdaptorParam) *Unknown {
	instance := (&Unknown{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
