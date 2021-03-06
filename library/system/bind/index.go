package bind

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newConstIndexBind(libs *tree.LibraryManager) *variable.SystemFunction {
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
			libs.Sandbox.Index.ConstIndex.Seeds[language.Value()] = func(pool concept.Pool, instance *index.ConstIndex) (string, concept.Exception) {
				seedInput := libs.Sandbox.Variable.Param.New()
				inputValue, suspend := instance.Value().ToLanguage(language.Value(), pool)
				if !nl_interface.IsNil(suspend) {
					return "", suspend
				}
				seedInput.SetOriginal("value", libs.Sandbox.Variable.String.New(inputValue))
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
			}
			return libs.Sandbox.Variable.Param.New(), nil
		},
		[]concept.String{
			languageParam,
			seedParam,
		},
		[]concept.String{},
	)
}

func newBubbleIndexBind(libs *tree.LibraryManager) *variable.SystemFunction {
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
			libs.Sandbox.Index.BubbleIndex.Seeds[language.Value()] = func(pool concept.Pool, instance *index.BubbleIndex) (string, concept.Exception) {
				seedInput := libs.Sandbox.Variable.Param.New()
				bubbleValue, exception := pool.KeyBubble(instance.Key()).ToLanguage(language.Value(), pool)
				if !nl_interface.IsNil(exception) {
					return "", exception
				}
				seedInput.SetOriginal("value", libs.Sandbox.Variable.String.New(bubbleValue))
				seedOutput, exception := seed.Exec(seedInput, nil)
				if !nl_interface.IsNil(exception) {
					return "", exception
				}
				valuePre := seedOutput.GetOriginal("value")
				value, yes := variable.VariableFamilyInstance.IsStringHome(valuePre)
				if !yes {
					return "", libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param value is not a string: %v", valuePre.ToString("")))
				}
				return value.Value(), nil
			}
			return libs.Sandbox.Variable.Param.New(), nil
		},
		[]concept.String{
			languageParam,
			seedParam,
		},
		[]concept.String{},
	)
}
