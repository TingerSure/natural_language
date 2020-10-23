package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	QuestionFromNumberSetInterrogativeName string = "structs.question.number_set_interrogative"
)

var (
	QuestionFromNumberSetInterrogativeList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Number,
		phrase_type.Set,
		phrase_type.PronounInterrogative,
	}
)

type QuestionFromNumberSetInterrogative struct {
	*adaptor.SourceAdaptor
}

func (p *QuestionFromNumberSetInterrogative) GetStructRules() []*tree.StructRule {
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
					Size:  len(QuestionFromNumberSetInterrogativeList),
					Types: phrase_type.Question,
					From:  p.GetName(),
				})
			},
			Types: QuestionFromNumberSetInterrogativeList,
			From:  p.GetName(),
		}),
	}
}

func (p *QuestionFromNumberSetInterrogative) GetName() string {
	return QuestionFromNumberSetInterrogativeName
}

func NewQuestionFromNumberSetInterrogative(param *adaptor.SourceAdaptorParam) *QuestionFromNumberSetInterrogative {
	return (&QuestionFromNumberSetInterrogative{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
