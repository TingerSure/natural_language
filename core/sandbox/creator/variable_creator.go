package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type VariableCreator struct {
	String         *variable.StringCreator
	Number         *variable.NumberCreator
	Null           *variable.NullCreator
	Bool           *variable.BoolCreator
	Class          *variable.ClassCreator
	Function       *variable.FunctionCreator
	Object         *variable.ObjectCreator
	MappingObject  *variable.MappingObjectCreator
	Param          *variable.ParamCreator
	SystemFunction *variable.SystemFunctionCreator
	Array          *variable.ArrayCreator
}

type VariableCreatorParam struct {
	CodeBlockCreator func() *code_block.CodeBlock
	ExceptionCreator func(string, string) concept.Exception
}

func NewVariableCreator(param *VariableCreatorParam) *VariableCreator {
	instance := &VariableCreator{}
	instance.Null = variable.NewNullCreator(&variable.NullCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})
	instance.SystemFunction = variable.NewSystemFunctionCreator(&variable.SystemFunctionCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
	})
	instance.Bool = variable.NewBoolCreator(&variable.BoolCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: param.ExceptionCreator,
	})
	instance.String = variable.NewStringCreator(&variable.StringCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: param.ExceptionCreator,
	})
	instance.Number = variable.NewNumberCreator(&variable.NumberCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: param.ExceptionCreator,
	})
	instance.Class = variable.NewClassCreator(&variable.ClassCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: param.ExceptionCreator,
	})

	instance.Param = variable.NewParamCreator(&variable.ParamCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: param.ExceptionCreator,
	})

	instance.Function = variable.NewFunctionCreator(&variable.FunctionCreatorParam{
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
		ParamCreator: func() concept.Param {
			return instance.Param.New()
		},
		NullCreator: func() concept.Null {
			return instance.Null.New()
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

	instance.Array = variable.NewArrayCreator(&variable.ArrayCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: param.ExceptionCreator,
	})

	return instance
}
