package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/source/word/question"
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
									question.QuestionParam: phrase[0].Index(),
								}),
							),
							question.QuestionResult,
						)
					},
					Size: len(anyFromAnySetQuestionList),
					DynamicTypes: func(phrase []tree.Phrase) string {
						return phrase[0].Types()
					},
					From: p.GetName(),
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
