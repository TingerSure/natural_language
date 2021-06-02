package runtime

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

func FunctionHomeInit(libs *tree.LibraryManager, instance concept.Function) {
	instance.SetField(
		libs.Sandbox.Variable.DelayString.New("paramList"),
		libs.Sandbox.Variable.DelayFunction.New(FunctionHomeParamList(libs, instance)),
	)
	instance.SetField(
		libs.Sandbox.Variable.DelayString.New("returnList"),
		libs.Sandbox.Variable.DelayFunction.New(FunctionHomeReturnList(libs, instance)),
	)
}

func FunctionHomeParamList(libs *tree.LibraryManager, instance concept.Function) func() concept.Function {
	return func() concept.Function {
		backList := libs.Sandbox.Variable.String.New("list")
		return libs.Sandbox.Variable.SystemFunction.New(
			func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				paramNames := libs.Sandbox.Variable.Array.New()
				for _, paramName := range instance.ParamNames() {
					paramNames.Append(paramName)
				}
				back := libs.Sandbox.Variable.Param.New()
				back.Set(backList, paramNames)
				return back, nil
			},
			nil,
			[]concept.String{},
			[]concept.String{
				backList,
			},
		)
	}
}

func FunctionHomeReturnList(libs *tree.LibraryManager, instance concept.Function) func() concept.Function {
	return func() concept.Function {
		backList := libs.Sandbox.Variable.String.New("list")
		return libs.Sandbox.Variable.SystemFunction.New(
			func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				returnNames := libs.Sandbox.Variable.Array.New()
				for _, returnName := range instance.ReturnNames() {
					returnNames.Append(returnName)
				}
				back := libs.Sandbox.Variable.Param.New()
				back.Set(backList, returnNames)
				return back, nil
			},
			nil,
			[]concept.String{},
			[]concept.String{
				backList,
			},
		)
	}
}
