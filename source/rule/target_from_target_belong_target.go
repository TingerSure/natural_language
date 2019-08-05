package rule

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	targetFromTargetBelongTargetName string = "rule.target.target_belong_target"
	targetFromTargetBelongTargetSize        = 3
)

type TargetFromTargetBelongTarget struct {
}

func (p *TargetFromTargetBelongTarget) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(targetFromTargetBelongTargetSize, phrase_types.Target)
		}, []string{
			phrase_types.Target,
			phrase_types.AuxiliaryBelong,
			phrase_types.Target,
		}, p.GetName()),
	}
}

func (p *TargetFromTargetBelongTarget) GetName() string {
	return targetFromTargetBelongTargetName
}

func (p *TargetFromTargetBelongTarget) GetWords(firstCharacter string) []*tree.Word {
	return nil
}

func (p *TargetFromTargetBelongTarget) GetVocabularyRules() []*tree.VocabularyRule {
	return nil
}

func NewTargetFromTargetBelongTarget() *TargetFromTargetBelongTarget {
	return (&TargetFromTargetBelongTarget{})
}
