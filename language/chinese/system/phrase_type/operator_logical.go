package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	OperatorLogicalName string = "types.operator.logical"
)

type OperatorLogical struct {
	*adaptor.SourceAdaptor
}

func (p *OperatorLogical) GetName() string {
	return OperatorLogicalName
}

func (p *OperatorLogical) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    OperatorLogicalName,
			From:    OperatorLogicalName,
			Parents: nil,
		}),
	}
}

func NewOperatorLogical(param *adaptor.SourceAdaptorParam) *OperatorLogical {
	instance := (&OperatorLogical{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
