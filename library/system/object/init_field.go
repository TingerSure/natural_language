package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	InitFieldContent      = variable.NewString("object")
	InitFieldKey          = variable.NewString("key")
	InitFieldDefaultValue = variable.NewString("default_value")
)

var (
	InitFieldObjectErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("InitFieldObjectErrorException"))
	InitFieldKeyErrorException    = interrupt.NewException(variable.NewString("type error"), variable.NewString("InitFieldKeyErrorException"))
	InitFieldKeyExistException    = interrupt.NewException(variable.NewString("type error"), variable.NewString("InitFieldKeyExistException"))
)

var (
	InitField *variable.SystemFunction = nil
)

func init() {
	InitField = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(InitFieldContent))
			if !ok {
				return nil, InitFieldObjectErrorException.Copy().AddStack(InitField)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(InitFieldKey))
			if !ok {
				return nil, InitFieldKeyErrorException.Copy().AddStack(InitField)
			}
			if object.HasField(key) {
				return nil, InitFieldKeyExistException.Copy().AddStack(InitField)
			}

			defaultValue := input.Get(InitFieldDefaultValue)
			if nl_interface.IsNil(defaultValue) {
				defaultValue = variable.NewNull()
			}

			suspend := object.InitField(key, defaultValue)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(InitField)
			}

			return variable.NewParam().Set(InitFieldContent, object), nil
		},
		[]concept.String{
			InitFieldContent,
			InitFieldKey,
			InitFieldDefaultValue,
		},
		[]concept.String{
			InitFieldContent,
		},
	)
}
