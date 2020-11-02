package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	OperatorRelationalName string = "types.operator.relational"
)

type OperatorRelational struct {
	*adaptor.SourceAdaptor
}

func (p *OperatorRelational) GetName() string {
	return OperatorRelationalName
}

func (p *OperatorRelational) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    OperatorRelationalName,
			From:    OperatorRelationalName,
			Parents: nil,
		}),
	}
}

func NewOperatorRelational(param *adaptor.SourceAdaptorParam) *OperatorRelational {
	instance := (&OperatorRelational{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
