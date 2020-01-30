package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/core/tree/phrase_types"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
)

const (
	EventFromTargetActionTargetName string = "structs.event.target_belong_target"
)

var (
	eventFromTargetActionTargetList []string = []string{
		phrase_types.Target,
		phrase_types.Action,
		phrase_types.Target,
	}
)

type EventFromTargetActionTarget struct {
	adaptor.SourceAdaptor
}

func (p *EventFromTargetActionTarget) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func([]tree.Phrase) concept.Index {
						return nil
						//TODO
					},
					Size:  len(eventFromTargetActionTargetList),
					Types: phrase_types.Event,
					From:  p.GetName(),
				})
			},
			Types: eventFromTargetActionTargetList,
			From:  p.GetName(),
		}),
	}
}

func (p *EventFromTargetActionTarget) GetName() string {
	return EventFromTargetActionTargetName
}

func NewEventFromTargetActionTarget() *EventFromTargetActionTarget {
	return (&EventFromTargetActionTarget{})
}
