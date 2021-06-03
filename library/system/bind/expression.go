package bind

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newExpressionCreator(libs *tree.LibraryManager) concept.Object {
	expressions := libs.Sandbox.Variable.Object.New()
	expressionCodeBlockBind(libs, expressions)
	return expressions
}

func expressionCodeBlockBind(libs *tree.LibraryManager, expressions concept.Object) {
	seedParam := libs.Sandbox.Variable.String.New("seed")
	languageParam := libs.Sandbox.Variable.String.New("language")
	expressions.SetField(
		libs.Sandbox.Variable.String.New("codeBlockBind"),
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
				libs.Sandbox.Expression.CodeBlock.Seeds[language.Value()] = func(_ string, pool concept.Pool, instance *expression.CodeBlock) string {
					seedInput := libs.Sandbox.Variable.Param.New()
					seedInput.Set(libs.Sandbox.Variable.String.New("pool"), pool)

					steps := libs.Sandbox.Variable.Array.New()
					for _, step := range instance.Steps() {
						steps.Append(libs.Sandbox.Variable.String.New(step.ToLanguage(language.Value(), pool)))
					}
					seedInput.Set(libs.Sandbox.Variable.String.New("steps"), steps)
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
}
