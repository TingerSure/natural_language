package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	DynamicFromBracketAnyBracketName string = "structs.dynamic.bracket_any_bracket"
)

var (
	DynamicFromBracketAnyBracketList []string = []string{
		phrase_type.BracketsLeftName,
		phrase_type.AnyName,
		phrase_type.BracketsRightName,
	}
)

type DynamicFromBracketAnyBracket struct {
	*adaptor.SourceAdaptor
}

func (p *DynamicFromBracketAnyBracket) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return phrase[1].Index()
					},
					Size: len(DynamicFromBracketAnyBracketList),
					DynamicTypes: func(phrase []tree.Phrase) string {
						return phrase[1].Types()
					},
					From: p.GetName(),
				})
			},
			Types: DynamicFromBracketAnyBracketList,
			From:  p.GetName(),
		}),
	}
}

func (p *DynamicFromBracketAnyBracket) GetName() string {
	return DynamicFromBracketAnyBracketName
}

func NewDynamicFromBracketAnyBracket(param *adaptor.SourceAdaptorParam) *DynamicFromBracketAnyBracket {
	return (&DynamicFromBracketAnyBracket{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
