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

	GetMethodObjectErrorExceptionTemplate = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("GetMethodObjectErrorException"))
	GetMethodKeyErrorExceptionTemplate    = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("GetMethodKeyErrorException"))
	GetMethodKeyNotExistExceptionTemplate = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("GetMethodKeyNotExistException"))
)

func initGetMethod(instance *Object) {
	GetMethodContent := libs.Sandbox.Variable.String.New(GetMethodContentName)
	GetMethodKey := libs.Sandbox.Variable.String.New(GetMethodKeyName)
	GetMethodFunction := libs.Sandbox.Variable.String.New(GetMethodFunctionName)

	GetMethodObjectErrorException := GetMethodObjectErrorExceptionTemplate.Copy()
	GetMethodKeyErrorException := GetMethodKeyErrorExceptionTemplate.Copy()
	GetMethodKeyNotExistException := GetMethodKeyNotExistExceptionTemplate.Copy()

	var GetMethod concept.Function
	GetMethod = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("GetMethod"),
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

			return libs.Sandbox.Variable.Param.New().Set(GetMethodFunction, function), nil
		},
		[]concept.String{
			GetMethodContent,
			GetMethodKey,
		},
		[]concept.String{
			GetMethodFunction,
		},
	)

	instance.SetException(libs.Sandbox.Variable.String.New("GetMethodObjectErrorException"), GetMethodObjectErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("GetMethodKeyErrorException"), GetMethodKeyErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("GetMethodKeyNotExistException"), GetMethodKeyNotExistException)

	instance.SetConst(libs.Sandbox.Variable.String.New("GetMethodContent"), GetMethodContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("GetMethodKey"), GetMethodKey)
	instance.SetConst(libs.Sandbox.Variable.String.New("GetMethodFunction"), GetMethodFunction)

	instance.SetFunction(libs.Sandbox.Variable.String.New("GetMethod"), GetMethod)

}
