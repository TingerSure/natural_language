package auto_number

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	AutoNumberValueName      = "value"
	AutoNumberClassValueName = AutoNumberValueName
	AutoNumberClassName      = "system.auto_number"

	NewAutoNumberContentName = "content"
	NewAutoNumberObjectName  = "object"
)

type AutoNumber struct {
	tree.Page
	AutoNumberValue                 concept.String
	AutoNumberClassValue            concept.String
	AutoNumberClass                 concept.Class
	NewAutoNumberContent            concept.String
	NewAutoNumberObject             concept.String
	NewAutoNumber                   concept.Function
	NewAutoNumberNotNumberException concept.Exception
	New                             func(value *variable.Number) concept.Object
}

func NewAutoNumber(libs *runtime.LibraryManager) *AutoNumber {
	instance := &AutoNumber{
		Page:                 tree.NewPageAdaptor(libs.Sandbox),
		AutoNumberValue:      libs.Sandbox.Variable.String.New(AutoNumberValueName),
		AutoNumberClassValue: libs.Sandbox.Variable.String.New(AutoNumberClassValueName),
		AutoNumberClass:      libs.Sandbox.Variable.Class.New(AutoNumberClassName),
		NewAutoNumberContent: libs.Sandbox.Variable.String.New(NewAutoNumberContentName),
		NewAutoNumberObject:  libs.Sandbox.Variable.String.New(NewAutoNumberObjectName),
	}

	instance.AutoNumberClass.SetField(instance.AutoNumberClassValue, libs.Sandbox.Variable.Number.New(0))

	instance.New = func(value *variable.Number) concept.Object {
		auto := libs.Sandbox.Variable.Object.New()
		auto.InitField(instance.AutoNumberValue, value)
		auto.AddClass(instance.AutoNumberClass, "", map[concept.String]concept.String{
			instance.AutoNumberClassValue: instance.AutoNumberValue,
		})
		object, suspend := libs.Sandbox.Variable.MappingObject.New(auto, AutoNumberClassName, "")
		if !nl_interface.IsNil(suspend) {
			panic(suspend)
		}
		return object
	}
	instance.NewAutoNumberNotNumberException = libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "NewAutoNumberNotNumberException")

	anticipateNumber := instance.New(libs.Sandbox.Variable.Number.New(0))

	instance.NewAutoNumber = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("NewAutoNumber"),
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			content := input.Get(instance.NewAutoNumberContent)

			number, ok := variable.VariableFamilyInstance.IsNumber(content)
			if !ok {
				return nil, instance.NewAutoNumberNotNumberException.Copy().AddStack(instance.NewAutoNumber)
			}

			return libs.Sandbox.Variable.Param.New().Set(instance.NewAutoNumberObject, instance.New(number)), nil
		},
		func(_ concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(instance.NewAutoNumberObject, anticipateNumber)
		},
		[]concept.String{
			instance.NewAutoNumberContent,
		},
		[]concept.String{
			instance.NewAutoNumberObject,
		},
	)

	initAddition(libs, instance)

	instance.SetClass(libs.Sandbox.Variable.String.New("AutoNumberClass"), instance.AutoNumberClass)
	instance.SetConst(libs.Sandbox.Variable.String.New("AutoNumberClassValue"), instance.AutoNumberClassValue)
	instance.SetConst(libs.Sandbox.Variable.String.New("AutoNumberValue"), instance.AutoNumberValue)
	instance.SetConst(libs.Sandbox.Variable.String.New("NewAutoNumberContent"), instance.NewAutoNumberContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("NewAutoNumberObject"), instance.NewAutoNumberObject)
	instance.SetFunction(libs.Sandbox.Variable.String.New("NewAutoNumber"), instance.NewAutoNumber)
	instance.SetException(libs.Sandbox.Variable.String.New("NewAutoNumberNotNumberException"), instance.NewAutoNumberNotNumberException)

	return instance
}
