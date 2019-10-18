package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

const (
	VariableSystemFunctionType = VariableFunctionType
	FunctionSystemFunctionType = "system"
)

type SystemFunction struct {
	funcs func(concept.Param) (concept.Param, concept.Exception)
}

func (f *SystemFunction) ToString(prefix string) string {
	return fmt.Sprintf("system_function")
}

func (f *SystemFunction) Exec(params concept.Param) (concept.Param, concept.Exception) {
	return f.funcs(params)
}

func (s *SystemFunction) Type() string {
	return VariableSystemFunctionType
}
func (s *SystemFunction) FunctionType() string {
	return FunctionSystemFunctionType
}

func NewSystemFunction(funcs func(concept.Param) (concept.Param, concept.Exception)) *SystemFunction {
	return &SystemFunction{
		funcs: funcs,
	}
}
