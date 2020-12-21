package number

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	NumberValueName      = "value"
	NumberClassValueName = NumberValueName
	NumberClassName      = "system.number"

	NewNumberContentName = "content"
	NewNumberObjectName  = "object"
)

type Number struct {
	tree.Page
	NumberValue                 concept.String
	NumberClassValue            concept.String
	NumberClass                 concept.Class
	NewNumberContent            concept.String
	NewNumberObject             concept.String
	NewNumber                   concept.Function
	NewNumberNotNumberException concept.Exception
	New                         func(value *variable.Number) concept.Object
}

func NewNumber(libs *runtime.LibraryManager) *Number {
	instance := &Number{
		Page:             tree.NewPageAdaptor(libs.Sandbox),
		NumberValue:      libs.Sandbox.Variable.String.New(NumberValueName),
		NumberClassValue: libs.Sandbox.Variable.String.New(NumberClassValueName),
		NumberClass:      libs.Sandbox.Variable.Class.New(NumberClassName),
		NewNumberContent: libs.Sandbox.Variable.String.New(NewNumberContentName),
		NewNumberObject:  libs.Sandbox.Variable.String.New(NewNumberObjectName),
	}

	instance.NumberClass.SetField(instance.NumberClassValue, libs.Sandbox.Variable.Number.New(0))

	instance.New = func(value *variable.Number) concept.Object {
		auto := libs.Sandbox.Variable.Object.New()
		auto.InitField(instance.NumberValue, value)
		auto.AddClass(instance.NumberClass, "", map[concept.String]concept.String{
			instance.NumberClassValue: instance.NumberValue,
		})
		object, suspend := libs.Sandbox.Variable.MappingObject.New(auto, NumberClassName, "")
		if !nl_interface.IsNil(suspend) {
			panic(suspend)
		}
		return object
	}
	instance.NewNumberNotNumberException = libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "NewNumberNotNumberException")

	anticipateNumber := instance.New(libs.Sandbox.Variable.Number.New(0))

	instance.NewNumber = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("NewNumber"),
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			content := input.Get(instance.NewNumberContent)

			number, ok := variable.VariableFamilyInstance.IsNumber(content)
			if !ok {
				return nil, instance.NewNumberNotNumberException.Copy().AddStack(instance.NewNumber)
			}

			return libs.Sandbox.Variable.Param.New().Set(instance.NewNumberObject, instance.New(number)), nil
		},
		func(_ concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(instance.NewNumberObject, anticipateNumber)
		},
		[]concept.String{
			instance.NewNumberContent,
		},
		[]concept.String{
			instance.NewNumberObject,
		},
	)

	instance.SetClass(libs.Sandbox.Variable.String.New("NumberClass"), instance.NumberClass)
	instance.SetConst(libs.Sandbox.Variable.String.New("NumberClassValue"), instance.NumberClassValue)
	instance.SetConst(libs.Sandbox.Variable.String.New("NumberValue"), instance.NumberValue)
	instance.SetConst(libs.Sandbox.Variable.String.New("NewNumberContent"), instance.NewNumberContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("NewNumberObject"), instance.NewNumberObject)
	instance.SetFunction(libs.Sandbox.Variable.String.New("NewNumber"), instance.NewNumber)
	instance.SetException(libs.Sandbox.Variable.String.New("NewNumberNotNumberException"), instance.NewNumberNotNumberException)

	return instance
}
