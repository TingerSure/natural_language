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
	libs.AddPage("system/std", libs.Sandbox.Index.ConstIndex.New(stdInstance))

	libs.AddPage("system/question", libs.Sandbox.Index.ConstIndex.New(question.NewQuestion(libs, stdInstance)))

	libs.AddPage("system/set", libs.Sandbox.Index.ConstIndex.New(set.NewSet(libs)))

	libs.AddPage("system/operator", libs.Sandbox.Index.ConstIndex.New(operator.NewOperator(libs)))

}
