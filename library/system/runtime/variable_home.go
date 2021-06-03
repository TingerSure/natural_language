package runtime

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func VariableHomeInit(libs *tree.LibraryManager, instance concept.Variable) {
	instance.SetField(
		libs.Sandbox.Variable.DelayString.New("toString"),
		libs.Sandbox.Variable.DelayFunction.New(VariableHomeToString(libs, instance)),
	)
}

func VariableHomeToString(libs *tree.LibraryManager, instance concept.Variable) func() concept.Function {
	return func() concept.Function {
		valueParam := libs.Sandbox.Variable.String.New("value")
		prefixParam := libs.Sandbox.Variable.String.New("prefix")
		return libs.Sandbox.Variable.SystemFunction.New(
			func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				prefixPre := input.Get(prefixParam)
				prefix, yes := variable.VariableFamilyInstance.IsStringHome(prefixPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param prefix is not a string: %v", prefixPre.ToString("")))
				}
				output := libs.Sandbox.Variable.Param.New()
				output.Set(valueParam, libs.Sandbox.Variable.String.New(instance.ToString(prefix.Value())))
				return output, nil
			},
			nil,
			[]concept.String{
				prefixParam,
			},
			[]concept.String{
				valueParam,
			},
		)
	}
}
