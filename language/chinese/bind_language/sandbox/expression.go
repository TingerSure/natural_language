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
	libs.Sandbox.Expression.NewParam.Seeds[language] = func(language string, instance *expression.NewParam) string {
		items := []string{}

		instance.Iterate(func(key concept.String, value concept.Pipe) bool {
			items = append(items, fmt.Sprintf("%v作为%v", value.ToLanguage(language), key.ToLanguage(language)))
			return false
		})

		return strings.Join(items, "")

	}

	libs.Sandbox.Expression.ParamGet.Seeds[language] = func(language string, instance *expression.ParamGet) string {
		key := instance.Key()
		callIndex, yesCallIndex := instance.Param().(*expression.Call)
		if yesCallIndex {
			constIndexFuncs, yesIndexFuncs := index.IndexFamilyInstance.IsConstIndex(callIndex.Function())
			if yesIndexFuncs {
				funcsHome, yesFuncs := variable.VariableFamilyInstance.IsFunctionHome(constIndexFuncs.Value())
				if yesFuncs {
					if len(funcsHome.ReturnNames()) == 1 {
						return instance.Param().ToLanguage(language)
					}
					key = funcsHome.ReturnFormat(key)
				}
			}
		}
		return fmt.Sprintf("%v的%v", instance.Param().ToLanguage(language), key.ToLanguage(language))

	}

	libs.Sandbox.Expression.Call.Seeds[language] = func(language string, instance *expression.Call) string {

		var funcs concept.Function = nil
		param := concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: libs.Sandbox.Variable.Null.New(),
		})
		paramCanUse := false

		constIndexFuncs, yesIndexFuncs := index.IndexFamilyInstance.IsConstIndex(instance.Function())
		if yesIndexFuncs {
			funcs, _ = variable.VariableFamilyInstance.IsFunctionHome(constIndexFuncs.Value())
		}

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
				newParamIndex.Iterate(func(key concept.String, value concept.Pipe) bool {
					param.Set(key, value)
					return false
				})
			}
		}

		if !paramCanUse {
			return fmt.Sprintf("以%v来%v", instance.Param().ToLanguage(language), instance.Function().ToLanguage(language))
		}

		if !nl_interface.IsNil(funcs) {
			seed := funcs.GetLanguageOnCallSeed(language)
			param = funcs.ParamFormat(param)
			if seed != nil {
				return seed(funcs, param)
			}
		}

		items := []string{}
		param.Iterate(func(key concept.String, value interface{}) bool {
			items = append(items, fmt.Sprintf("%v作为%v", value.(concept.ToString).ToLanguage(language), key.ToLanguage(language)))
			return false
		})

		return fmt.Sprintf("以%v来%v", strings.Join(items, ""), instance.Function().ToLanguage(language))

	}

}
