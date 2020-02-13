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
	adaptor.SourceAdaptor
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

func (o *OperatorLevel) getIndex(phrase tree.Phrase) int {
	if phrase.Size() == 0 {
		return 1
	}
	count := 0
	for index := 0; index < phrase.Size(); index++ {
		count += o.getIndex(phrase.GetChild(index))
	}
	return count
}

func (o *OperatorLevel) GetPriorityRules() []*tree.PriorityRule {
	return []*tree.PriorityRule{
		tree.NewPriorityRule(&tree.PriorityRuleParam{
			Match: func(left tree.Phrase, right tree.Phrase) bool {
				return left.From() == structs.NumberFromNumberOperatorNumberName &&
					right.From() == structs.NumberFromNumberOperatorNumberName
			},
			Chooser: func(left tree.Phrase, right tree.Phrase) int {

				operatorLeft := o.getLevel(left.GetChild(1).From())
				operatorRight := o.getLevel(right.GetChild(1).From())
				if operatorLeft > operatorRight {
					return 1
				}
				if operatorLeft < operatorRight {
					return -1
				}

				indexLeft := o.getIndex(left.GetChild(0))
				indexRight := o.getIndex(right.GetChild(0))
				if indexLeft < indexRight {
					return 1
				}
				if indexLeft > indexRight {
					return -1
				}

				return 0
			},
		}),
	}
}

func (p *OperatorLevel) GetName() string {
	return operatorLevelName
}

func NewOperatorLevel() *OperatorLevel {
	return (&OperatorLevel{})
}
