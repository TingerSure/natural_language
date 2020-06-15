package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

func OperatorBindLanguage(libs *runtime.LibraryManager, language string) {
	operator := libs.GetLibraryPage("system", "operator")

	AdditionFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("AdditionFunc"))
	AdditionLeft := operator.GetConst(libs.Sandbox.Variable.String.New("AdditionLeft"))
	AdditionRight := operator.GetConst(libs.Sandbox.Variable.String.New("AdditionRight"))
	AdditionResult := operator.GetConst(libs.Sandbox.Variable.String.New("AdditionResult"))

	AdditionLeft.SetLanguage(language, "augend")
	AdditionRight.SetLanguage(language, "addend")
	AdditionResult.SetLanguage(language, "sum")

	AdditionFunc.Name().SetLanguage(language, "add")

	AdditionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(AdditionLeft).(concept.ToString)
		right := param.Get(AdditionRight).(concept.ToString)
		return fmt.Sprintf("%v plus %v", left.ToLanguage(language), right.ToLanguage(language))
	})

	SubtractionFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("SubtractionFunc"))
	SubtractionLeft := operator.GetConst(libs.Sandbox.Variable.String.New("SubtractionLeft"))
	SubtractionRight := operator.GetConst(libs.Sandbox.Variable.String.New("SubtractionRight"))
	SubtractionResult := operator.GetConst(libs.Sandbox.Variable.String.New("SubtractionResult"))

	SubtractionLeft.SetLanguage(language, "minuend")
	SubtractionRight.SetLanguage(language, "subtrahend")
	SubtractionResult.SetLanguage(language, "difference")

	SubtractionFunc.Name().SetLanguage(language, "subtract")

	SubtractionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(SubtractionLeft).(concept.ToString)
		right := param.Get(SubtractionRight).(concept.ToString)
		return fmt.Sprintf("%v minus %v", left.ToLanguage(language), right.ToLanguage(language))
	})

	MultiplicationFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("MultiplicationFunc"))
	MultiplicationLeft := operator.GetConst(libs.Sandbox.Variable.String.New("MultiplicationLeft"))
	MultiplicationRight := operator.GetConst(libs.Sandbox.Variable.String.New("MultiplicationRight"))
	MultiplicationResult := operator.GetConst(libs.Sandbox.Variable.String.New("MultiplicationResult"))

	MultiplicationLeft.SetLanguage(language, "multiplicand")
	MultiplicationRight.SetLanguage(language, "multiplier")
	MultiplicationResult.SetLanguage(language, "product")

	MultiplicationFunc.Name().SetLanguage(language, "multiply")

	MultiplicationFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(MultiplicationLeft).(concept.ToString)
		right := param.Get(MultiplicationRight).(concept.ToString)
		return fmt.Sprintf("%v times %v", left.ToLanguage(language), right.ToLanguage(language))
	})

	DivisionFunc := operator.GetFunction(libs.Sandbox.Variable.String.New("DivisionFunc"))
	DivisionLeft := operator.GetConst(libs.Sandbox.Variable.String.New("DivisionLeft"))
	DivisionRight := operator.GetConst(libs.Sandbox.Variable.String.New("DivisionRight"))
	DivisionResult := operator.GetConst(libs.Sandbox.Variable.String.New("DivisionResult"))

	DivisionLeft.SetLanguage(language, "dividend")
	DivisionRight.SetLanguage(language, "divisor")
	DivisionResult.SetLanguage(language, "quotient")

	DivisionFunc.Name().SetLanguage(language, "divide")

	DivisionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(DivisionLeft).(concept.ToString)
		right := param.Get(DivisionRight).(concept.ToString)
		return fmt.Sprintf("%v divided by %v", left.ToLanguage(language), right.ToLanguage(language))
	})

}
