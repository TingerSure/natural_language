package bind

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newIndexCreator(libs *tree.LibraryManager) concept.Object {
	indexes := libs.Sandbox.Variable.Object.New()
	indexConstBind(libs, indexes)
	indexBubbleBind(libs, indexes)
	return indexes
}

func indexConstBind(libs *tree.LibraryManager, indexes concept.Object) {
	seedParam := libs.Sandbox.Variable.String.New("seed")
	languageParam := libs.Sandbox.Variable.String.New("language")
	indexes.SetField(
		libs.Sandbox.Variable.String.New("constIndexBind"),
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
				libs.Sandbox.Index.ConstIndex.Seeds[language.Value()] = func(pool concept.Pool, instance *index.ConstIndex) (string, concept.Exception) {
					seedInput := libs.Sandbox.Variable.Param.New()
					inputValue, suspend := instance.Value().ToLanguage(language.Value(), pool)
					if !nl_interface.IsNil(suspend) {
						return "", suspend
					}
					seedInput.Set(libs.Sandbox.Variable.String.New("value"), libs.Sandbox.Variable.String.New(inputValue))
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

func indexBubbleBind(libs *tree.LibraryManager, indexes concept.Object) {
	seedParam := libs.Sandbox.Variable.String.New("seed")
	languageParam := libs.Sandbox.Variable.String.New("language")
	indexes.SetField(
		libs.Sandbox.Variable.String.New("bubbleIndexBind"),
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
				libs.Sandbox.Index.BubbleIndex.Seeds[language.Value()] = func(pool concept.Pool, instance *index.BubbleIndex) (string, concept.Exception) {
					seedInput := libs.Sandbox.Variable.Param.New()
					bubble, suspend := instance.Get(pool)
					if !nl_interface.IsNil(suspend) {
						return "", suspend.(concept.Exception)
					}
					bubbleValue, exception := bubble.ToLanguage(language.Value(), pool)
					if !nl_interface.IsNil(exception) {
						return "", exception
					}
					seedInput.Set(libs.Sandbox.Variable.String.New("value"), libs.Sandbox.Variable.String.New(bubbleValue))
					seedOutput, exception := seed.Exec(seedInput, nil)
					if !nl_interface.IsNil(exception) {
						return "", exception
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
