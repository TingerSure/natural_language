package rule

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	eventFromTargetActionTargetName string = "rule.target.target_belong_target"
	eventFromTargetActionTargetSize        = 3
)

type EventFromTargetActionTarget struct {
}

func (p *EventFromTargetActionTarget) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(eventFromTargetActionTargetSize, phrase_types.Target)
		}, []string{
			phrase_types.Target,
			phrase_types.Action,
			phrase_types.Target,
		}, p.GetName()),
	}
}

func (p *EventFromTargetActionTarget) GetName() string {
	return eventFromTargetActionTargetName
}

func (p *EventFromTargetActionTarget) GetWords(firstCharacter string) []*tree.Word {
	return nil
}

func (p *EventFromTargetActionTarget) GetVocabularyRules() []*tree.VocabularyRule {
	return nil
}

func NewEventFromTargetActionTarget() *EventFromTargetActionTarget {
	return (&EventFromTargetActionTarget{})
}
