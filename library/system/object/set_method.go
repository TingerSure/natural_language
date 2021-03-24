package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	SetMethodContentName  = "object"
	SetMethodKeyName      = "key"
	SetMethodFunctionName = "function"
)

func initSetMethod(libs *tree.LibraryManager, instance *Object) {
	SetMethodContent := libs.Sandbox.Variable.String.New(SetMethodContentName)
	SetMethodKey := libs.Sandbox.Variable.String.New(SetMethodKeyName)
	SetMethodFunction := libs.Sandbox.Variable.String.New(SetMethodFunctionName)

	SetMethodObjectErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "SetMethodObjectErrorException")
	SetMethodKeyErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "SetMethodKeyErrorException")
	SetMethodFunctionErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "SetMethodFunctionErrorException")

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
		func(input concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(SetMethodContent, input.Get(SetMethodContent))
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
