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
	*adaptor.SourceAdaptor
}

func (o *Belong) GetPriorityRules() []*tree.PriorityRule {
	return []*tree.PriorityRule{
		tree.NewPriorityRule(&tree.PriorityRuleParam{
			Match: func(left tree.Phrase, right tree.Phrase) bool {
				return left.From() == structs.AnyFromAnyBelongAnyName &&
					right.From() == structs.AnyFromAnyBelongAnyName
			},
			Chooser: func(left tree.Phrase, right tree.Phrase) int {
				indexLeft := left.GetChild(0).ContentSize()
				indexRight := right.GetChild(0).ContentSize()
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

func (p *Belong) GetName() string {
	return belongName
}

func NewBelong(param *adaptor.SourceAdaptorParam) *Belong {
	return (&Belong{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
