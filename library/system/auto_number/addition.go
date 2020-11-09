package auto_number

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	AdditionTargetName = "target"
	AdditionResultName = "result"
	AdditionKeyName    = "addition"
)

func initAddition(libs *runtime.LibraryManager, instance *AutoNumber) {

	AdditionTarget := libs.Sandbox.Variable.String.New(AdditionTargetName)
	AdditionResult := libs.Sandbox.Variable.String.New(AdditionResultName)
	AdditionKey := libs.Sandbox.Variable.String.New(AdditionKeyName)
	AdditionTargetNotExistException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "AdditionTargetNotExistException")
	AdditionAutoObjectValueTypeErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "AdditionAutoObjectValueTypeErrorException")

	var Addition concept.Function = nil

	anticipateNumber := instance.New(libs.Sandbox.Variable.Number.New(0))

	Addition = libs.Sandbox.Variable.SystemFunction.New(
		AdditionKey.Clone(),
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			preLeft, suspend := object.GetField(instance.AutoNumberClassValue)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend
			}

			left, ok := variable.VariableFamilyInstance.IsNumber(preLeft)
			if !ok {
				return nil, AdditionAutoObjectValueTypeErrorException.Copy().AddStack(Addition)
			}

			right, ok := variable.VariableFamilyInstance.IsNumber(input.Get(AdditionTarget))
			if !ok {
				return nil, AdditionTargetNotExistException.Copy().AddStack(Addition)
			}

			return libs.Sandbox.Variable.Param.New().Set(AdditionResult, libs.Sandbox.Variable.Number.New(left.Value()+right.Value())), nil
		},
		func(_ concept.Param, _ concept.Object) concept.Param {

			return libs.Sandbox.Variable.Param.New().Set(AdditionResult, anticipateNumber)
		},
		[]concept.String{
			AdditionTarget,
		},
		[]concept.String{
			AdditionResult,
		},
	)

	instance.AutoNumberClass.SetMethod(AdditionKey, Addition)

	instance.SetException(libs.Sandbox.Variable.String.New("AdditionTargetNotExistException"), AdditionTargetNotExistException)
	instance.SetException(libs.Sandbox.Variable.String.New("AdditionAutoObjectValueTypeErrorException"), AdditionAutoObjectValueTypeErrorException)

	instance.SetConst(libs.Sandbox.Variable.String.New("AdditionTarget"), AdditionTarget)
	instance.SetConst(libs.Sandbox.Variable.String.New("AdditionResult"), AdditionResult)
	instance.SetConst(libs.Sandbox.Variable.String.New("AdditionKey"), AdditionKey)
}
