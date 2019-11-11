package structs

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	TargetFromUnknownName string = "structs.target.unknown"
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
			}, len(targetFromUnknownList), phrase_types.Target, p.GetName())
		}, targetFromUnknownList, p.GetName()),
	}
}

func (p *TargetFromUnknown) GetName() string {
	return TargetFromUnknownName
}

func NewTargetFromUnknown() *TargetFromUnknown {
	return (&TargetFromUnknown{})
}
