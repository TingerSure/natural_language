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

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(func([]tree.Phrase) concept.Index {
				return nil
				//TODO
			}, len(targetFromTargetTargetList), phrase_types.Target, p.GetName())
		}, targetFromTargetTargetList, p.GetName()),
	}
}

func (p *TargetFromTargetTarget) GetName() string {
	return TargetFromTargetTargetName
}

func NewTargetFromTargetTarget() *TargetFromTargetTarget {
	return (&TargetFromTargetTarget{})
}