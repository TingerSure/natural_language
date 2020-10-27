package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	EntityFromNounName string = "structs.entity.noun"
)

var (
	EntityFromNounList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Noun,
	}
)

type EntityFromNoun struct {
	*adaptor.SourceAdaptor
}

func (p *EntityFromNoun) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return phrase[0].Index()
					},
					Size:  len(EntityFromNounList),
					Types: phrase_type.Entity,
					From:  p.GetName(),
				})
			},
			Types: EntityFromNounList,
			From:  p.GetName(),
		}),
	}
}

func (p *EntityFromNoun) GetName() string {
	return EntityFromNounName
}

func NewEntityFromNoun(param *adaptor.SourceAdaptorParam) *EntityFromNoun {
	instance := (&EntityFromNoun{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
