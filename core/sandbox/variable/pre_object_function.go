package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariablePreObjectFunctionType = "pre_object_function"
	FunctionPreObjectFunctionType = "pre_object"
)

type PreObjectFunction struct {
	function concept.Function
	object   concept.Object
}

func (s *PreObjectFunction) ParamNames() []concept.String {
	return s.function.ParamNames()
}

func (s *PreObjectFunction) ReturnNames() []concept.String {
	return s.function.ReturnNames()
}

func (f *PreObjectFunction) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", f.object.ToString(prefix), f.function.ToString(prefix))
}

func (f *PreObjectFunction) Exec(params concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	if nl_interface.IsNil(object) {
		object = f.object
	}
	return f.function.Exec(params, object)
}

func (s *PreObjectFunction) Type() string {
	return VariablePreObjectFunctionType
}

func (s *PreObjectFunction) FunctionType() string {
	return FunctionPreObjectFunctionType
}

func NewPreObjectFunction(
	function concept.Function,
	object concept.Object,
) *PreObjectFunction {
	return &PreObjectFunction{
		function: function,
		object:   object,
	}
}
