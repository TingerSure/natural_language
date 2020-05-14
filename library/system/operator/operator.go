package operator

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

const (
	ItemLeft   = "left"
	ItemRight  = "right"
	ItemResult = "result"
)

var (
	OperatorTypeErrorExceptionTemplate   = interrupt.NewException(variable.NewString("type error"), variable.NewString("OperatorTypeErrorException"))
	OperatorDivisorZeroExceptionTemplate = interrupt.NewException(variable.NewString("param error"), variable.NewString("OperatorDivisorZeroException"))
)

type OperatorItem struct {
	Name   string
	Func   *variable.SystemFunction
	Left   concept.String
	Right  concept.String
	Result concept.String
}

func NewOperatorItem(name string, exec func(*variable.Number, *variable.Number) (*variable.Number, concept.Exception)) *OperatorItem {
	instance := &OperatorItem{
		Left:   variable.NewString(ItemLeft),
		Right:  variable.NewString(ItemRight),
		Result: variable.NewString(ItemResult),
	}
	instance.Func = variable.NewSystemFunction(
		variable.NewString(name),
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			left, yesLeft := variable.VariableFamilyInstance.IsNumber(input.Get(instance.Left))
			right, yesRight := variable.VariableFamilyInstance.IsNumber(input.Get(instance.Right))
			if !yesLeft || !yesRight {
				return nil, OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			return variable.NewParam().Set(instance.Result, result), nil
		},
		[]concept.String{
			instance.Left,
			instance.Right,
		},
		[]concept.String{
			instance.Result,
		},
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

			AdditionName: NewOperatorItem(AdditionName, func(left *variable.Number, right *variable.Number) (*variable.Number, concept.Exception) {
				return variable.NewNumber(left.Value() + right.Value()), nil
			}),
			DivisionName: NewOperatorItem(DivisionName, func(left *variable.Number, right *variable.Number) (*variable.Number, concept.Exception) {
				if right.Value() == 0 {
					return nil, OperatorDivisorZeroExceptionTemplate
				}
				return variable.NewNumber(left.Value() / right.Value()), nil
			}),
			MultiplicationName: NewOperatorItem(MultiplicationName, func(left *variable.Number, right *variable.Number) (*variable.Number, concept.Exception) {
				return variable.NewNumber(left.Value() * right.Value()), nil
			}),
			SubtractionName: NewOperatorItem(SubtractionName, func(left *variable.Number, right *variable.Number) (*variable.Number, concept.Exception) {
				return variable.NewNumber(left.Value() - right.Value()), nil
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
