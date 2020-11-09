package operator

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

const (
	OperatorLeft   = "left"
	OperatorRight  = "right"
	OperatorResult = "result"
)

type OperatorItem struct {
	Name   string
	Func   *variable.SystemFunction
	Left   concept.String
	Right  concept.String
	Result concept.String
}

type OperatorUnaryItem struct {
	Name   string
	Func   *variable.SystemFunction
	Right  concept.String
	Result concept.String
}

func (o *Operator) NewNumberOperatorNumberItem(name string, exec func(*variable.Number, *variable.Number) (concept.Variable, concept.Exception), anticipateValue concept.Variable) *OperatorItem {
	instance := &OperatorItem{
		Left:   o.Libs.Sandbox.Variable.String.New(OperatorLeft),
		Right:  o.Libs.Sandbox.Variable.String.New(OperatorRight),
		Result: o.Libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = o.Libs.Sandbox.Variable.SystemFunction.New(
		o.Libs.Sandbox.Variable.String.New(name),
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			left, yesLeft := variable.VariableFamilyInstance.IsNumber(input.Get(instance.Left))
			right, yesRight := variable.VariableFamilyInstance.IsNumber(input.Get(instance.Right))
			if !yesLeft || !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			return o.Libs.Sandbox.Variable.Param.New().Set(instance.Result, result), nil
		},
		func(input concept.Param, object concept.Object) concept.Param {
			return o.Libs.Sandbox.Variable.Param.New().Set(instance.Result, anticipateValue)
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

func (o *Operator) NewBoolOperatorBoolItem(name string, exec func(*variable.Bool, *variable.Bool) (concept.Variable, concept.Exception), anticipateValue concept.Variable) *OperatorItem {
	instance := &OperatorItem{
		Left:   o.Libs.Sandbox.Variable.String.New(OperatorLeft),
		Right:  o.Libs.Sandbox.Variable.String.New(OperatorRight),
		Result: o.Libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = o.Libs.Sandbox.Variable.SystemFunction.New(
		o.Libs.Sandbox.Variable.String.New(name),
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			left, yesLeft := variable.VariableFamilyInstance.IsBool(input.Get(instance.Left))
			right, yesRight := variable.VariableFamilyInstance.IsBool(input.Get(instance.Right))
			if !yesLeft || !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			return o.Libs.Sandbox.Variable.Param.New().Set(instance.Result, result), nil
		},
		func(input concept.Param, object concept.Object) concept.Param {
			return o.Libs.Sandbox.Variable.Param.New().Set(instance.Result, anticipateValue)
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

func (o *Operator) NewOperatorBoolItem(name string, exec func(*variable.Bool) (concept.Variable, concept.Exception), anticipateValue concept.Variable) *OperatorUnaryItem {
	instance := &OperatorUnaryItem{
		Right:  o.Libs.Sandbox.Variable.String.New(OperatorRight),
		Result: o.Libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = o.Libs.Sandbox.Variable.SystemFunction.New(
		o.Libs.Sandbox.Variable.String.New(name),
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			right, yesRight := variable.VariableFamilyInstance.IsBool(input.Get(instance.Right))
			if !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			return o.Libs.Sandbox.Variable.Param.New().Set(instance.Result, result), nil
		},
		func(input concept.Param, object concept.Object) concept.Param {
			return o.Libs.Sandbox.Variable.Param.New().Set(instance.Result, anticipateValue)
		},
		[]concept.String{
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
	NotName                  = "Not"
)

const (
	FuncName   = "Func"
	LeftName   = "Left"
	RightName  = "Right"
	ResultName = "Result"
)

type Operator struct {
	tree.Page
	Items                                map[string]*OperatorItem
	UnaryItems                           map[string]*OperatorUnaryItem
	Libs                                 *runtime.LibraryManager
	OperatorTypeErrorExceptionTemplate   concept.Exception
	OperatorDivisorZeroExceptionTemplate concept.Exception
}

func NewOperator(libs *runtime.LibraryManager) *Operator {
	instance := &Operator{
		Page:                                 tree.NewPageAdaptor(libs.Sandbox),
		OperatorTypeErrorExceptionTemplate:   libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "OperatorTypeErrorException"),
		OperatorDivisorZeroExceptionTemplate: libs.Sandbox.Interrupt.Exception.NewOriginal("param error", "OperatorDivisorZeroException"),
		Libs:                                 libs,
	}

	anticipateNumber := libs.Sandbox.Variable.Number.New(0)
	anticipateBool := libs.Sandbox.Variable.Bool.New(false)

	instance.Items = map[string]*OperatorItem{

		AdditionName: instance.NewNumberOperatorNumberItem(AdditionName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Number.New(left.Value() + right.Value()), nil
		}, anticipateNumber),
		DivisionName: instance.NewNumberOperatorNumberItem(DivisionName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			if right.Value() == 0 {
				return nil, instance.OperatorDivisorZeroExceptionTemplate.Copy()
			}
			return libs.Sandbox.Variable.Number.New(left.Value() / right.Value()), nil
		}, anticipateNumber),
		MultiplicationName: instance.NewNumberOperatorNumberItem(MultiplicationName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Number.New(left.Value() * right.Value()), nil
		}, anticipateNumber),
		SubtractionName: instance.NewNumberOperatorNumberItem(SubtractionName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Number.New(left.Value() - right.Value()), nil
		}, anticipateNumber),
		EqualToName: instance.NewNumberOperatorNumberItem(EqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() == right.Value()), nil
		}, anticipateBool),
		NotEqualToName: instance.NewNumberOperatorNumberItem(NotEqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() != right.Value()), nil
		}, anticipateBool),
		GreaterThanName: instance.NewNumberOperatorNumberItem(GreaterThanName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() > right.Value()), nil
		}, anticipateBool),
		LessThanName: instance.NewNumberOperatorNumberItem(LessThanName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() < right.Value()), nil
		}, anticipateBool),
		GreaterThanOrEqualToName: instance.NewNumberOperatorNumberItem(GreaterThanOrEqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() >= right.Value()), nil
		}, anticipateBool),
		LessThanOrEqualToName: instance.NewNumberOperatorNumberItem(LessThanOrEqualToName, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() <= right.Value()), nil
		}, anticipateBool),
		OrName: instance.NewBoolOperatorBoolItem(OrName, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() || right.Value()), nil
		}, anticipateBool),
		AndName: instance.NewBoolOperatorBoolItem(AndName, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() && right.Value()), nil
		}, anticipateBool),
	}

	instance.UnaryItems = map[string]*OperatorUnaryItem{
		NotName: instance.NewOperatorBoolItem(NotName, func(right *variable.Bool) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(!right.Value()), nil
		}, anticipateBool),
	}

	for name, item := range instance.Items {
		instance.SetFunction(libs.Sandbox.Variable.String.New(name+FuncName), item.Func)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+LeftName), item.Left)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+RightName), item.Right)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+ResultName), item.Result)
	}

	for name, item := range instance.UnaryItems {
		instance.SetFunction(libs.Sandbox.Variable.String.New(name+FuncName), item.Func)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+RightName), item.Right)
		instance.SetConst(libs.Sandbox.Variable.String.New(name+ResultName), item.Result)
	}

	instance.SetException(libs.Sandbox.Variable.String.New("OperatorTypeErrorException"), instance.OperatorTypeErrorExceptionTemplate)
	instance.SetException(libs.Sandbox.Variable.String.New("OperatorDivisorZeroException"), instance.OperatorDivisorZeroExceptionTemplate)
	return instance
}
