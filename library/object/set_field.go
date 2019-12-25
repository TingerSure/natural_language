package object

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

const (
	SetFieldContent = "object"
	SetFieldKey     = "key"
	SetFieldValue   = "value"
)

var (
	SetFieldObjectErrorException = interrupt.NewException("type error", "SetFieldObjectErrorException")
	SetFieldKeyErrorException    = interrupt.NewException("type error", "SetFieldKeyErrorException")
	SetFieldKeyNotExistException = interrupt.NewException("type error", "SetFieldKeyNotExistException")
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
			if !object.HasField(key.Value()) {
				return nil, SetFieldKeyNotExistException.Copy().AddStack(SetField)
			}

			value := input.Get(SetFieldValue)
			if nl_interface.IsNil(value) {
				value = variable.NewNull()
			}

			suspend := object.SetField(key.Value(), value)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(SetField)
			}

			return variable.NewParamWithInit(map[string]concept.Variable{
				SetFieldContent: object,
			}), nil
		},
		[]string{
			SetFieldContent,
			SetFieldKey,
			SetFieldValue,
		},
		[]string{
			SetFieldContent,
		},
	)
}
