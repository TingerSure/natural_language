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
				return left.From() == structs.EntityFromEntityBelongNounName &&
					right.From() == structs.EntityFromEntityBelongNounName
			},
			Chooser: func(left tree.Phrase, right tree.Phrase) (int, *tree.AbandonGroup) {
				indexLeft := left.GetChild(0).ContentSize()
				indexRight := right.GetChild(0).ContentSize()
				if indexLeft < indexRight {
					return 1, nil
				}
				if indexLeft > indexRight {
					return -1, nil
				}

				return 0, nil
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
