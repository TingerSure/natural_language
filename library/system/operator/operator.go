package operator

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

const (
	OperatorLeft   = "left"
	OperatorRight  = "right"
	OperatorResult = "result"
)

var (
	OperatorTypeErrorExceptionTemplate   = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("OperatorTypeErrorException"))
	OperatorDivisorZeroExceptionTemplate = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("param error"), libs.Sandbox.Variable.String.New("OperatorDivisorZeroException"))
)

type OperatorItem struct {
	Name   string
	Func   *variable.SystemFunction
	Left   concept.String
	Right  concept.String
	Result concept.String
}

func NewNumberOperatorNumberItem(name string, exec func(*variable.Number, *variable.Number) (concept.Variable, concept.Exception)) *OperatorItem {
	instance := &OperatorItem{
		Left:   libs.Sandbox.Variable.String.New(OperatorLeft),
		Right:  libs.Sandbox.Variable.String.New(OperatorRight),
		Result: libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New(name),
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
			return libs.Sandbox.Variable.Param.New().Set(instance.Result, result), nil
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

func NewBoolOperatorBoolItem(name string, exec func(*variable.Bool, *variable.Bool) (concept.Variable, concept.Exception)) *OperatorItem {
	instance := &OperatorItem{
		Left:   libs.Sandbox.Variable.String.New(OperatorLeft),
		Right:  libs.Sandbox.Variable.String.New(OperatorRight),
		Result: libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New(name),
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			left, yesLeft := variable.VariableFamilyInstance.IsBool(input.Get(instance.Left))
			right, yesRight := variable.VariableFamilyInstance.IsBool(input.Get(instance.Right))
			if !yesLeft || !yesRight {
				return nil, OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			return libs.Sandbox.Variable.Param.New().Set(instance.Result, result), nil
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
	AdditionName             = "Addition"
	DivisionName             = "Division"
	MultiplicationName       = "Multiplication"
	SubtractionName          = "Subtraction"
	EqualToName              = "EqualTo"
	NotEqualToName           = "NotEqualTo"
	GreaterThanName          = "GreaterThan"
	LessThanName             = "LessThan"
	GreaterThanOrEqualToName = "GreaterThanOrEqualTo"
	LessThanOrEqualToName    = "LessThanOrEqualTo"
	OrName                   = "Or"
	AndName                  = "And"
)

const (
	FuncName   = "Func"
	LeftName   = "Left"
	RightName  = "Right"
	ResultName = "Result"
)

type Operator struct {
	tree.Page
	Items map[string]*OperatorItem
}

func NewOperator(libs *runtime.LibraryManager) *Operator {
	instance := (&Operator{
		Page: tree.NewPageAdaptor(libs.Sandbox),
		Items: map[string]*OperatorItem{

			AdditionName: NewNumberOperatorNumberItem(AdditionName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return libs.Sandbox.Variable.Number.New(left.Value() + right.Value()), nil
			}),
			DivisionName: NewNumberOperatorNumberItem(DivisionName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				if right.Value() == 0 {
					return nil, OperatorDivisorZeroExceptionTemplate
				}
				return libs.Sandbox.Variable.Number.New(left.Value() / right.Value()), nil
			}),
			MultiplicationName: NewNumberOperatorNumberItem(MultiplicationName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return libs.Sandbox.Variable.Number.New(left.Value() * right.Value()), nil
			}),
			SubtractionName: NewNumberOperatorNumberItem(SubtractionName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return libs.Sandbox.Variable.Number.New(left.Value() - right.Value()), nil
			}),
			EqualToName: NewNumberOperatorNumberItem(EqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() == right.Value()), nil
			}),
			NotEqualToName: NewNumberOperatorNumberItem(NotEqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() != right.Value()), nil
			}),
			GreaterThanName: NewNumberOperatorNumberItem(GreaterThanName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() > right.Value()), nil
			}),
			LessThanName: NewNumberOperatorNumberItem(LessThanName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() < right.Value()), nil
			}),
			GreaterThanOrEqualToName: NewNumberOperatorNumberItem(GreaterThanOrEqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() >= right.Value()), nil
			}),
			LessThanOrEqualToName: NewNumberOperatorNumberItem(LessThanOrEqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() <= right.Value()), nil
			}),
			OrName: NewBoolOperatorBoolItem(OrName, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() || right.Value()), nil
			}),
			AndName: NewBoolOperatorBoolItem(AndName, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Exception) {
				return variable.NewBool(left.Value() && right.Value()), nil
			}),
		},
	})

	for name, item := range instance.Items {
		instance.SetFunction(libs.Sandbox.Variable.String.New(name+FuncName), item.Func)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+LeftName), item.Left)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+RightName), item.Right)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+ResultName), item.Result)
	}

	instance.SetException(libs.Sandbox.Variable.String.New("OperatorTypeErrorException"), OperatorTypeErrorExceptionTemplate)
	instance.SetException(libs.Sandbox.Variable.String.New("OperatorDivisorZeroException"), OperatorDivisorZeroExceptionTemplate)
	return instance
}
