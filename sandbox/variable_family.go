package sandbox

import (
	"unsafe"
)

type VariableFamily struct {
}

func (v *VariableFamily) IsNumber(value Variable) (*Number, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableNumberType {
		return (*Number)(unsafe.Pointer(value)), true
	}
	return nil, false
}

func newVariableFamily() *VariableFamily {
	return &VariableFamily{}
}

var (
	VariableFamilyInstance *VariableFamily = newVariableFamily()
)
