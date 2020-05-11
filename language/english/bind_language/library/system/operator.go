package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func OperatorBindLanguage(libs *tree.LibraryManager, language string) {
	operator := libs.GetLibraryPage("system", "operator")

	AdditionFunc := operator.GetFunction(variable.NewString("AdditionFunc"))
	AdditionLeft := operator.GetConst(variable.NewString("AdditionLeft"))
	AdditionRight := operator.GetConst(variable.NewString("AdditionRight"))
	AdditionResult := operator.GetConst(variable.NewString("AdditionResult"))

	AdditionLeft.SetLanguage(language, "augend")
	AdditionRight.SetLanguage(language, "addend")
	AdditionResult.SetLanguage(language, "sum")

	AdditionFunc.Name().SetLanguage(language, "add")

	AdditionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(AdditionLeft).(concept.ToString)
		right := param.Get(AdditionRight).(concept.ToString)
		return fmt.Sprintf("%v plus %v", left.ToLanguage(language), right.ToLanguage(language))
	})

	SubtractionFunc := operator.GetFunction(variable.NewString("SubtractionFunc"))
	SubtractionLeft := operator.GetConst(variable.NewString("SubtractionLeft"))
	SubtractionRight := operator.GetConst(variable.NewString("SubtractionRight"))
	SubtractionResult := operator.GetConst(variable.NewString("SubtractionResult"))

	SubtractionLeft.SetLanguage(language, "minuend")
	SubtractionRight.SetLanguage(language, "subtrahend")
	SubtractionResult.SetLanguage(language, "difference")

	SubtractionFunc.Name().SetLanguage(language, "subtract")

	SubtractionFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(SubtractionLeft).(concept.ToString)
		right := param.Get(SubtractionRight).(concept.ToString)
		return fmt.Sprintf("%v minus %v", left.ToLanguage(language), right.ToLanguage(language))
	})

	MultiplicationFunc := operator.GetFunction(variable.NewString("MultiplicationFunc"))
	MultiplicationLeft := operator.GetConst(variable.NewString("MultiplicationLeft"))
	MultiplicationRight := operator.GetConst(variable.NewString("MultiplicationRight"))
	MultiplicationResult := operator.GetConst(variable.NewString("MultiplicationResult"))

	MultiplicationLeft.SetLanguage(language, "multiplicand")
	MultiplicationRight.SetLanguage(language, "multiplier")
	MultiplicationResult.SetLanguage(language, "product")

	MultiplicationFunc.Name().SetLanguage(language, "multiply")

	MultiplicationFunc.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		left := param.Get(MultiplicationLeft).(concept.ToString)
		right := param.Get(MultiplicationRight).(concept.ToString)
		return fmt.Sprintf("%v times %v", left.ToLanguage(language), right.ToLanguage(language))
	})

	DivisionFunc := operator.GetFunction(variable.NewString("DivisionFunc"))
	DivisionLeft := operator.GetConst(variable.NewString("DivisionLeft"))
	DivisionRight := operator.GetConst(variable.NewString("DivisionRight"))
	DivisionResult := operator.GetConst(variable.NewString("DivisionResult"))

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
