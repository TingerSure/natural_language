package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	AnyFromAnySetQuestionName string = "structs.any.any_set_question"
)

var (
	anyFromAnySetQuestionList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Any,
		phrase_type.Set,
		phrase_type.Question,
	}
)

type AnyFromAnySetQuestion struct {
	*adaptor.SourceAdaptor
}

func (p *AnyFromAnySetQuestion) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								phrase[2].Index(),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Index{
									p.Libs.Sandbox.Variable.String.New(QuestionParam): phrase[0].Index(),
								}),
							),
							p.Libs.Sandbox.Variable.String.New(QuestionResult),
						)
					},
					Size: len(anyFromAnySetQuestionList),
					DynamicTypes: func(phrase []tree.Phrase) *tree.PhraseType {
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

func NewAnyFromAnySetQuestion(param *adaptor.SourceAdaptorParam) *AnyFromAnySetQuestion {
	return (&AnyFromAnySetQuestion{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
