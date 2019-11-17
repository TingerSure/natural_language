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
	AnyFromAnySetQuestionName string = "structs.any.any_set_question"
)

var (
	anyFromAnySetQuestionList []string = []string{
		phrase_types.Any,
		phrase_types.Set,
		phrase_types.Question,
	}
)

type AnyFromAnySetQuestion struct {
	adaptor.SourceAdaptor
}

func (p *AnyFromAnySetQuestion) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return expression.NewParamGet(
							expression.NewCall(
								phrase[2].Index(),
								expression.NewNewParamWithInit(map[string]concept.Index{
									std.PrintfContent: phrase[0].Index(),
								}),
							),
							std.PrintfContent,
						)
					},
					Size:  len(anyFromAnySetQuestionList),
					Types: phrase_types.Any,
					From:  p.GetName(),
				})
			},
			Types: anyFromAnySetQuestionList,
			From:  p.GetName(),
		}),
	}
}

func (p *AnyFromAnySetQuestion) GetName() string {
	return AnyFromAnySetQuestionName
}

func NewAnyFromAnySetQuestion() *AnyFromAnySetQuestion {
	return (&AnyFromAnySetQuestion{})
}
