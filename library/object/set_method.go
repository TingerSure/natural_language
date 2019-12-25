package object

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

const (
	SetMethodContent  = "object"
	SetMethodKey      = "key"
	SetMethodFunction = "function"
)

var (
	SetMethodObjectErrorException   = interrupt.NewException("type error", "SetMethodObjectErrorException")
	SetMethodKeyErrorException      = interrupt.NewException("type error", "SetMethodKeyErrorException")
	SetMethodFunctionErrorException = interrupt.NewException("type error", "SetMethodFunctionErrorException")
)

var (
	SetMethod *variable.SystemFunction = nil
)

func init() {
	SetMethod = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(SetMethodContent))
			if !ok {
				return nil, SetMethodObjectErrorException.Copy().AddStack(SetMethod)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(SetMethodKey))
			if !ok {
				return nil, SetMethodKeyErrorException.Copy().AddStack(SetMethod)
			}

			function, ok := variable.VariableFamilyInstance.IsFunctionHome(input.Get(SetMethodFunction))
			if !ok {
				return nil, SetMethodFunctionErrorException.Copy().AddStack(SetMethod)
			}

			suspend := object.SetMethod(key.Value(), function)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(SetMethod)
			}

			return variable.NewParamWithInit(map[string]concept.Variable{
				SetMethodContent: object,
			}), nil
		},
		[]string{
			SetMethodContent,
			SetMethodKey,
			SetMethodFunction,
		},
		[]string{
			SetMethodContent,
		},
	)
}
