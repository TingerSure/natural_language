package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

const (
	ItemLeft   = "left"
	ItemRight  = "right"
	ItemResult = "result"
)

type OperatorItem struct {
	Name   string
	Func   *variable.Function
	Left   concept.String
	Right  concept.String
	Result concept.String
}

func NewOperatorItem(create func(concept.Index, concept.Index) concept.Expression) *OperatorItem {
	instance := &OperatorItem{
		Left:   variable.NewString(ItemLeft),
		Right:  variable.NewString(ItemRight),
		Result: variable.NewString(ItemResult),
	}
	instance.Func = variable.NewFunction(nil)
	instance.Func.AddParamName(instance.Left)
	instance.Func.AddParamName(instance.Right)
	instance.Func.Body().AddStep(
		expression.NewReturn(
			instance.Result,
			create(
				index.NewLocalIndex(instance.Left),
				index.NewLocalIndex(instance.Right),
			),
		),
	)
	return instance
}

const (
	AdditionName       = "Addition"
	DivisionName       = "Division"
	MultiplicationName = "Multiplication"
	SubtractionName    = "Subtraction"

	FuncName   = "Func"
	LeftName   = "Left"
	RightName  = "Right"
	ResultName = "Result"
)

type Operator struct {
	tree.Page
	Items map[string]*OperatorItem
}

func NewOperator(libs *tree.LibraryManager) *Operator {
	instance := (&Operator{
		Page: tree.NewPageAdaptor(),
		Items: map[string]*OperatorItem{
			AdditionName: NewOperatorItem(func(left concept.Index, right concept.Index) concept.Expression {
				return expression.NewAddition(left, right)
			}),
			DivisionName: NewOperatorItem(func(left concept.Index, right concept.Index) concept.Expression {
				return expression.NewDivision(left, right)
			}),
			MultiplicationName: NewOperatorItem(func(left concept.Index, right concept.Index) concept.Expression {
				return expression.NewMultiplication(left, right)
			}),
			SubtractionName: NewOperatorItem(func(left concept.Index, right concept.Index) concept.Expression {
				return expression.NewSubtraction(left, right)
			}),
		},
	})

	for name, item := range instance.Items {
		instance.SetFunction(variable.NewString(name+FuncName), item.Func)
		instance.SetConst(variable.NewString(name+LeftName), item.Left)
		instance.SetConst(variable.NewString(name+RightName), item.Right)
		instance.SetConst(variable.NewString(name+ResultName), item.Result)
	}
	return instance
}
