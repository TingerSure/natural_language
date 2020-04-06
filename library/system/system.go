package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/auto_number"
	"github.com/TingerSure/natural_language/library/system/object"
	"github.com/TingerSure/natural_language/library/system/operator"
	"github.com/TingerSure/natural_language/library/system/question"
	"github.com/TingerSure/natural_language/library/system/std"
)

type SystemLibraryParam struct {
	Std *std.StdParam
}

func NewSystemLibrary(param *SystemLibraryParam) tree.Library {
	stdObject := std.NewStd(param.Std)
	system := tree.NewLibraryAdaptor()
	system.SetPage("std", stdObject)
	system.SetPage("question", question.NewQuestion(stdObject))
	system.SetPage("operator", operator.NewOperator())
	system.SetPage("object", object.NewObject())
	system.SetPage("auto_number", auto_number.NewAutoNumber())

	return system
}
