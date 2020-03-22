package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

var (
	QuestionName   string         = "word.question"
	QuestionParam  concept.String = variable.NewString("param")
	QuestionResult concept.String = variable.NewString("result")
)

type Question struct {
	adaptor.SourceAdaptor
	libs           *tree.LibraryManager
	QuestionParam  concept.String
	QuestionResult concept.String
}

func (q *Question) GetName() string {
	return QuestionName
}

func NewQuestion(libs *tree.LibraryManager) *Question {
	return &Question{
		libs:           libs,
		QuestionParam:  QuestionParam,
		QuestionResult: QuestionResult,
	}
}
