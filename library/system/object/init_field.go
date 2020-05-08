package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	InitFieldContentName      = "object"
	InitFieldKeyName          = "key"
	InitFieldDefaultValueName = "default_value"

	InitFieldObjectErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("InitFieldObjectErrorException"))
	InitFieldKeyErrorExceptionTemplate    = interrupt.NewException(variable.NewString("type error"), variable.NewString("InitFieldKeyErrorException"))
	InitFieldKeyExistExceptionTemplate    = interrupt.NewException(variable.NewString("type error"), variable.NewString("InitFieldKeyExistException"))
)

func initInitField(instance *Object) {
	InitFieldContent := variable.NewString(InitFieldContentName)
	InitFieldKey := variable.NewString(InitFieldKeyName)
	InitFieldDefaultValue := variable.NewString(InitFieldDefaultValueName)

	InitFieldObjectErrorException := InitFieldObjectErrorExceptionTemplate.Copy()
	InitFieldKeyErrorException := InitFieldKeyErrorExceptionTemplate.Copy()
	InitFieldKeyExistException := InitFieldKeyExistExceptionTemplate.Copy()

	var InitField concept.Function
	InitField = variable.NewSystemFunction(
		variable.NewString("InitField"),
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

	instance.SetException(variable.NewString("InitFieldObjectErrorException"), InitFieldObjectErrorException)
	instance.SetException(variable.NewString("InitFieldKeyErrorException"), InitFieldKeyErrorException)
	instance.SetException(variable.NewString("InitFieldKeyExistException"), InitFieldKeyExistException)

	instance.SetConst(variable.NewString("InitFieldContent"), InitFieldContent)
	instance.SetConst(variable.NewString("InitFieldKey"), InitFieldKey)
	instance.SetConst(variable.NewString("InitFieldDefaultValue"), InitFieldDefaultValue)

	instance.SetFunction(variable.NewString("InitField"), InitField)

}
