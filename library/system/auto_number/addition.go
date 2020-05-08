package auto_number

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	AdditionTargetNotExistExceptionTemplate           = interrupt.NewException(variable.NewString("type error"), variable.NewString("AdditionTargetNotExistException"))
	AdditionAutoObjectValueTypeErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("AdditionAutoObjectValueTypeErrorException"))

	AdditionTargetName = "target"
	AdditionResultName = "result"
	AdditionKeyName    = "addition"
)

func initAddition(instance *AutoNumber) {

	AdditionTarget := variable.NewString(AdditionTargetName)
	AdditionResult := variable.NewString(AdditionResultName)
	AdditionKey := variable.NewString(AdditionKeyName)
	AdditionTargetNotExistException := AdditionTargetNotExistExceptionTemplate.Copy()
	AdditionAutoObjectValueTypeErrorException := AdditionAutoObjectValueTypeErrorExceptionTemplate.Copy()

	var Addition concept.Function = nil

	Addition = variable.NewSystemFunction(
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

			return variable.NewParam().Set(AdditionResult, variable.NewNumber(left.Value()+right.Value())), nil
		},
		[]concept.String{
			AdditionTarget,
		},
		[]concept.String{
			AdditionResult,
		},
	)

	instance.AutoNumberClass.SetMethod(AdditionKey, Addition)

	instance.SetException(variable.NewString("AdditionTargetNotExistException"), AdditionTargetNotExistException)
	instance.SetException(variable.NewString("AdditionAutoObjectValueTypeErrorException"), AdditionAutoObjectValueTypeErrorException)

	instance.SetConst(variable.NewString("AdditionTarget"), AdditionTarget)
	instance.SetConst(variable.NewString("AdditionResult"), AdditionResult)
	instance.SetConst(variable.NewString("AdditionKey"), AdditionKey)
}
