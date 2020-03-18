package auto_number

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	AdditionTargetNotExistException           = interrupt.NewException(variable.NewString("type error"), variable.NewString("AdditionTargetNotExistException"))
	AdditionAutoObjectValueTypeErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("AdditionAutoObjectValueTypeErrorException"))

	AdditionTarget                  = variable.NewString("target")
	AdditionResult                  = variable.NewString("result")
	AdditionKey                     = variable.NewString("addition")
	Addition       concept.Function = nil
)

func init() {

	Addition = variable.NewSystemFunction(
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			preLeft, suspend := object.GetField(AutoNumberClassValue)
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

	AutoNumberClass.SetMethod(AdditionKey, Addition)
}
