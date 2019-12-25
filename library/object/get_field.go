package object

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

const (
	GetFieldContent = "object"
	GetFieldKey     = "key"
	GetFieldValue   = "value"
)

var (
	GetFieldObjectErrorException = interrupt.NewException("type error", "GetFieldObjectErrorException")
	GetFieldKeyErrorException    = interrupt.NewException("type error", "GetFieldKeyErrorException")
	GetFieldKeyNotExistException = interrupt.NewException("type error", "GetFieldKeyNotExistException")
)

var (
	GetField *variable.SystemFunction = nil
)

func init() {
	GetField = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(GetFieldContent))
			if !ok {
				return nil, GetFieldObjectErrorException.Copy().AddStack(GetField)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(GetFieldKey))
			if !ok {
				return nil, GetFieldKeyErrorException.Copy().AddStack(GetField)
			}
			if !object.HasField(key.Value()) {
				return nil, GetFieldKeyNotExistException.Copy().AddStack(GetField)
			}

			value, suspend := object.GetField(key.Value())
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(GetField)
			}

			return variable.NewParamWithInit(map[string]concept.Variable{
				GetFieldValue: value,
			}), nil
		},
		[]string{
			GetFieldContent,
			GetFieldKey,
		},
		[]string{
			GetFieldValue,
		},
	)
}
