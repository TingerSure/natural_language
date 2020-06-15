package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	SetMethodContentName  = "object"
	SetMethodKeyName      = "key"
	SetMethodFunctionName = "function"

	SetMethodObjectErrorExceptionTemplate   = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("SetMethodObjectErrorException"))
	SetMethodKeyErrorExceptionTemplate      = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("SetMethodKeyErrorException"))
	SetMethodFunctionErrorExceptionTemplate = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("SetMethodFunctionErrorException"))
)

func initSetMethod(instance *Object) {
	SetMethodContent := libs.Sandbox.Variable.String.New(SetMethodContentName)
	SetMethodKey := libs.Sandbox.Variable.String.New(SetMethodKeyName)
	SetMethodFunction := libs.Sandbox.Variable.String.New(SetMethodFunctionName)

	SetMethodObjectErrorException := SetMethodObjectErrorExceptionTemplate.Copy()
	SetMethodKeyErrorException := SetMethodKeyErrorExceptionTemplate.Copy()
	SetMethodFunctionErrorException := SetMethodFunctionErrorExceptionTemplate.Copy()

	var SetMethod concept.Function

	SetMethod = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("SetMethod"),
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

			return libs.Sandbox.Variable.Param.New().Set(SetMethodContent, object), nil
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

	instance.SetException(libs.Sandbox.Variable.String.New("SetMethodObjectErrorException"), SetMethodObjectErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("SetMethodKeyErrorException"), SetMethodKeyErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("SetMethodFunctionErrorException"), SetMethodFunctionErrorException)

	instance.SetConst(libs.Sandbox.Variable.String.New("SetMethodContent"), SetMethodContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("SetMethodKey"), SetMethodKey)
	instance.SetConst(libs.Sandbox.Variable.String.New("SetMethodFunction"), SetMethodFunction)

	instance.SetFunction(libs.Sandbox.Variable.String.New("SetMethod"), SetMethod)

}
