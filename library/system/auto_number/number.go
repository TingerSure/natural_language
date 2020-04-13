package auto_number

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	AutoNumberClassValue   = variable.NewString("value")
	AutoNumberClassName    = "system.auto_number"
	AutoNumberClass        = variable.NewClass(AutoNumberClassName)
	AutoNumberClassMapping = map[concept.String]concept.String{
		AutoNumberClassValue: AutoNumberClassValue,
	}
)

func init() {
	AutoNumberClass.SetField(AutoNumberClassValue, variable.NewNumber(0))
}

func NewAutoNumberObject(value *variable.Number) concept.Object {
	auto := variable.NewObject()
	auto.InitField(AutoNumberClassValue, value)
	auto.AddClass(AutoNumberClass, "", AutoNumberClassMapping)
	object, suspend := variable.NewMappingObject(auto, AutoNumberClassName, "")
	if !nl_interface.IsNil(suspend) {
		panic(suspend)
	}
	return object
}

type AutoNumber struct {
	tree.Page
	AutoNumberClass concept.Class
}

func NewAutoNumber(libs *tree.LibraryManager) *AutoNumber {
	instance := &AutoNumber{
		Page:            tree.NewPageAdaptor(),
		AutoNumberClass: AutoNumberClass,
	}
	instance.SetClass(variable.NewString("AutoNumberClass"), AutoNumberClass)
	instance.SetConst(variable.NewString("AutoNumberClassValue"), AutoNumberClassValue)
	instance.SetConst(variable.NewString("AutoNumberClassName"), variable.NewString(AutoNumberClassName))

	return instance
}
