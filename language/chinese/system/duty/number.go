package duty

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	NumberName string = "duty.number"
)

type Number struct {
	*adaptor.SourceAdaptor
}

func NewNumber(param *adaptor.SourceAdaptorParam) *Number {
	return &Number{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	}
}

func (p *Number) GetName() string {
	return NumberName
}

func (n *Number) GetDutyRules() []*tree.DutyRule {
	return []*tree.DutyRule{
		tree.NewDutyRule(&tree.DutyRuleParam{
			From: n.GetName(),
			Match: func(value concept.Variable) bool {
				_, yes := variable.VariableFamilyInstance.IsNumber(value)
				return yes
			},
			Create: func(value concept.Variable) string {
				return phrase_type.NumberName
			},
		}),
	}
}
