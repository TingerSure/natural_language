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

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func([]tree.Phrase) concept.Index {
						return nil
						//TODO
					},
					Size:  len(targetFromUnknownList),
					Types: phrase_types.Target,
					From:  p.GetName(),
				})
			},
			Types: targetFromUnknownList,
			From:  p.GetName(),
		}),
	}
}

func (p *TargetFromUnknown) GetName() string {
	return TargetFromUnknownName
}

func NewTargetFromUnknown() *TargetFromUnknown {
	return (&TargetFromUnknown{})
}
