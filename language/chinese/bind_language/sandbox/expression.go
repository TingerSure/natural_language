package sandbox

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func ExpressionBindLanguage(libs *tree.LibraryManager, language string) {
	expression.CallLanguageSeeds[language] = func(language string, instance *expression.Call) string {

		defaultLanguage := func() string {
			return fmt.Sprintf("以%v来%v", instance.Param().ToLanguage(language), instance.Function().ToLanguage(language))
		}

		var funcs concept.Function = nil
		constIndexFuncs, yesIndexFuncs := index.IndexFamilyInstance.IsConstIndex(instance.Function())
		if yesIndexFuncs {
			funcsHome, yesFuncs := variable.VariableFamilyInstance.IsFunctionHome(constIndexFuncs.Value())
			if yesFuncs {
				funcs = funcsHome
			}
		}
		if !nl_interface.IsNil(funcs) {
			return defaultLanguage()
		}

		seed := funcs.GetLanguageOnCallSeed(language)
		if seed == nil {
			return defaultLanguage()
		}

		paramCanUse := false
		param := concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: variable.NewNull(),
		})

		constIndexParam, yesIndexParam := index.IndexFamilyInstance.IsConstIndex(instance.Param())
		if yesIndexParam {
			paramObject, yesParamObject := variable.VariableFamilyInstance.IsParam(constIndexParam.Value())
			if yesParamObject {
				paramCanUse = true
				paramObject.Iterate(func(key concept.String, value concept.Variable) bool {
					param.Set(key, value)
					return false
				})
			}
		}

		if !paramCanUse {
			newParamIndex, yesNewParamIndex := instance.Param().(*expression.NewParam)
			if yesNewParamIndex {
				paramCanUse = true
				newParamIndex.Iterate(func(key concept.String, value concept.Index) bool {
					param.Set(key, value)
					return false
				})
			}
		}
		if !paramCanUse {
			return defaultLanguage()
		}
		return seed(funcs, param)
	}

}
