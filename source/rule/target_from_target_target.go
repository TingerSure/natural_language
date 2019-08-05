package rule

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	targetFromTargetTargetName string = "rule.target.target_target"
	targetFromTargetTargetSize        = 2
)

type TargetFromTargetTarget struct {
}

func (p *TargetFromTargetTarget) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(targetFromTargetTargetSize, phrase_types.Target)
		}, []string{
			phrase_types.Target,
			phrase_types.Target,
		}, p.GetName()),
	}
}

func (p *TargetFromTargetTarget) GetName() string {
	return targetFromTargetTargetName
}

func (p *TargetFromTargetTarget) GetWords(firstCharacter string) []*tree.Word {
	return nil
}

func (p *TargetFromTargetTarget) GetVocabularyRules() []*tree.VocabularyRule {
	return nil
}

func NewTargetFromTargetTarget() *TargetFromTargetTarget {
	return (&TargetFromTargetTarget{})
}
