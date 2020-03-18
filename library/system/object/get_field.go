package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/library/system/auto_number"
)

var (
	GetFieldContent = variable.NewString("object")
	GetFieldKey     = variable.NewString("key")
	GetFieldValue   = variable.NewString("value")
)

var (
	GetFieldObjectErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetFieldObjectErrorException"))
	GetFieldKeyErrorException    = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetFieldKeyErrorException"))
	GetFieldKeyNotExistException = interrupt.NewException(variable.NewString("type error"), variable.NewString("GetFieldKeyNotExistException"))
)

var (
	GetField *variable.SystemFunction = nil
)

func init() {
	GetField = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			content := input.Get(GetFieldContent)
			var object concept.Object
			if objectHome, ok := variable.VariableFamilyInstance.IsObjectHome(content); ok {
				object = objectHome
			} else if number, ok := variable.VariableFamilyInstance.IsNumber(content); ok {
				object = auto_number.NewAutoNumberObject(number)
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
}
