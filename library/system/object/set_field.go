package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	SetFieldContentName = "object"
	SetFieldKeyName     = "key"
	SetFieldValueName   = "value"
)

func initSetField(libs *tree.LibraryManager, instance *Object) {
	SetFieldContent := libs.Sandbox.Variable.String.New(SetFieldContentName)
	SetFieldKey := libs.Sandbox.Variable.String.New(SetFieldKeyName)
	SetFieldValue := libs.Sandbox.Variable.String.New(SetFieldValueName)

	SetFieldObjectErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "SetFieldObjectErrorException")
	SetFieldKeyErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "SetFieldKeyErrorException")
	SetFieldKeyNotExistException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "SetFieldKeyNotExistException")

	var SetField concept.Function

	SetField = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("SetField"),
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(SetFieldContent))
			if !ok {
				return nil, SetFieldObjectErrorException.Copy().AddStack(SetField)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(SetFieldKey))
			if !ok {
				return nil, SetFieldKeyErrorException.Copy().AddStack(SetField)
			}
			if !object.HasField(key) {
				return nil, SetFieldKeyNotExistException.Copy().AddStack(SetField)
			}

			value := input.Get(SetFieldValue)
			if nl_interface.IsNil(value) {
				value = libs.Sandbox.Variable.Null.New()
			}

			suspend := object.SetField(key, value)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(SetField)
			}

			return libs.Sandbox.Variable.Param.New().Set(SetFieldContent, object), nil
		},
		func(input concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(SetFieldContent, input.Get(SetFieldContent))
		},
		[]concept.String{
			SetFieldContent,
			SetFieldKey,
			SetFieldValue,
		},
		[]concept.String{
			SetFieldContent,
		},
	)

	instance.SetException(libs.Sandbox.Variable.String.New("SetFieldObjectErrorException"), SetFieldObjectErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("SetFieldKeyErrorException"), SetFieldKeyErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("SetFieldKeyNotExistException"), SetFieldKeyNotExistException)

	instance.SetConst(libs.Sandbox.Variable.String.New("SetFieldContent"), SetFieldContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("SetFieldKey"), SetFieldKey)
	instance.SetConst(libs.Sandbox.Variable.String.New("SetFieldValue"), SetFieldValue)

	instance.SetFunction(libs.Sandbox.Variable.String.New("SetField"), SetField)

}
