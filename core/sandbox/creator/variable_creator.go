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
	Exception      *variable.ExceptionCreator
	Page           *variable.PageCreator
	DefineFunction *variable.DefineFunctionCreator
	DelayFunction  *variable.DelayFunctionCreator
	DelayString    *variable.DelayStringCreator
}

type VariableCreatorParam struct {
	CodeBlockCreator func() *code_block.CodeBlock
	ClosureCreator   func(concept.Closure) concept.Closure
}

func NewVariableCreator(param *VariableCreatorParam) *VariableCreator {
	instance := &VariableCreator{}
	instance.DelayString = variable.NewDelayStringCreator(&variable.DelayStringCreatorParam{
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
	})
	instance.DelayFunction = variable.NewDelayFunctionCreator(&variable.DelayFunctionCreatorParam{})
	instance.DefineFunction = variable.NewDefineFunctionCreator(&variable.DefineFunctionCreatorParam{
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
		ParamCreator: func() concept.Param {
			return instance.Param.New()
		},
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		DelayStringCreator: func(original string) concept.String {
			return instance.DelayString.New(original)
		},
		DelayFunctionCreator: func(create func() concept.Function) concept.Function {
			return instance.DelayFunction.New(create)
		},
		SystemFunctionCreator: func(
			funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
			anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
			paramNames []concept.String,
			returnNames []concept.String,
		) concept.Function {
			return instance.SystemFunction.New(funcs, anticipateFuncs, paramNames, returnNames)
		},
		ArrayCreator: func() concept.Array {
			return instance.Array.New()
		},
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
	})
	instance.Exception = variable.NewExceptionCreator(&variable.ExceptionCreatorParam{
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
	})
	instance.Page = variable.NewPageCreator(&variable.PageCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
		ClosureCreator: param.ClosureCreator,
	})
	instance.Null = variable.NewNullCreator(&variable.NullCreatorParam{
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
	})
	instance.SystemFunction = variable.NewSystemFunctionCreator(&variable.SystemFunctionCreatorParam{
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ParamCreator: func() concept.Param {
			return instance.Param.New()
		},
		DelayStringCreator: func(original string) concept.String {
			return instance.DelayString.New(original)
		},
		DelayFunctionCreator: func(create func() concept.Function) concept.Function {
			return instance.DelayFunction.New(create)
		},
		SystemFunctionCreator: func(
			funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
			anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
			paramNames []concept.String,
			returnNames []concept.String,
		) concept.Function {
			return instance.SystemFunction.New(funcs, anticipateFuncs, paramNames, returnNames)
		},
		ArrayCreator: func() concept.Array {
			return instance.Array.New()
		},
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
	})
	instance.Bool = variable.NewBoolCreator(&variable.BoolCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
	})
	instance.String = variable.NewStringCreator(&variable.StringCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
	})
	instance.Number = variable.NewNumberCreator(&variable.NumberCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
	})
	instance.Class = variable.NewClassCreator(&variable.ClassCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
	})

	instance.Param = variable.NewParamCreator(&variable.ParamCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
	})

	instance.Function = variable.NewFunctionCreator(&variable.FunctionCreatorParam{
		DelayStringCreator: func(original string) concept.String {
			return instance.DelayString.New(original)
		},
		DelayFunctionCreator: func(create func() concept.Function) concept.Function {
			return instance.DelayFunction.New(create)
		},
		SystemFunctionCreator: func(
			funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
			anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
			paramNames []concept.String,
			returnNames []concept.String,
		) concept.Function {
			return instance.SystemFunction.New(funcs, anticipateFuncs, paramNames, returnNames)
		},
		ArrayCreator: func() concept.Array {
			return instance.Array.New()
		},
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
		ParamCreator: func() concept.Param {
			return instance.Param.New()
		},
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
		CodeBlockCreator: param.CodeBlockCreator,
	})

	instance.Object = variable.NewObjectCreator(&variable.ObjectCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
	})

	instance.MappingObject = variable.NewMappingObjectCreator(&variable.MappingObjectCreatorParam{
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
	})

	instance.Array = variable.NewArrayCreator(&variable.ArrayCreatorParam{
		NullCreator: func() concept.Null {
			return instance.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Exception.NewOriginal(name, message)
		},
		ParamCreator: func() concept.Param {
			return instance.Param.New()
		},
		StringCreator: func(value string) concept.String {
			return instance.String.New(value)
		},
		DelayStringCreator: func(original string) concept.String {
			return instance.DelayString.New(original)
		},
		NumberCreator: func(value float64) concept.Number {
			return instance.Number.New(value)
		},
		DelayFunctionCreator: func(create func() concept.Function) concept.Function {
			return instance.DelayFunction.New(create)
		},
		SystemFunctionCreator: func(
			funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
			anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
			paramNames []concept.String,
			returnNames []concept.String,
		) concept.Function {
			return instance.SystemFunction.New(funcs, anticipateFuncs, paramNames, returnNames)
		},
	})
	return instance
}
