package operator

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
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

func (o *Operator) NewNumberOperatorNumberItem(exec func(*variable.Number, *variable.Number) (concept.Variable, concept.Exception), anticipateValue concept.Variable) *OperatorItem {
	instance := &OperatorItem{
		Left:   o.Libs.Sandbox.Variable.String.New(OperatorLeft),
		Right:  o.Libs.Sandbox.Variable.String.New(OperatorRight),
		Result: o.Libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = o.Libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
			leftPre := input.Get(instance.Left)
			rightPre := input.Get(instance.Right)
			if leftPre.IsNull() || rightPre.IsNull() {
				return nil, o.OperatorParamMissingExceptionTemplate.Copy().AddStack(instance.Func)
			}

			left, yesLeft := variable.VariableFamilyInstance.IsNumber(leftPre)
			right, yesRight := variable.VariableFamilyInstance.IsNumber(rightPre)
			if !yesLeft || !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			param := o.Libs.Sandbox.Variable.Param.New()
			param.Set(instance.Result, result)
			return param, nil
		},
		func(input concept.Param, object concept.Variable) concept.Param {
			param := o.Libs.Sandbox.Variable.Param.New()
			param.Set(instance.Result, anticipateValue)
			return param
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

func (o *Operator) NewBoolOperatorBoolItem(exec func(*variable.Bool, *variable.Bool) (concept.Variable, concept.Exception), anticipateValue concept.Variable) *OperatorItem {
	instance := &OperatorItem{
		Left:   o.Libs.Sandbox.Variable.String.New(OperatorLeft),
		Right:  o.Libs.Sandbox.Variable.String.New(OperatorRight),
		Result: o.Libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = o.Libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
			leftPre := input.Get(instance.Left)
			rightPre := input.Get(instance.Right)
			if leftPre.IsNull() || rightPre.IsNull() {
				return nil, o.OperatorParamMissingExceptionTemplate.Copy().AddStack(instance.Func)
			}

			left, yesLeft := variable.VariableFamilyInstance.IsBool(leftPre)
			right, yesRight := variable.VariableFamilyInstance.IsBool(rightPre)
			if !yesLeft || !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			param := o.Libs.Sandbox.Variable.Param.New()
			param.Set(instance.Result, result)
			return param, nil
		},
		func(input concept.Param, object concept.Variable) concept.Param {
			param := o.Libs.Sandbox.Variable.Param.New()
			param.Set(instance.Result, anticipateValue)
			return param
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

func (o *Operator) NewOperatorBoolItem(exec func(*variable.Bool) (concept.Variable, concept.Exception), anticipateValue concept.Variable) *OperatorUnaryItem {
	instance := &OperatorUnaryItem{
		Right:  o.Libs.Sandbox.Variable.String.New(OperatorRight),
		Result: o.Libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = o.Libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
			rightPre := input.Get(instance.Right)
			if rightPre.IsNull() {
				return nil, o.OperatorParamMissingExceptionTemplate.Copy().AddStack(instance.Func)
			}
			right, yesRight := variable.VariableFamilyInstance.IsBool(rightPre)

			if !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy().AddStack(instance.Func)
			}
			result, exception := exec(right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy().AddStack(instance.Func)
			}
			param := o.Libs.Sandbox.Variable.Param.New()
			param.Set(instance.Result, result)
			return param, nil
		},
		func(input concept.Param, object concept.Variable) concept.Param {
			param := o.Libs.Sandbox.Variable.Param.New()
			param.Set(instance.Result, anticipateValue)
			return param
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
	concept.Page
	NumberItems                           map[string]*OperatorItem
	BoolItems                             map[string]*OperatorItem
	BoolUnaryItems                        map[string]*OperatorUnaryItem
	Libs                                  *tree.LibraryManager
	OperatorTypeErrorExceptionTemplate    concept.Exception
	OperatorParamMissingExceptionTemplate concept.Exception
	OperatorDivisorZeroExceptionTemplate  concept.Exception
}

func NewOperator(libs *tree.LibraryManager) *Operator {
	instance := &Operator{
		Page:                                  libs.Sandbox.Variable.Page.New(),
		OperatorTypeErrorExceptionTemplate:    libs.Sandbox.Variable.Exception.NewOriginal("type error", "OperatorTypeErrorException"),
		OperatorParamMissingExceptionTemplate: libs.Sandbox.Variable.Exception.NewOriginal("type error", "OperatorParamMissingException"),
		OperatorDivisorZeroExceptionTemplate:  libs.Sandbox.Variable.Exception.NewOriginal("param error", "OperatorDivisorZeroException"),
		Libs:                                  libs,
	}

	anticipateNumber := libs.Sandbox.Variable.Number.New(0)
	anticipateBool := libs.Sandbox.Variable.Bool.New(false)

	instance.NumberItems = map[string]*OperatorItem{

		AdditionName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Number.New(left.Value() + right.Value()), nil
		}, anticipateNumber),
		DivisionName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			if right.Value() == 0 {
				return nil, instance.OperatorDivisorZeroExceptionTemplate.Copy()
			}
			return libs.Sandbox.Variable.Number.New(left.Value() / right.Value()), nil
		}, anticipateNumber),
		MultiplicationName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Number.New(left.Value() * right.Value()), nil
		}, anticipateNumber),
		SubtractionName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Number.New(left.Value() - right.Value()), nil
		}, anticipateNumber),
		EqualToName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() == right.Value()), nil
		}, anticipateBool),
		NotEqualToName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() != right.Value()), nil
		}, anticipateBool),
		GreaterThanName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() > right.Value()), nil
		}, anticipateBool),
		LessThanName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() < right.Value()), nil
		}, anticipateBool),
		GreaterThanOrEqualToName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() >= right.Value()), nil
		}, anticipateBool),
		LessThanOrEqualToName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() <= right.Value()), nil
		}, anticipateBool),
	}

	instance.BoolItems = map[string]*OperatorItem{
		OrName: instance.NewBoolOperatorBoolItem(func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() || right.Value()), nil
		}, anticipateBool),
		AndName: instance.NewBoolOperatorBoolItem(func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() && right.Value()), nil
		}, anticipateBool),
	}

	instance.BoolUnaryItems = map[string]*OperatorUnaryItem{
		NotName: instance.NewOperatorBoolItem(func(right *variable.Bool) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(!right.Value()), nil
		}, anticipateBool),
	}

	for name, item := range instance.NumberItems {
		instance.SetPublic(libs.Sandbox.Variable.String.New(name), libs.Sandbox.Index.PublicIndex.New(name, libs.Sandbox.Index.ConstIndex.New(item.Func)))
	}

	for name, item := range instance.BoolItems {
		instance.SetPublic(libs.Sandbox.Variable.String.New(name), libs.Sandbox.Index.PublicIndex.New(name, libs.Sandbox.Index.ConstIndex.New(item.Func)))
	}

	for name, item := range instance.BoolUnaryItems {
		instance.SetPublic(libs.Sandbox.Variable.String.New(name), libs.Sandbox.Index.PublicIndex.New(name, libs.Sandbox.Index.ConstIndex.New(item.Func)))
	}

	return instance
}
