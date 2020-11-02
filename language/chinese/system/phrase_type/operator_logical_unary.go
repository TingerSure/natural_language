package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	OperatorLogicalUnaryName string = "types.operator.logical.unary"
)

type OperatorLogicalUnary struct {
	*adaptor.SourceAdaptor
}

func (p *OperatorLogicalUnary) GetName() string {
	return OperatorLogicalUnaryName
}

func (p *OperatorLogicalUnary) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    OperatorLogicalUnaryName,
			From:    OperatorLogicalUnaryName,
			Parents: nil,
		}),
	}
}

func NewOperatorLogicalUnary(param *adaptor.SourceAdaptorParam) *OperatorLogicalUnary {
	instance := (&OperatorLogicalUnary{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
