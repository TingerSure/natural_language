package structs

import (
	"github.com/TingerSure/natural_language/library/std"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	AnyFromQuestionSetAnyName string = "structs.any.question_set_any"
)

var (
	anyFromQuestionSetAnyList []string = []string{
		phrase_types.Question,
		phrase_types.Set,
		phrase_types.Any,
	}
)

type AnyFromQuestionSetAny struct {
	adaptor.SourceAdaptor
}

func (p *AnyFromQuestionSetAny) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(func(phrase []tree.Phrase) concept.Index {
					return expression.NewParamGet(
						expression.NewCall(
							phrase[0].Index(),
							expression.NewNewParamWithInit(map[string]concept.Index{
								std.PrintfContent: phrase[2].Index(),
							}),
						),
						std.PrintfContent,
					)
				}, len(anyFromQuestionSetAnyList), phrase_types.Any, p.GetName())
			},
			Types: anyFromQuestionSetAnyList,
			From:  p.GetName(),
		}),
	}
}

func (p *AnyFromQuestionSetAny) GetName() string {
	return AnyFromQuestionSetAnyName
}

func NewAnyFromQuestionSetAny() *AnyFromQuestionSetAny {
	return (&AnyFromQuestionSetAny{})
}
