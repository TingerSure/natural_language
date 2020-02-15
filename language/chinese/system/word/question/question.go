package question

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	QuestionName   string = "word.question"
	QuestionParam  string = "param"
	QuestionResult string = "result"
)

type Question struct {
	adaptor.SourceAdaptor
	libs           *tree.LibraryManager
	QuestionParam  string
	QuestionResult string
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
