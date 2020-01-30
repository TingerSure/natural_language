package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/core/tree/phrase_types"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
)

const (
	TargetFromTargetTargetName string = "structs.target.target_target"
)

var (
	targetFromTargetTargetList []string = []string{
		phrase_types.Target,
		phrase_types.Target,
	}
)

type TargetFromTargetTarget struct {
	adaptor.SourceAdaptor
}

func (p *TargetFromTargetTarget) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func([]tree.Phrase) concept.Index {
						return nil
						//TODO
					},
					Size:  len(targetFromTargetTargetList),
					Types: phrase_types.Target,
					From:  p.GetName(),
				})
			},
			Types: targetFromTargetTargetList,
			From:  p.GetName(),
		}),
	}
}

func (p *TargetFromTargetTarget) GetName() string {
	return TargetFromTargetTargetName
}

func NewTargetFromTargetTarget() *TargetFromTargetTarget {
	return (&TargetFromTargetTarget{})
}
