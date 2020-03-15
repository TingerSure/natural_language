package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/auto_number"
	"github.com/TingerSure/natural_language/library/system/object"
	"github.com/TingerSure/natural_language/library/system/operator"
	"github.com/TingerSure/natural_language/library/system/question"
	"github.com/TingerSure/natural_language/library/system/std"
)

type SystemLibrary struct {
	pages map[string]*tree.Page
}

func (s *SystemLibrary) GetPage(name string) *tree.Page {
	return s.pages[name]
}

func (s *SystemLibrary) SetPage(name string, value *tree.Page) tree.Library {
	s.pages[name] = value
	return s
}

type SystemLibraryParam struct {
	Std *std.StdParam
}

func NewSystemLibrary(param *SystemLibraryParam) *SystemLibrary {
	stdObject := std.NewStd(param.Std)
	system := &SystemLibrary{
		pages: map[string]*tree.Page{
			"std":         stdObject,
			"question":    question.NewQuestion(stdObject),
			"operator":    operator.NewOperator(),
			"object":      object.NewObject(),
			"auto_number": auto_number.NewAutoNumber(),
		},
	}

	return system
}
