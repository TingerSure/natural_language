package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	AnyFromQuestionSetAnyName string = "structs.any.question_set_any"
)

const (
	QuestionParam  = "param"
	QuestionResult = "result"
)

var (
	anyFromQuestionSetAnyList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Question,
		phrase_type.Set,
		phrase_type.Any,
	}
)

type AnyFromQuestionSetAny struct {
	*adaptor.SourceAdaptor
}

func (p *AnyFromQuestionSetAny) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return expression.NewParamGet(
							expression.NewCall(
								phrase[0].Index(),
								expression.NewNewParamWithInit(map[concept.String]concept.Index{
									variable.NewString(QuestionParam): phrase[2].Index(),
								}),
							),
							variable.NewString(QuestionResult),
						)
					},
					Size: len(anyFromQuestionSetAnyList),
					DynamicTypes: func(phrase []tree.Phrase) *tree.PhraseType {
						return phrase[2].Types()
					},
					From: p.GetName(),
				})
			},
			Types: anyFromQuestionSetAnyList,
			From:  p.GetName(),
		}),
	}
}

func (p *AnyFromQuestionSetAny) GetName() string {
	return AnyFromQuestionSetAnyName
}

func NewAnyFromQuestionSetAny(param *adaptor.SourceAdaptorParam) *AnyFromQuestionSetAny {
	return (&AnyFromQuestionSetAny{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
