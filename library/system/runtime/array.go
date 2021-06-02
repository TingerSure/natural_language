package runtime

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func ArrayInit(libs *tree.LibraryManager, instance *variable.Array) {
	instance.SetField(
		libs.Sandbox.Variable.DelayString.New("size"),
		libs.Sandbox.Variable.DelayFunction.New(ArraySize(libs, instance)),
	)
}

func ArraySize(libs *tree.LibraryManager, instance *variable.Array) func() concept.Function {
	return func() concept.Function {
		backSize := libs.Sandbox.Variable.String.New("size")
		return libs.Sandbox.Variable.SystemFunction.New(
			func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				back := libs.Sandbox.Variable.Param.New()
				back.Set(backSize, libs.Sandbox.Variable.Number.New(float64(instance.Length())))
				return back, nil
			},
			nil,
			[]concept.String{},
			[]concept.String{
				backSize,
			},
		)
	}
}
