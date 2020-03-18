package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	SetMethodContent  = variable.NewString("object")
	SetMethodKey      = variable.NewString("key")
	SetMethodFunction = variable.NewString("function")
)

var (
	SetMethodObjectErrorException   = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetMethodObjectErrorException"))
	SetMethodKeyErrorException      = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetMethodKeyErrorException"))
	SetMethodFunctionErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetMethodFunctionErrorException"))
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

			suspend := object.SetMethod(key, function)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(SetMethod)
			}

			return variable.NewParam().Set(SetMethodContent, object), nil
		},
		[]concept.String{
			SetMethodContent,
			SetMethodKey,
			SetMethodFunction,
		},
		[]concept.String{
			SetMethodContent,
		},
	)
}
