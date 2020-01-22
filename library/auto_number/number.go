package auto_number

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

var (
	AutoNumberClassValue = "value"
	AutoNumberClass      = variable.NewClass("system.auto.number")
)

func init() {
	AutoNumberClass.SetField(AutoNumberClassValue, variable.NewNumber(0))
}

func NewAutoNumber(value *variable.Number) concept.Object {
	auto := variable.NewObject()
	auto.SetField(AutoNumberClassValue, value)
	auto.AddClass(AutoNumberClass, "", map[string]string{
		AutoNumberClassValue: AutoNumberClassValue,
	})
	return auto
}
