package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/core/tree/phrase_types"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
)

const (
	AnyFromBracketAnyBracketName string = "structs.any.bracket_any_bracket"
)

var (
	AnyFromBracketAnyBracketList []string = []string{
		phrase_types.BracketsLeft,
		phrase_types.Any,
		phrase_types.BracketsRight,
	}
)

type AnyFromBracketAnyBracket struct {
	adaptor.SourceAdaptor
}

func (p *AnyFromBracketAnyBracket) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return phrase[1].Index()
					},
					Size: len(AnyFromBracketAnyBracketList),
					DynamicTypes: func(phrase []tree.Phrase) string {
						return phrase[1].Types()
					},
					From: p.GetName(),
				})
			},
			Types: AnyFromBracketAnyBracketList,
			From:  p.GetName(),
		}),
	}
}

func (p *AnyFromBracketAnyBracket) GetName() string {
	return AnyFromBracketAnyBracketName
}

func NewAnyFromBracketAnyBracket() *AnyFromBracketAnyBracket {
	return (&AnyFromBracketAnyBracket{})
}
