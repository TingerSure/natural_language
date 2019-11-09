package rule

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	eventFromTargetActionTargetName string = "rule.event.target_belong_target"
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

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(func([]tree.Phrase) concept.Index {
				return nil
				//TODO
			}, len(eventFromTargetActionTargetList), phrase_types.Event)
		}, eventFromTargetActionTargetList, p.GetName()),
	}
}

func (p *EventFromTargetActionTarget) GetName() string {
	return eventFromTargetActionTargetName
}

func NewEventFromTargetActionTarget() *EventFromTargetActionTarget {
	return (&EventFromTargetActionTarget{})
}
