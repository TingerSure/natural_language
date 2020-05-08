package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	SetFieldContentName = "object"
	SetFieldKeyName     = "key"
	SetFieldValueName   = "value"

	SetFieldObjectErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetFieldObjectErrorException"))
	SetFieldKeyErrorExceptionTemplate    = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetFieldKeyErrorException"))
	SetFieldKeyNotExistExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetFieldKeyNotExistException"))
)

func initSetField(instance *Object) {
	SetFieldContent := variable.NewString(SetFieldContentName)
	SetFieldKey := variable.NewString(SetFieldKeyName)
	SetFieldValue := variable.NewString(SetFieldValueName)

	SetFieldObjectErrorException := SetFieldObjectErrorExceptionTemplate.Copy()
	SetFieldKeyErrorException := SetFieldKeyErrorExceptionTemplate.Copy()
	SetFieldKeyNotExistException := SetFieldKeyNotExistExceptionTemplate.Copy()

	var SetField concept.Function

	SetField = variable.NewSystemFunction(
		variable.NewString("SetField"),
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
				value = variable.NewNull()
			}

			suspend := object.SetField(key, value)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(SetField)
			}

			return variable.NewParam().Set(SetFieldContent, object), nil
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

	instance.SetException(variable.NewString("SetFieldObjectErrorException"), SetFieldObjectErrorException)
	instance.SetException(variable.NewString("SetFieldKeyErrorException"), SetFieldKeyErrorException)
	instance.SetException(variable.NewString("SetFieldKeyNotExistException"), SetFieldKeyNotExistException)

	instance.SetConst(variable.NewString("SetFieldContent"), SetFieldContent)
	instance.SetConst(variable.NewString("SetFieldKey"), SetFieldKey)
	instance.SetConst(variable.NewString("SetFieldValue"), SetFieldValue)

	instance.SetFunction(variable.NewString("SetField"), SetField)

}
