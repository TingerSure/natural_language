package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/number"
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

func NewSystemLibrary(libs *tree.LibraryManager, param *SystemLibraryParam) tree.Library {
	system := tree.NewLibraryAdaptor()

	stdInstance := std.NewStd(libs, param.Std)
	system.SetPage("std", stdInstance)

	system.SetPage("question", question.NewQuestion(libs, stdInstance))

	system.SetPage("set", set.NewSet(libs))

	system.SetPage("number", number.NewNumber(libs))

	system.SetPage("operator", operator.NewOperator(libs))

	system.SetPage("object", object.NewObject(libs))

	system.SetPage("pronoun", pronoun.NewPronoun(libs))

	return system
}
