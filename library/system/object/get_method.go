package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	GetMethodContent  = variable.NewString("object")
	GetMethodKey      = variable.NewString("key")
	GetMethodFunction = variable.NewString("function")
)

var (
	GetMethodObjectErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetMethodObjectErrorException"))
	GetMethodKeyErrorException    = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetMethodKeyErrorException"))
	GetMethodKeyNotExistException = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetMethodKeyNotExistException"))
)

var (
	GetMethod *variable.SystemFunction = nil
)

func init() {
	GetMethod = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(GetMethodContent))
			if !ok {
				return nil, GetMethodObjectErrorException.Copy().AddStack(GetMethod)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(GetMethodKey))
			if !ok {
				return nil, GetMethodKeyErrorException.Copy().AddStack(GetMethod)
			}
			if !object.HasMethod(key) {
				return nil, GetMethodKeyNotExistException.Copy().AddStack(GetMethod)
			}

			function, suspend := object.GetMethod(key)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(GetMethod)
			}

			return variable.NewParam().Set(GetMethodFunction, function), nil
		},
		[]concept.String{
			GetMethodContent,
			GetMethodKey,
		},
		[]concept.String{
			GetMethodFunction,
		},
	)
}
