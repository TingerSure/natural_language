package runtime

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
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
	instance.SetField(
		libs.Sandbox.Variable.DelayString.New("setLanguageOnCallSeed"),
		libs.Sandbox.Variable.DelayFunction.New(FunctionHomeSetLanguageOnCallSeed(libs, instance)),
	)
}

func FunctionHomeSetLanguageOnCallSeed(libs *tree.LibraryManager, instance concept.Function) func() concept.Function {
	return func() concept.Function {
		seedParam := libs.Sandbox.Variable.String.New("seed")
		languageParam := libs.Sandbox.Variable.String.New("language")
		return libs.Sandbox.Variable.SystemFunction.New(
			func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
				languagePre := input.Get(languageParam)
				language, yes := variable.VariableFamilyInstance.IsStringHome(languagePre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param language is not a string: %v", languagePre.ToString("")))
				}
				seedPre := input.Get(seedParam)
				seed, yes := variable.VariableFamilyInstance.IsFunctionHome(seedPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param seed is not a function: %v", seedPre.ToString("")))
				}
				instance.SetLanguageOnCallSeed(language.Value(), func(_ concept.Function, pool concept.Pool, name string, params concept.Param) (string, concept.Exception) {
					seedInput := libs.Sandbox.Variable.Param.New()
					seedInput.SetOriginal("instance", instance)
					seedInput.SetOriginal("name", libs.Sandbox.Variable.String.New(name))
					seedInput.SetOriginal("params", params)
					seedOutput, suspend := seed.Exec(seedInput, nil)
					if !nl_interface.IsNil(suspend) {
						return "", suspend
					}
					valuePre := seedOutput.GetOriginal("value")
					value, yes := variable.VariableFamilyInstance.IsStringHome(valuePre)
					if !yes {
						return "", libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param value is not a string: %v", valuePre.ToString("")))
					}
					return value.Value(), nil
				})
				return libs.Sandbox.Variable.Param.New(), nil
			},
			func(input concept.Param, object concept.Variable) concept.Param {
				return libs.Sandbox.Variable.Param.New()
			},
			[]concept.String{
				languageParam,
				seedParam,
			},
			[]concept.String{},
		)
	}
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
