package structs

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
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
