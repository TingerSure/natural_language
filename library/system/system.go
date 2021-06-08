package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/bind"
	"github.com/TingerSure/natural_language/library/system/operator"
	"github.com/TingerSure/natural_language/library/system/parser"
	"github.com/TingerSure/natural_language/library/system/runtime"
	"github.com/TingerSure/natural_language/library/system/std"
)

type SystemLibraryParam struct {
	StdParam     *std.StdParam
	RuntimeParam *runtime.RuntimeParam
}

func BindSystem(libs *tree.LibraryManager, param *SystemLibraryParam) {
	libs.AddPage("system/std", libs.Sandbox.Index.ConstIndex.New(std.NewStd(libs, param.StdParam)))

	libs.AddPage("system/operator", libs.Sandbox.Index.ConstIndex.New(operator.NewOperator(libs)))

	libs.AddPage("system/runtime", libs.Sandbox.Index.ConstIndex.New(runtime.NewRuntime(libs, param.RuntimeParam)))

	libs.AddPage("system/bind", libs.Sandbox.Index.ConstIndex.New(bind.NewBind(libs)))

	libs.AddPage("system/parser", libs.Sandbox.Index.ConstIndex.New(parser.NewParser(libs)))
}
