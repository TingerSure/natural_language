package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	QuestionFromNounSetInterrogativeName string = "structs.question.noun_set_interrogative"
)

var (
	QuestionFromNounSetInterrogativeList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Noun,
		phrase_type.Set,
		phrase_type.PronounInterrogative,
	}
)

type QuestionFromNounSetInterrogative struct {
	*adaptor.SourceAdaptor
}

func (p *QuestionFromNounSetInterrogative) GetStructRules() []*tree.StructRule {
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
					Size:  len(QuestionFromNounSetInterrogativeList),
					Types: phrase_type.Question,
					From:  p.GetName(),
				})
			},
			Types: QuestionFromNounSetInterrogativeList,
			From:  p.GetName(),
		}),
	}
}

func (p *QuestionFromNounSetInterrogative) GetName() string {
	return QuestionFromNounSetInterrogativeName
}

func NewQuestionFromNounSetInterrogative(param *adaptor.SourceAdaptorParam) *QuestionFromNounSetInterrogative {
	return (&QuestionFromNounSetInterrogative{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
