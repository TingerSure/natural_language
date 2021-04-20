package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/operator"
	"github.com/TingerSure/natural_language/library/system/question"
	"github.com/TingerSure/natural_language/library/system/set"
	"github.com/TingerSure/natural_language/library/system/std"
)

type SystemLibraryParam struct {
	Std *std.StdParam
}

func BindSystem(libs *tree.LibraryManager, param *SystemLibraryParam) {
	stdInstance := std.NewStd(libs, param.Std)
	libs.SetPage("system/std", stdInstance)

	libs.SetPage("system/question", question.NewQuestion(libs, stdInstance))

	libs.SetPage("system/set", set.NewSet(libs))

	libs.SetPage("system/operator", operator.NewOperator(libs))

}
