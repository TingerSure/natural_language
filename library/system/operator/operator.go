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

func (o *Operator) EqualItem(resultFormat func(result bool) concept.Bool, anticipateValue concept.Variable) *OperatorItem {
	instance := &OperatorItem{
		Left:   o.Libs.Sandbox.Variable.String.New(OperatorLeft),
		Right:  o.Libs.Sandbox.Variable.String.New(OperatorRight),
		Result: o.Libs.Sandbox.Variable.String.New(OperatorResult),
	}
	instance.Func = o.Libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, object concept.Variable) (param concept.Param, exception concept.Exception) {
			param = o.Libs.Sandbox.Variable.Param.New()
			leftPre := input.Get(instance.Left)
			rightPre := input.Get(instance.Right)
			if leftPre.IsNull() && rightPre.IsNull() {
				param.Set(instance.Result, resultFormat(true))
				return
			}
			if leftPre.IsNull() || rightPre.IsNull() {
				param.Set(instance.Result, resultFormat(false))
				return
			}
			leftString, yesLeft := variable.VariableFamilyInstance.IsStringHome(leftPre)
			rightString, yesRight := variable.VariableFamilyInstance.IsStringHome(rightPre)
			if yesLeft && yesRight {
				param.Set(instance.Result, resultFormat(leftString.Value() == rightString.Value()))
				return
			}
			leftNumber, yesLeft := variable.VariableFamilyInstance.IsNumber(leftPre)
			rightNumber, yesRight := variable.VariableFamilyInstance.IsNumber(rightPre)
			if yesLeft && yesRight {
				param.Set(instance.Result, resultFormat(leftNumber.Value() == rightNumber.Value()))
				return
			}
			leftBool, yesLeft := variable.VariableFamilyInstance.IsBool(leftPre)
			rightBool, yesRight := variable.VariableFamilyInstance.IsBool(rightPre)
			if yesLeft && yesRight {
				param.Set(instance.Result, resultFormat(leftBool.Value() == rightBool.Value()))
				return
			}
			param.Set(instance.Result, resultFormat(leftPre == rightPre))
			return
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
				return nil, o.OperatorParamMissingExceptionTemplate.Copy()
			}

			left, yesLeft := variable.VariableFamilyInstance.IsNumber(leftPre)
			right, yesRight := variable.VariableFamilyInstance.IsNumber(rightPre)
			if !yesLeft || !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy()
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy()
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
				return nil, o.OperatorParamMissingExceptionTemplate.Copy()
			}

			left, yesLeft := variable.VariableFamilyInstance.IsBool(leftPre)
			right, yesRight := variable.VariableFamilyInstance.IsBool(rightPre)
			if !yesLeft || !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy()
			}
			result, exception := exec(left, right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy()
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
				return nil, o.OperatorParamMissingExceptionTemplate.Copy()
			}
			right, yesRight := variable.VariableFamilyInstance.IsBool(rightPre)

			if !yesRight {
				return nil, o.OperatorTypeErrorExceptionTemplate.Copy()
			}
			result, exception := exec(right)
			if !nl_interface.IsNil(exception) {
				return nil, exception.Copy()
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
	AdditionName       = "addition"
	DivisionName       = "division"
	MultiplicationName = "multiplication"
	SubtractionName    = "subtraction"
	EqualName          = "equal"
	NotEqualName       = "notEqual"
	GreaterName        = "greater"
	LessName           = "less"
	GreaterOrEqualName = "greaterOrEqual"
	LessOrEqualName    = "lessOrEqual"
	OrName             = "or"
	AndName            = "and"
	NotName            = "not"
)

type Operator struct {
	concept.Page
	NumberItems                           map[string]*OperatorItem
	EqualItems                            map[string]*OperatorItem
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
		GreaterName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() > right.Value()), nil
		}, anticipateBool),
		LessName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() < right.Value()), nil
		}, anticipateBool),
		GreaterOrEqualName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() >= right.Value()), nil
		}, anticipateBool),
		LessOrEqualName: instance.NewNumberOperatorNumberItem(func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Exception) {
			return libs.Sandbox.Variable.Bool.New(left.Value() <= right.Value()), nil
		}, anticipateBool),
	}

	instance.EqualItems = map[string]*OperatorItem{
		EqualName: instance.EqualItem(func(result bool) concept.Bool {
			return libs.Sandbox.Variable.Bool.New(result)
		}, anticipateBool),
		NotEqualName: instance.EqualItem(func(result bool) concept.Bool {
			return libs.Sandbox.Variable.Bool.New(!result)
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

	for name, item := range instance.EqualItems {
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
