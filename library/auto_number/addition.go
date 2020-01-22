package auto_number

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

var (
	AdditionTargetNotExistException           = interrupt.NewException("type error", "AdditionTargetNotExistException")
	AdditionAutoObjectValueTypeErrorException = interrupt.NewException("type error", "AdditionAutoObjectValueTypeErrorException")

	AdditionTarget = "target"
	AdditionResult = "result"

	Addition    concept.Function = nil
	AdditionKey                  = "addition"
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

			return variable.NewParamWithInit(map[string]concept.Variable{
				AdditionResult: variable.NewNumber(left.Value() + right.Value()),
			}), nil

		},
		[]string{
			AdditionTarget,
		},
		[]string{
			AdditionResult,
		},
	)

	AutoNumberClass.SetMethod(AdditionKey, Addition)
}
