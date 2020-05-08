package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	GetMethodContentName  = "object"
	GetMethodKeyName      = "key"
	GetMethodFunctionName = "function"

	GetMethodObjectErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetMethodObjectErrorException"))
	GetMethodKeyErrorExceptionTemplate    = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetMethodKeyErrorException"))
	GetMethodKeyNotExistExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetMethodKeyNotExistException"))
)

func initGetMethod(instance *Object) {
	GetMethodContent := variable.NewString(GetMethodContentName)
	GetMethodKey := variable.NewString(GetMethodKeyName)
	GetMethodFunction := variable.NewString(GetMethodFunctionName)

	GetMethodObjectErrorException := GetMethodObjectErrorExceptionTemplate.Copy()
	GetMethodKeyErrorException := GetMethodKeyErrorExceptionTemplate.Copy()
	GetMethodKeyNotExistException := GetMethodKeyNotExistExceptionTemplate.Copy()

	var GetMethod concept.Function
	GetMethod = variable.NewSystemFunction(
		variable.NewString("GetMethod"),
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

	instance.SetException(variable.NewString("GetMethodObjectErrorException"), GetMethodObjectErrorException)
	instance.SetException(variable.NewString("GetMethodKeyErrorException"), GetMethodKeyErrorException)
	instance.SetException(variable.NewString("GetMethodKeyNotExistException"), GetMethodKeyNotExistException)

	instance.SetConst(variable.NewString("GetMethodContent"), GetMethodContent)
	instance.SetConst(variable.NewString("GetMethodKey"), GetMethodKey)
	instance.SetConst(variable.NewString("GetMethodFunction"), GetMethodFunction)

	instance.SetFunction(variable.NewString("GetMethod"), GetMethod)

}
