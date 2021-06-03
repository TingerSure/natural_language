package bind

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newVariableCreator(libs *tree.LibraryManager) concept.Object {
	variables := libs.Sandbox.Variable.Object.New()
	variableStringBind(libs, variables)
	return variables
}

func variableStringBind(libs *tree.LibraryManager, variables concept.Object) {
	seedParam := libs.Sandbox.Variable.String.New("seed")
	languageParam := libs.Sandbox.Variable.String.New("language")
	variables.SetField(
		libs.Sandbox.Variable.String.New("stringBind"),
		libs.Sandbox.Variable.SystemFunction.New(
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
				libs.Sandbox.Variable.String.Seeds[language.Value()] = func(pool concept.Pool, instance *variable.String) (string, concept.Exception) {
					seedInput := libs.Sandbox.Variable.Param.New()
					seedInput.Set(libs.Sandbox.Variable.String.New("instance"), instance)
					seedOutput, suspend := seed.Exec(seedInput, nil)
					if !nl_interface.IsNil(suspend) {
						return "", suspend
					}
					valuePre := seedOutput.Get(libs.Sandbox.Variable.String.New("value"))
					value, yes := variable.VariableFamilyInstance.IsStringHome(valuePre)
					if !yes {
						return "", libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param value is not a string: %v", valuePre.ToString("")))
					}
					return value.Value(), nil
				}
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
		),
	)
}
