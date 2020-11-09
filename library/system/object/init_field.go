package object

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	InitFieldContentName      = "object"
	InitFieldKeyName          = "key"
	InitFieldDefaultValueName = "default_value"
)

func initInitField(libs *runtime.LibraryManager, instance *Object) {
	InitFieldContent := libs.Sandbox.Variable.String.New(InitFieldContentName)
	InitFieldKey := libs.Sandbox.Variable.String.New(InitFieldKeyName)
	InitFieldDefaultValue := libs.Sandbox.Variable.String.New(InitFieldDefaultValueName)

	InitFieldObjectErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "InitFieldObjectErrorException")
	InitFieldKeyErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "InitFieldKeyErrorException")
	InitFieldKeyExistException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "InitFieldKeyExistException")

	var InitField concept.Function
	InitField = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("InitField"),
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
				defaultValue = libs.Sandbox.Variable.Null.New()
			}

			suspend := object.InitField(key, defaultValue)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend.AddStack(InitField)
			}

			return libs.Sandbox.Variable.Param.New().Set(InitFieldContent, object), nil
		},
		func(input concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(InitFieldContent, input.Get(InitFieldContent))
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

	instance.SetException(libs.Sandbox.Variable.String.New("InitFieldObjectErrorException"), InitFieldObjectErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("InitFieldKeyErrorException"), InitFieldKeyErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("InitFieldKeyExistException"), InitFieldKeyExistException)

	instance.SetConst(libs.Sandbox.Variable.String.New("InitFieldContent"), InitFieldContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("InitFieldKey"), InitFieldKey)
	instance.SetConst(libs.Sandbox.Variable.String.New("InitFieldDefaultValue"), InitFieldDefaultValue)

	instance.SetFunction(libs.Sandbox.Variable.String.New("InitField"), InitField)

}
