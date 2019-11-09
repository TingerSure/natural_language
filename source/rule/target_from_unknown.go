package rule

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
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
	adaptor.SourceAdaptor
}

func (p *TargetFromUnknown) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(func([]tree.Phrase) concept.Index {
				return nil
				//TODO
			}, len(targetFromUnknownList), phrase_types.Target)
		}, targetFromUnknownList, p.GetName()),
	}
}

func (p *TargetFromUnknown) GetName() string {
	return targetFromUnknownName
}

func NewTargetFromUnknown() *TargetFromUnknown {
	return (&TargetFromUnknown{})
}
