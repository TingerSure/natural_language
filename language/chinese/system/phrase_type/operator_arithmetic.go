package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	OperatorArithmeticName string = "types.operator.arithmetic"
)

type OperatorArithmetic struct {
	*adaptor.SourceAdaptor
}

func (p *OperatorArithmetic) GetName() string {
	return OperatorArithmeticName
}

func (p *OperatorArithmetic) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    OperatorArithmeticName,
			From:    OperatorArithmeticName,
			Parents: nil,
		}),
	}
}

func NewOperatorArithmetic(param *adaptor.SourceAdaptorParam) *OperatorArithmetic {
	instance := (&OperatorArithmetic{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
