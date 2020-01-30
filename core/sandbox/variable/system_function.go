package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariableSystemFunctionType = "system_function"
	FunctionSystemFunctionType = "system"
)

type SystemFunction struct {
	paramNames  []string
	returnNames []string
	funcs       func(concept.Param, concept.Object) (concept.Param, concept.Exception)
}

func (s *SystemFunction) ParamNames() []string {
	return s.paramNames
}

func (s *SystemFunction) ReturnNames() []string {
	return s.returnNames
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

func NewSystemFunction(
	funcs func(concept.Param, concept.Object) (concept.Param, concept.Exception),
	paramNames []string,
	returnNames []string,
) *SystemFunction {
	return &SystemFunction{
		funcs: funcs,
	}
}
