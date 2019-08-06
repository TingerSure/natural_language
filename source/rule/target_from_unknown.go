package rule

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	targetFromUnknownName string = "rule.target.unknown"
)

var (
	targetFromUnknownList []string = []string{
		phrase_types.Unknown,
	}
)

type TargetFromUnknown struct {
}

func (p *TargetFromUnknown) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(len(targetFromUnknownList), phrase_types.Target)
		}, targetFromUnknownList, p.GetName()),
	}
}

func (p *TargetFromUnknown) GetName() string {
	return targetFromUnknownName
}

func (p *TargetFromUnknown) GetWords(firstCharacter string) []*tree.Word {
	return nil
}

func (p *TargetFromUnknown) GetVocabularyRules() []*tree.VocabularyRule {
	return nil
}

func NewTargetFromUnknown() *TargetFromUnknown {
	return (&TargetFromUnknown{})
}
