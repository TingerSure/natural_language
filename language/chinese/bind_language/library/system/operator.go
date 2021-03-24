package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

func OperatorBindLanguage(libs *tree.LibraryManager, language string) {
	operator := libs.GetLibraryPage("system", "operator")

	AdditionFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("AdditionFunc"))
	AdditionLeft := operator.GetConst(libs.Sandbox.Variable.String.New("AdditionLeft"))
	AdditionRight := operator.GetConst(libs.Sandbox.Variable.String.New("AdditionRight"))
	AdditionResult := operator.GetConst(libs.Sandbox.Variable.String.New("AdditionResult"))

	AdditionLeft.SetLanguage(language, "被加数")
	AdditionRight.SetLanguage(language, "加数")
	AdditionResult.SetLanguage(language, "和")

	AdditionFunc.Name().SetLanguage(language, "相加")

	AdditionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(AdditionLeft).(concept.ToString)
		right := param.Get(AdditionRight).(concept.ToString)
		return fmt.Sprintf("%v加上%v", left.ToLanguage(language), right.ToLanguage(language))
	})

	SubtractionFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("SubtractionFunc"))
	SubtractionLeft := operator.GetConst(libs.Sandbox.Variable.String.New("SubtractionLeft"))
	SubtractionRight := operator.GetConst(libs.Sandbox.Variable.String.New("SubtractionRight"))
	SubtractionResult := operator.GetConst(libs.Sandbox.Variable.String.New("SubtractionResult"))

	SubtractionLeft.SetLanguage(language, "被减数")
	SubtractionRight.SetLanguage(language, "减数")
	SubtractionResult.SetLanguage(language, "差")

	SubtractionFunc.Name().SetLanguage(language, "相减")

	SubtractionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(SubtractionLeft).(concept.ToString)
		right := param.Get(SubtractionRight).(concept.ToString)
		return fmt.Sprintf("%v减去%v", left.ToLanguage(language), right.ToLanguage(language))
	})

	MultiplicationFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("MultiplicationFunc"))
	MultiplicationLeft := operator.GetConst(libs.Sandbox.Variable.String.New("MultiplicationLeft"))
	MultiplicationRight := operator.GetConst(libs.Sandbox.Variable.String.New("MultiplicationRight"))
	MultiplicationResult := operator.GetConst(libs.Sandbox.Variable.String.New("MultiplicationResult"))

	MultiplicationLeft.SetLanguage(language, "被乘数")
	MultiplicationRight.SetLanguage(language, "乘数")
	MultiplicationResult.SetLanguage(language, "积")

	MultiplicationFunc.Name().SetLanguage(language, "相乘")

	MultiplicationFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(MultiplicationLeft).(concept.ToString)
		right := param.Get(MultiplicationRight).(concept.ToString)
		return fmt.Sprintf("%v乘以%v", left.ToLanguage(language), right.ToLanguage(language))
	})

	DivisionFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("DivisionFunc"))
	DivisionLeft := operator.GetConst(libs.Sandbox.Variable.String.New("DivisionLeft"))
	DivisionRight := operator.GetConst(libs.Sandbox.Variable.String.New("DivisionRight"))
	DivisionResult := operator.GetConst(libs.Sandbox.Variable.String.New("DivisionResult"))

	DivisionLeft.SetLanguage(language, "被除数")
	DivisionRight.SetLanguage(language, "除数")
	DivisionResult.SetLanguage(language, "商")

	DivisionFunc.Name().SetLanguage(language, "相除")

	DivisionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(DivisionLeft).(concept.ToString)
		right := param.Get(DivisionRight).(concept.ToString)
		return fmt.Sprintf("%v除以%v", left.ToLanguage(language), right.ToLanguage(language))
	})

}
