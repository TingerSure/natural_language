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

	GetFieldObjectErrorExceptionTemplate = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("GetFieldObjectErrorException"))
	GetFieldKeyErrorExceptionTemplate    = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("GetFieldKeyErrorException"))
	GetFieldKeyNotExistExceptionTemplate = libs.Sandbox.Interrupt.Exception.New(libs.Sandbox.Variable.String.New("type error"), libs.Sandbox.Variable.String.New("GetFieldKeyNotExistException"))
)

func initGetField(instance *Object) {
	GetFieldContent := libs.Sandbox.Variable.String.New(GetFieldContentName)
	GetFieldKey := libs.Sandbox.Variable.String.New(GetFieldKeyName)
	GetFieldValue := libs.Sandbox.Variable.String.New(GetFieldValueName)

	GetFieldObjectErrorException := GetFieldObjectErrorExceptionTemplate.Copy()
	GetFieldKeyErrorException := GetFieldKeyErrorExceptionTemplate.Copy()
	GetFieldKeyNotExistException := GetFieldKeyNotExistExceptionTemplate.Copy()

	var GetField concept.Function
	GetField = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("GetField"),
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

			return libs.Sandbox.Variable.Param.New().Set(GetFieldValue, value), nil
		},
		[]concept.String{
			GetFieldContent,
			GetFieldKey,
		},
		[]concept.String{
			GetFieldValue,
		},
	)

	instance.SetException(libs.Sandbox.Variable.String.New("GetFieldObjectErrorException"), GetFieldObjectErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("GetFieldKeyErrorException"), GetFieldKeyErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("GetFieldKeyNotExistException"), GetFieldKeyNotExistException)

	instance.SetConst(libs.Sandbox.Variable.String.New("GetFieldContent"), GetFieldContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("GetFieldKey"), GetFieldKey)
	instance.SetConst(libs.Sandbox.Variable.String.New("GetFieldValue"), GetFieldValue)

	instance.SetFunction(libs.Sandbox.Variable.String.New("GetField"), GetField)
}
