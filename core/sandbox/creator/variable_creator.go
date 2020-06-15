package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type VariableCreator struct {
	String            *variable.StringCreator
	Number            *variable.NumberCreator
	Null              *variable.NullCreator
	Bool              *variable.BoolCreator
	Class             *variable.ClassCreator
	Function          *variable.FunctionCreator
	Object            *variable.ObjectCreator
	MappingObject     *variable.MappingObjectCreator
	Param             *variable.ParamCreator
	SystemFunction    *variable.SystemFunctionCreator
	PreObjectFunction *variable.PreObjectFunctionCreator
}

type VariableCreatorParam struct {
	CodeBlockCreator func() *code_block.CodeBlock
	ExceptionCreator func(string, string) concept.Exception
}

func NewVariableCreator(param *VariableCreatorParam) *VariableCreator {
	instance := &VariableCreator{}
	instance.PreObjectFunction = variable.NewPreObjectFunctionCreator(&variable.PreObjectFunctionCreatorParam{})
	instance.SystemFunction = variable.NewSystemFunctionCreator(&variable.SystemFunctionCreatorParam{})
	instance.Bool = variable.NewBoolCreator()
	instance.Null = variable.NewNullCreator()
	instance.String = variable.NewStringCreator()
	instance.Number = variable.NewNumberCreator()
	instance.Class = variable.NewClassCreator(&variable.ClassCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
	})

	instance.Param = variable.NewParamCreator(&variable.ParamCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
	})

	instance.Function = variable.NewFunctionCreator(&variable.FunctionCreatorParam{
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
		ParamCreator: func() concept.Param {
			return instance.Param.New()
		},
		ExceptionCreator: param.ExceptionCreator,
		CodeBlockCreator: param.CodeBlockCreator,
	})

	instance.Object = variable.NewObjectCreator(&variable.ObjectCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: param.ExceptionCreator,
	})

	instance.MappingObject = variable.NewMappingObjectCreator(&variable.MappingObjectCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})

	return instance
}
