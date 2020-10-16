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
				return left.From() == structs.NumberFromNumberOperatorNumberName &&
					right.From() == structs.NumberFromNumberOperatorNumberName
			},
			Chooser: func(left tree.Phrase, right tree.Phrase) (int, *tree.AbandonGroup) {

				levelLeft := o.getLevel(left.GetChild(1).From())
				levelRight := o.getLevel(right.GetChild(1).From())
				indexLeft := left.GetChild(0)
				indexRight := right.GetChild(0)

				if levelLeft > levelRight {
					if indexLeft.ContentSize() < indexRight.ContentSize() {
						return 1, tree.NewAbandonGroup().Add(&tree.Abandon{
							Offset: 0,
							Value:  left.GetChild(2),
						})
					}
					return 1, tree.NewAbandonGroup().Add(&tree.Abandon{
						Offset: indexLeft.ContentSize() - left.ContentSize(),
						Value:  indexLeft,
					})
				}
				if levelLeft < levelRight {
					if indexLeft.ContentSize() > indexRight.ContentSize() {
						return -1, tree.NewAbandonGroup().Add(&tree.Abandon{
							Offset: 0,
							Value:  right.GetChild(2),
						})
					}
					return -1, tree.NewAbandonGroup().Add(&tree.Abandon{
						Offset: indexRight.ContentSize() - right.ContentSize(),
						Value:  indexRight,
					})
				}

				if indexLeft.ContentSize() < indexRight.ContentSize() {
					return 1, tree.NewAbandonGroup().Add(&tree.Abandon{
						Offset: 0,
						Value:  left.GetChild(2),
					})
				}
				if indexLeft.ContentSize() > indexRight.ContentSize() {
					return -1, tree.NewAbandonGroup().Add(&tree.Abandon{
						Offset: 0,
						Value:  right.GetChild(2),
					})
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
