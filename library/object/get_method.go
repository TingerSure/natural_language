package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

const (
	GetMethodContent  = "object"
	GetMethodKey      = "key"
	GetMethodFunction = "function"
)

var (
	GetMethodObjectErrorException = interrupt.NewException("type error", "GetMethodObjectErrorException")
	GetMethodKeyErrorException    = interrupt.NewException("type error", "GetMethodKeyErrorException")
	GetMethodKeyNotExistException = interrupt.NewException("type error", "GetMethodKeyNotExistException")
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
			if !object.HasMethod(key.Value()) {
				return nil, GetMethodKeyNotExistException.Copy().AddStack(GetMethod)
			}

			function, suspend := object.GetMethod(key.Value())
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(GetMethod)
			}

			return variable.NewParamWithInit(map[string]concept.Variable{
				GetMethodFunction: function,
			}), nil
		},
		[]string{
			GetMethodContent,
			GetMethodKey,
		},
		[]string{
			GetMethodFunction,
		},
	)
}
