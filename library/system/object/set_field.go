package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	SetFieldContent = variable.NewString("object")
	SetFieldKey     = variable.NewString("key")
	SetFieldValue   = variable.NewString("value")
)

var (
	SetFieldObjectErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetFieldObjectErrorException"))
	SetFieldKeyErrorException    = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetFieldKeyErrorException"))
	SetFieldKeyNotExistException = interrupt.NewException(variable.NewString("type error"), variable.NewString("SetFieldKeyNotExistException"))
)

var (
	SetField *variable.SystemFunction = nil
)

func init() {
	SetField = variable.NewSystemFunction(
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
}
