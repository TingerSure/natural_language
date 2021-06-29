package priority

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/structs"
	"github.com/TingerSure/natural_language/language/chinese/system/word/operator"
)

const (
	operatorLevelName string = "priority.operator_level"
)

type OperatorLevel struct {
	*adaptor.SourceAdaptor
}

func (o *OperatorLevel) getLevel(key string) int {
	switch key {
	case operator.AdditionName, operator.SubtractionName:
		return 1
	case operator.MultiplicationName, operator.DivisionName:
		return 2
	}
	return 0
}

func (o *OperatorLevel) GetPriorityRules() []*tree.PriorityRule {
	return []*tree.PriorityRule{
		tree.NewPriorityRule(&tree.PriorityRuleParam{
			Match: func(left tree.Phrase, right tree.Phrase) bool {
				return left.From() == structs.NumberFromNumberArithmeticNumberName &&
					right.From() == structs.NumberFromNumberArithmeticNumberName
			},
			Chooser: func(left tree.Phrase, right tree.Phrase) *tree.PriorityResult {

				levelLeft := o.getLevel(left.GetChild(1).From())
				levelRight := o.getLevel(right.GetChild(1).From())
				indexLeft := left.GetChild(0)
				indexRight := right.GetChild(0)

				if levelLeft > levelRight {
					if indexLeft.ContentSize() < indexRight.ContentSize() {
						return tree.NewPriorityResult(1).AddAbandon(2)
					}
					return tree.NewPriorityResult(1).AddAbandon(0)
				}
				if levelLeft < levelRight {
					if indexLeft.ContentSize() > indexRight.ContentSize() {
						return tree.NewPriorityResult(-1).AddAbandon(2)
					}
					return tree.NewPriorityResult(-1).AddAbandon(0)
				}

				if indexLeft.ContentSize() < indexRight.ContentSize() {
					return tree.NewPriorityResult(1).AddAbandon(2)
				}
				if indexLeft.ContentSize() > indexRight.ContentSize() {
					return tree.NewPriorityResult(-1).AddAbandon(2)
				}
				return 0, nil
			},
		}),
	}
}

func (p *OperatorLevel) GetName() string {
	return operatorLevelName
}

func NewOperatorLevel(param *adaptor.SourceAdaptorParam) *OperatorLevel {
	return (&OperatorLevel{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
