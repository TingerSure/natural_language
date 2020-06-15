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
)

type AutoNumber struct {
	tree.Page
	AutoNumberValue      concept.String
	AutoNumberClassValue concept.String
	AutoNumberClass      concept.Class
	NewAutoNumberObject  func(*variable.Number) concept.Object
}

func NewAutoNumber(libs *runtime.LibraryManager) *AutoNumber {
	instance := &AutoNumber{
		Page:                 tree.NewPageAdaptor(libs.Sandbox),
		AutoNumberValue:      libs.Sandbox.Variable.String.New(AutoNumberValueName),
		AutoNumberClassValue: libs.Sandbox.Variable.String.New(AutoNumberClassValueName),
		AutoNumberClass:      variable.NewClass(AutoNumberClassName),
	}

	instance.AutoNumberClass.SetField(instance.AutoNumberClassValue, libs.Sandbox.Variable.Number.New(0))

	instance.NewAutoNumberObject = func(value *variable.Number) concept.Object {
		auto := variable.NewObject()
		auto.InitField(instance.AutoNumberValue, value)
		auto.AddClass(instance.AutoNumberClass, "", map[concept.String]concept.String{
			instance.AutoNumberClassValue: instance.AutoNumberValue,
		})
		object, suspend := variable.NewMappingObject(auto, AutoNumberClassName, "")
		if !nl_interface.IsNil(suspend) {
			panic(suspend)
		}
		return object
	}

	initAddition(instance)

	instance.SetClass(libs.Sandbox.Variable.String.New("AutoNumberClass"), instance.AutoNumberClass)
	instance.SetConst(libs.Sandbox.Variable.String.New("AutoNumberClassValue"), instance.AutoNumberClassValue)
	instance.SetConst(libs.Sandbox.Variable.String.New("AutoNumberValue"), instance.AutoNumberValue)

	return instance
}
