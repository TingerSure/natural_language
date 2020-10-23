package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	QuestionFromInterrogativeSetNounName string = "structs.question.interrogative_set_noun"
)

const (
	QuestionParam  = "param"
	QuestionResult = "result"
)

var (
	QuestionFromInterrogativeSetNounList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.PronounInterrogative,
		phrase_type.Set,
		phrase_type.Noun,
	}
)

type QuestionFromInterrogativeSetNoun struct {
	*adaptor.SourceAdaptor
}

func (p *QuestionFromInterrogativeSetNoun) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
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
					Size:  len(QuestionFromInterrogativeSetNounList),
					Types: phrase_type.Question,
					From:  p.GetName(),
				})
			},
			Types: QuestionFromInterrogativeSetNounList,
			From:  p.GetName(),
		}),
	}
}

func (p *QuestionFromInterrogativeSetNoun) GetName() string {
	return QuestionFromInterrogativeSetNounName
}

func NewQuestionFromInterrogativeSetNoun(param *adaptor.SourceAdaptorParam) *QuestionFromInterrogativeSetNoun {
	return (&QuestionFromInterrogativeSetNoun{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
