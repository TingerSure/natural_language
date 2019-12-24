package object

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

const (
	InitFieldContent      = "object"
	InitFieldKey          = "key"
	InitFieldDefaultValue = "default_value"
)

var (
	InitFieldObjectErrorException = interrupt.NewException("type error", "TODO")
	InitFieldKeyErrorException    = interrupt.NewException("type error", "TODO")
	InitFieldKeyExistException    = interrupt.NewException("type error", "TODO")
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
			if object.HasField(key.Value()) {
				return nil, InitFieldKeyExistException.Copy().AddStack(InitField)
			}

			defaultValue := input.Get(InitFieldDefaultValue)
			if nl_interface.IsNil(defaultValue) {
				defaultValue = variable.NewNull()
			}

			suspend := object.InitField(key.Value(), defaultValue)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(InitField)
			}

			return variable.NewParamWithInit(map[string]concept.Variable{
				InitFieldContent: object,
			}), nil
		},
		[]string{
			InitFieldContent,
			InitFieldKey,
			InitFieldDefaultValue,
		},
		[]string{
			InitFieldContent,
		},
	)
}
