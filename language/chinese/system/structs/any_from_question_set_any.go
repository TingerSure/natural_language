package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
	"github.com/TingerSure/natural_language/language/chinese/system/word/question"
)

const (
	AnyFromQuestionSetAnyName string = "structs.any.question_set_any"
)

var (
	anyFromQuestionSetAnyList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Question,
		phrase_type.Set,
		phrase_type.Any,
	}
)

type AnyFromQuestionSetAny struct {
	adaptor.SourceAdaptor
	questionPackage *question.Question
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
									p.questionPackage.QuestionParam: phrase[2].Index(),
								}),
							),
							p.questionPackage.QuestionResult,
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

func NewAnyFromQuestionSetAny(libs *tree.LibraryManager, questionPackage *question.Question) *AnyFromQuestionSetAny {
	return (&AnyFromQuestionSetAny{
		questionPackage: questionPackage,
	})
}
