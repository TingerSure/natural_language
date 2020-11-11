package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	QuestionFromInterrogativeSetEntityName string = "structs.question.interrogative_set_entity"
)

const (
	QuestionParam  = "param"
	QuestionResult = "result"
)

var (
	QuestionFromInterrogativeSetEntityList []string = []string{
		phrase_type.PronounInterrogativeName,
		phrase_type.SetName,
		phrase_type.EntityName,
	}
)

type QuestionFromInterrogativeSetEntity struct {
	*adaptor.SourceAdaptor
}

func (p *QuestionFromInterrogativeSetEntity) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStruct(&tree.PhraseStructParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								phrase[0].Index(),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Index{
									p.Libs.Sandbox.Variable.String.New(QuestionParam): phrase[2].Index(),
								}),
							),
							p.Libs.Sandbox.Variable.String.New(QuestionResult),
						)
					},
					Size:  len(QuestionFromInterrogativeSetEntityList),
					Types: phrase_type.QuestionName,
					From:  p.GetName(),
				})
			},
			Types: QuestionFromInterrogativeSetEntityList,
			From:  p.GetName(),
		}),
	}
}

func (p *QuestionFromInterrogativeSetEntity) GetName() string {
	return QuestionFromInterrogativeSetEntityName
}

func NewQuestionFromInterrogativeSetEntity(param *adaptor.SourceAdaptorParam) *QuestionFromInterrogativeSetEntity {
	return (&QuestionFromInterrogativeSetEntity{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
