package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	QuestionFromEntitySetInterrogativeName string = "structs.question.entity_set_interrogative"
)

var (
	QuestionFromEntitySetInterrogativeList []string = []string{
		phrase_type.EntityName,
		phrase_type.SetName,
		phrase_type.PronounInterrogativeName,
	}
)

type QuestionFromEntitySetInterrogative struct {
	*adaptor.SourceAdaptor
}

func (p *QuestionFromEntitySetInterrogative) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStruct(&tree.PhraseStructParam{
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
					Size:  len(QuestionFromEntitySetInterrogativeList),
					Types: phrase_type.QuestionName,
					From:  p.GetName(),
				})
			},
			Types: QuestionFromEntitySetInterrogativeList,
			From:  p.GetName(),
		}),
	}
}

func (p *QuestionFromEntitySetInterrogative) GetName() string {
	return QuestionFromEntitySetInterrogativeName
}

func NewQuestionFromEntitySetInterrogative(param *adaptor.SourceAdaptorParam) *QuestionFromEntitySetInterrogative {
	return (&QuestionFromEntitySetInterrogative{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
