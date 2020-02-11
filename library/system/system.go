package library

import (
	"github.com/TingerSure/natural_language/core/tree/page"
	// "github.com/TingerSure/natural_language/library/system/system/object"
	// "github.com/TingerSure/natural_language/library/system/system/operator"
	"github.com/TingerSure/natural_language/library/system/system/question"
	"github.com/TingerSure/natural_language/library/system/system/std"
)

type SystemLibrary struct {
	functions map[string]tree.Page
}

func (s *SystemLibrary) GetPage(name string) tree.Page {
	return s.functions[name]
}

func NewSystemLibrary() SystemLibrary {
	system := &SystemLibrary{
		functions: map[string]tree.Page{
			"std":      std.NewStd(),
			"question": question.NewQuestion(),

			// "question.HowMany": question.HowMany,
			// "question.What":    question.What,
			//
			// "operator.AdditionFunc":       operator.AdditionFunc,
			// "operator.DivisionFunc":       operator.DivisionFunc,
			// "operator.MultiplicationFunc": operator.MultiplicationFunc,
			// "operator.SubtractionFunc":    operator.SubtractionFunc,
			//
			// "object.GetField": object.GetField,
		},
	}

	return system
}
