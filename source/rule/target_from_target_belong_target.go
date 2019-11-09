package rule

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	targetFromTargetBelongTargetName string = "rule.target.target_belong_target"
)

var (
	targetFromTargetBelongTargetList []string = []string{
		phrase_types.Target,
		phrase_types.AuxiliaryBelong,
		phrase_types.Target,
	}
)

type TargetFromTargetBelongTarget struct {
	adaptor.Adaptor
}

func (p *TargetFromTargetBelongTarget) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(func([]tree.Phrase) concept.Index {
				return nil
				//TODO
			}, len(targetFromTargetBelongTargetList), phrase_types.Target)
		}, targetFromTargetBelongTargetList, p.GetName()),
	}
}

func (p *TargetFromTargetBelongTarget) GetName() string {
	return targetFromTargetBelongTargetName
}

func NewTargetFromTargetBelongTarget() *TargetFromTargetBelongTarget {
	return (&TargetFromTargetBelongTarget{})
}
