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

	SetMethodObjectErrorExceptionTemplate   = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetMethodObjectErrorException"))
	SetMethodKeyErrorExceptionTemplate      = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetMethodKeyErrorException"))
	SetMethodFunctionErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetMethodFunctionErrorException"))
)

func initSetMethod(instance *Object) {
	SetMethodContent := variable.NewString(SetMethodContentName)
	SetMethodKey := variable.NewString(SetMethodKeyName)
	SetMethodFunction := variable.NewString(SetMethodFunctionName)

	SetMethodObjectErrorException := SetMethodObjectErrorExceptionTemplate.Copy()
	SetMethodKeyErrorException := SetMethodKeyErrorExceptionTemplate.Copy()
	SetMethodFunctionErrorException := SetMethodFunctionErrorExceptionTemplate.Copy()

	var SetMethod concept.Function

	SetMethod = variable.NewSystemFunction(
		variable.NewString("SetMethod"),
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

	instance.SetException(variable.NewString("SetMethodObjectErrorException"), SetMethodObjectErrorException)
	instance.SetException(variable.NewString("SetMethodKeyErrorException"), SetMethodKeyErrorException)
	instance.SetException(variable.NewString("SetMethodFunctionErrorException"), SetMethodFunctionErrorException)

	instance.SetConst(variable.NewString("SetMethodContent"), SetMethodContent)
	instance.SetConst(variable.NewString("SetMethodKey"), SetMethodKey)
	instance.SetConst(variable.NewString("SetMethodFunction"), SetMethodFunction)

	instance.SetFunction(variable.NewString("SetMethod"), SetMethod)

}
