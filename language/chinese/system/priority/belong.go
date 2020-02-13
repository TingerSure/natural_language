package priority

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/structs"
)

const (
	belongName string = "priority.belong"
)

type Belong struct {
	adaptor.SourceAdaptor
}

func (o *Belong) GetPriorityRules() []*tree.PriorityRule {
	return []*tree.PriorityRule{
		tree.NewPriorityRule(&tree.PriorityRuleParam{
			Match: func(left tree.Phrase, right tree.Phrase) bool {
				return left.From() == structs.AnyFromAnyBelongAnyName ||
					right.From() == structs.AnyFromAnyBelongAnyName
				return false
			},
			Chooser: func(left tree.Phrase, right tree.Phrase) int {
				if left.From() == structs.AnyFromAnyBelongAnyName {
					return 1
				}

				if right.From() == structs.AnyFromAnyBelongAnyName {
					return -1
				}

				return 0
			},
		}),
	}
}

func (p *Belong) GetName() string {
	return belongName
}

func NewBelong() *Belong {
	return (&Belong{})
}
