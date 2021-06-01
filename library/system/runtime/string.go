package runtime

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newStringCreator(libs *tree.LibraryManager) concept.Object {
	creator := libs.Sandbox.Variable.Object.New()
	seedParam := libs.Sandbox.Variable.String.New("seed")
	languageParam := libs.Sandbox.Variable.String.New("language")
	creator.SetField(
		libs.Sandbox.Variable.String.New("BindSeed"),
		libs.Sandbox.Variable.SystemFunction.New(
			func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
				languagePre := input.Get(languageParam)
				language, yes := variable.VariableFamilyInstance.IsStringHome(languagePre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param language is not a string: %v", languagePre))
				}
				seedPre := input.Get(seedParam)
				seed, yes := variable.VariableFamilyInstance.IsFunctionHome(seedPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param seed is not a function: %v", seedPre))
				}
				libs.Sandbox.Variable.String.Seeds[language.Value()] = func(_ string, pool concept.Pool, instance *variable.String) string {
					seedInput := libs.Sandbox.Variable.Param.New()
					seedInput.Set(libs.Sandbox.Variable.String.New("pool"), newPoolObject(libs, pool))
					seedInput.Set(libs.Sandbox.Variable.String.New("instance"), instance)
					seedOutput, suspend := seed.Exec(seedInput, nil)
					if !nl_interface.IsNil(suspend) {
						return instance.ToString("")
					}
					valuePre := seedOutput.Get(libs.Sandbox.Variable.String.New("value"))
					value, yes := variable.VariableFamilyInstance.IsStringHome(valuePre)
					if !yes {
						return instance.ToString("")
					}
					return value.Value()
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
	return creator
}
