package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	GetFieldContentName = "object"
	GetFieldKeyName     = "key"
	GetFieldValueName   = "value"

	GetFieldObjectErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetFieldObjectErrorException"))
	GetFieldKeyErrorExceptionTemplate    = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetFieldKeyErrorException"))
	GetFieldKeyNotExistExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetFieldKeyNotExistException"))
)

func initGetField(instance *Object) {
	GetFieldContent := variable.NewString(GetFieldContentName)
	GetFieldKey := variable.NewString(GetFieldKeyName)
	GetFieldValue := variable.NewString(GetFieldValueName)

	GetFieldObjectErrorException := GetFieldObjectErrorExceptionTemplate.Copy()
	GetFieldKeyErrorException := GetFieldKeyErrorExceptionTemplate.Copy()
	GetFieldKeyNotExistException := GetFieldKeyNotExistExceptionTemplate.Copy()

	var GetField concept.Function
	GetField = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			content := input.Get(GetFieldContent)
			var object concept.Object
			if objectHome, ok := variable.VariableFamilyInstance.IsObjectHome(content); ok {
				object = objectHome
			} else if number, ok := variable.VariableFamilyInstance.IsNumber(content); ok {
				object = instance.AutoNumber.NewAutoNumberObject(number)
			} else {
				return nil, GetFieldObjectErrorException.Copy().AddStack(GetField)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(GetFieldKey))
			if !ok {
				return nil, GetFieldKeyErrorException.Copy().AddStack(GetField)
			}
			if !object.HasField(key) {
				return nil, GetFieldKeyNotExistException.Copy().AddStack(GetField)
			}

			value, suspend := object.GetField(key)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(GetField)
			}

			return variable.NewParam().Set(GetFieldValue, value), nil
		},
		[]concept.String{
			GetFieldContent,
			GetFieldKey,
		},
		[]concept.String{
			GetFieldValue,
		},
	)

	instance.SetException(variable.NewString("GetFieldObjectErrorException"), GetFieldObjectErrorException)
	instance.SetException(variable.NewString("GetFieldKeyErrorException"), GetFieldKeyErrorException)
	instance.SetException(variable.NewString("GetFieldKeyNotExistException"), GetFieldKeyNotExistException)

	instance.SetConst(variable.NewString("GetFieldContent"), GetFieldContent)
	instance.SetConst(variable.NewString("GetFieldKey"), GetFieldKey)
	instance.SetConst(variable.NewString("GetFieldValue"), GetFieldValue)

	instance.SetFunction(variable.NewString("GetField"), GetField)
}
