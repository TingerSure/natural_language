package auto_number

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

var (
	AutoNumberClassValue   = "value"
	AutoNumberClassName    = "system.auto.number"
	AutoNumberClass        = variable.NewClass(AutoNumberClassName)
	AutoNumberClassMapping = map[string]string{
		AutoNumberClassValue: AutoNumberClassValue,
	}
)

func init() {
	AutoNumberClass.SetField(AutoNumberClassValue, variable.NewNumber(0))
}

func NewAutoNumber(value *variable.Number) concept.Object {
	auto := variable.NewObject()
	auto.InitField(AutoNumberClassValue, value)
	auto.AddClass(AutoNumberClass, "", AutoNumberClassMapping)
	object, suspend := variable.NewMappingObject(auto, AutoNumberClassName, "")
	if !nl_interface.IsNil(suspend) {
		panic(suspend)
	}
	return object
}
