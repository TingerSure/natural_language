package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

const (
	VariableSystemFunctionType = "system_function"
	FunctionSystemFunctionType = "system"
)

type SystemFunction struct {
	funcs func(concept.Param, concept.Object) (concept.Param, concept.Exception)
}

func (f *SystemFunction) ToString(prefix string) string {
	return fmt.Sprintf("system_function")
}

func (f *SystemFunction) Exec(params concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	return f.funcs(params, object)
}

func (s *SystemFunction) Type() string {
	return VariableSystemFunctionType
}
func (s *SystemFunction) FunctionType() string {
	return FunctionSystemFunctionType
}

func NewSystemFunction(funcs func(concept.Param, concept.Object) (concept.Param, concept.Exception)) *SystemFunction {
	return &SystemFunction{
		funcs: funcs,
	}
}
