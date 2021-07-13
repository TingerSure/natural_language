package bind

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newCodeBlockBind(libs *tree.LibraryManager) *variable.SystemFunction {
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
			libs.Sandbox.Expression.CodeBlock.Seeds[language.Value()] = func(pool concept.Pool, instance *expression.CodeBlock) (string, concept.Exception) {
				seedInput := libs.Sandbox.Variable.Param.New()
				steps := libs.Sandbox.Variable.Array.New()
				for _, step := range instance.Steps() {
					stepValue, suspend := step.ToLanguage(language.Value(), pool)
					if !nl_interface.IsNil(suspend) {
						return "", suspend
					}
					steps.Append(libs.Sandbox.Variable.String.New(stepValue))
				}
				seedInput.SetOriginal("steps", steps)
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
