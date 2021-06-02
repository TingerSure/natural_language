package runtime

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func StringInit(libs *tree.LibraryManager, instance *variable.String) {
	instance.SetField(
		libs.Sandbox.Variable.DelayString.New("setLanguage"),
		libs.Sandbox.Variable.DelayFunction.New(StringSetLanguage(libs, instance)),
	)
	instance.SetField(
		libs.Sandbox.Variable.DelayString.New("getLanguage"),
		libs.Sandbox.Variable.DelayFunction.New(StringGetLanguage(libs, instance)),
	)

}

func StringSetLanguage(libs *tree.LibraryManager, instance *variable.String) func() concept.Function {
	return func() concept.Function {
		languageParam := libs.Sandbox.Variable.String.New("language")
		valueParam := libs.Sandbox.Variable.String.New("value")
		return libs.Sandbox.Variable.SystemFunction.New(
			func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				languagePre := param.Get(languageParam)
				language, yes := variable.VariableFamilyInstance.IsStringHome(languagePre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param language is not a string: %v", languagePre))
				}
				valuePre := param.Get(valueParam)
				value, yes := variable.VariableFamilyInstance.IsStringHome(valuePre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param value is not a string: %v", languagePre))
				}
				instance.SetLanguage(language.Value(), value.Value())
				return libs.Sandbox.Variable.Param.New(), nil
			},
			func(param concept.Param, _ concept.Variable) concept.Param {
				return libs.Sandbox.Variable.Param.New()
			},
			[]concept.String{
				languageParam,
				valueParam,
			},
			[]concept.String{},
		)
	}
}

func StringGetLanguage(libs *tree.LibraryManager, instance *variable.String) func() concept.Function {
	return func() concept.Function {
		languageParam := libs.Sandbox.Variable.String.New("language")
		valueBack := libs.Sandbox.Variable.String.New("value")
		return libs.Sandbox.Variable.SystemFunction.New(
			func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				languagePre := input.Get(languageParam)
				language, yes := variable.VariableFamilyInstance.IsStringHome(languagePre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param language is not a string: %v", languagePre))
				}
				value := instance.GetLanguage(language.Value())
				output := libs.Sandbox.Variable.Param.New()
				output.Set(valueBack, libs.Sandbox.Variable.String.New(value))
				return output, nil
			},
			nil,
			[]concept.String{
				languageParam,
			},
			[]concept.String{
				valueBack,
			},
		)
	}
}
