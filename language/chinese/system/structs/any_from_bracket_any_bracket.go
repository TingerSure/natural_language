package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	AnyFromBracketAnyBracketName string = "structs.any.bracket_any_bracket"
)

var (
	AnyFromBracketAnyBracketList []string = []string{
		phrase_type.BracketsLeftName,
		phrase_type.AnyName,
		phrase_type.BracketsRightName,
	}
)

type AnyFromBracketAnyBracket struct {
	*adaptor.SourceAdaptor
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

func NewAnyFromBracketAnyBracket(param *adaptor.SourceAdaptorParam) *AnyFromBracketAnyBracket {
	return (&AnyFromBracketAnyBracket{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
