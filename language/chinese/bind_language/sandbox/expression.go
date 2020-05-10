package sandbox

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"strings"
)

func ExpressionBindLanguage(libs *tree.LibraryManager, language string) {

	expression.NewParamLanguageSeeds[language] = func(language string, instance *expression.NewParam) string {
		items := []string{}

		instance.Iterate(func(key concept.String, value concept.Index) bool {
			items = append(items, fmt.Sprintf("%v作为%v", value.ToLanguage(language), key.ToLanguage(language)))
			return false
		})

		return strings.Join(items, "")

	}

	expression.ParamGetLanguageSeeds[language] = func(language string, instance *expression.ParamGet) string {
		callIndex, yesCallIndex := instance.Param().(*expression.Call)
		if yesCallIndex {
			constIndexFuncs, yesIndexFuncs := index.IndexFamilyInstance.IsConstIndex(callIndex.Function())
			if yesIndexFuncs {
				funcsHome, yesFuncs := variable.VariableFamilyInstance.IsFunctionHome(constIndexFuncs.Value())
				if yesFuncs {
					if len(funcsHome.ReturnNames()) == 1 {
						return instance.Param().ToLanguage(language)
					}
				}
			}
		}
		return fmt.Sprintf("%v的%v", instance.Param().ToLanguage(language), instance.Key().ToLanguage(language))

	}

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
