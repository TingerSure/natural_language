package system

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/auto_number"
	"github.com/TingerSure/natural_language/library/system/object"
	"github.com/TingerSure/natural_language/library/system/operator"
	"github.com/TingerSure/natural_language/library/system/pronoun"
	"github.com/TingerSure/natural_language/library/system/question"
	"github.com/TingerSure/natural_language/library/system/set"
	"github.com/TingerSure/natural_language/library/system/std"
)

type SystemLibraryParam struct {
	Std *std.StdParam
}

func NewSystemLibrary(libs *runtime.LibraryManager, param *SystemLibraryParam) tree.Library {
	system := tree.NewLibraryAdaptor()

	stdInstance := std.NewStd(libs, param.Std)
	system.SetPage("std", stdInstance)

	system.SetPage("question", question.NewQuestion(libs, stdInstance))

	system.SetPage("set", set.NewSet(libs))

	system.SetPage("operator", operator.NewOperator(libs))

	autoNumber := auto_number.NewAutoNumber(libs)
	system.SetPage("auto_number", autoNumber)

	system.SetPage("object", object.NewObject(libs, autoNumber))

	system.SetPage("pronoun", pronoun.NewPronoun(libs))

	return system
}
