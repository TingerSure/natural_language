package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

const (
	Left   = "left"
	Right  = "right"
	Result = "result"
)

type Operator struct {
	*tree.Page
	AdditionFunc       *variable.Function
	DivisionFunc       *variable.Function
	MultiplicationFunc *variable.Function
	SubtractionFunc    *variable.Function
}

func NewOperator() *Operator {
	instance := &Operator{
		Page:               tree.NewPage(),
		AdditionFunc:       additionFunc,
		DivisionFunc:       divisionFunc,
		MultiplicationFunc: multiplicationFunc,
		SubtractionFunc:    subtractionFunc,
	}
	instance.SetFunction(variable.NewString("AdditionFunc"), instance.AdditionFunc)
	instance.SetFunction(variable.NewString("DivisionFunc"), instance.DivisionFunc)
	instance.SetFunction(variable.NewString("MultiplicationFunc"), instance.MultiplicationFunc)
	instance.SetFunction(variable.NewString("SubtractionFunc"), instance.SubtractionFunc)
	instance.SetConst(variable.NewString("Left"), Left)
	instance.SetConst(variable.NewString("Right"), Right)
	instance.SetConst(variable.NewString("Result"), Result)
	return instance
}

var (
	additionFunc       *variable.Function = nil
	divisionFunc       *variable.Function = nil
	multiplicationFunc *variable.Function = nil
	subtractionFunc    *variable.Function = nil
)

func init() {
	additionFunc = variable.NewFunction(nil)
	additionFunc.AddParamName(Left)
	additionFunc.AddParamName(Right)
	additionFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewAddition(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

	divisionFunc = variable.NewFunction(nil)
	divisionFunc.AddParamName(Left)
	divisionFunc.AddParamName(Right)
	divisionFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewDivision(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

	multiplicationFunc = variable.NewFunction(nil)
	multiplicationFunc.AddParamName(Left)
	multiplicationFunc.AddParamName(Right)
	multiplicationFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewMultiplication(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

	subtractionFunc = variable.NewFunction(nil)
	subtractionFunc.AddParamName(Left)
	subtractionFunc.AddParamName(Right)
	subtractionFunc.Body().AddStep(
		expression.NewReturn(
			Result,
			expression.NewSubtraction(
				index.NewLocalIndex(Left),
				index.NewLocalIndex(Right),
			),
		),
	)

}
