package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

const (
	Left   = "left"
	Right  = "right"
	Result = "result"
)

var (
	AdditionFunc       *variable.Function = nil
	DivisionFunc       *variable.Function = nil
	MultiplicationFunc *variable.Function = nil
	SubtractionFunc    *variable.Function = nil
)

func init() {
	AdditionFunc = variable.NewFunction(nil)
	AdditionFunc.AddParamName(Left)
	AdditionFunc.AddParamName(Right)
	AdditionFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewAddition(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

	DivisionFunc = variable.NewFunction(nil)
	DivisionFunc.AddParamName(Left)
	DivisionFunc.AddParamName(Right)
	DivisionFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewDivision(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

	MultiplicationFunc = variable.NewFunction(nil)
	MultiplicationFunc.AddParamName(Left)
	MultiplicationFunc.AddParamName(Right)
	MultiplicationFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewMultiplication(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

	SubtractionFunc = variable.NewFunction(nil)
	SubtractionFunc.AddParamName(Left)
	SubtractionFunc.AddParamName(Right)
	SubtractionFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewSubtraction(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

}
