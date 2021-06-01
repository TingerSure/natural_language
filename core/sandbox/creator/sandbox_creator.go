package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/pool"
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type SandboxCreator struct {
	Variable   *VariableCreator
	Index      *IndexCreator
	Interrupt  *InterruptCreator
	Expression *ExpressionCreator
	Pool    *pool.PoolCreator
	CodeBlock  *code_block.CodeBlockCreator
}

func NewSandboxCreator() *SandboxCreator {
	instance := &SandboxCreator{}

	instance.Interrupt = NewInterruptCreator(&InterruptCreatorParam{})

	instance.Pool = pool.NewPoolCreator(&pool.PoolCreatorParam{
		EmptyCreator: func() concept.Null {
			return instance.Variable.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Variable.Exception.NewOriginal(name, message)
		},
	})

	instance.CodeBlock = code_block.NewCodeBlockCreator(&code_block.CodeBlockCreatorParam{
		PoolCreator: func(parent concept.Pool) concept.Pool {
			return instance.Pool.New(parent)
		},
	})

	instance.Index = NewIndexCreator(&IndexCreatorParam{
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Variable.Exception.NewOriginal(name, message)
		},
		NullCreator: func() concept.Null {
			return instance.Variable.Null.New()
		},
		ParamCreator: func() concept.Param {
			return instance.Variable.Param.New()
		},
		StringCreator: func(value string) concept.String {
			return instance.Variable.String.New(value)
		},
	})

	instance.Variable = NewVariableCreator(&VariableCreatorParam{
		CodeBlockCreator: func() *code_block.CodeBlock {
			return instance.CodeBlock.New()
		},
		PoolCreator: func(parent concept.Pool) concept.Pool {
			return instance.Pool.New(parent)
		},
	})

	instance.Expression = NewExpressionCreator(&ExpressionCreatorParam{
		MappingObjectCreator: func(object concept.Variable, class concept.Class) *variable.MappingObject {
			return instance.Variable.MappingObject.New(object, class)
		},
		ClassCreator: func() concept.Class {
			return instance.Variable.Class.New()
		},
		DefineFunctionCreator: func(paramNames []concept.String, returnNames []concept.String) *variable.DefineFunction {
			return instance.Variable.DefineFunction.New(paramNames, returnNames)
		},
		FunctionCreator: func(parent concept.Pool) *variable.Function {
			return instance.Variable.Function.New(parent)
		},
		CodeBlockCreator: func() *code_block.CodeBlock {
			return instance.CodeBlock.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Variable.Exception.NewOriginal(name, message)
		},
		NullCreator: func() concept.Null {
			return instance.Variable.Null.New()
		},
		ParamCreator: func() concept.Param {
			return instance.Variable.Param.New()
		},
		ConstIndexCreator: func(value concept.Variable) *index.ConstIndex {
			return instance.Index.ConstIndex.New(value)
		},
		StringCreator: func(value string) concept.String {
			return instance.Variable.String.New(value)
		},
		BoolCreator: func(value bool) concept.Bool {
			return instance.Variable.Bool.New(value)
		},
		NumberCreator: func(value float64) concept.Number {
			return instance.Variable.Number.New(value)
		},
		ReturnCreator: func() *interrupt.Return {
			return instance.Interrupt.Return.New()
		},
		PoolCreator: func(parent concept.Pool) concept.Pool {
			return instance.Pool.New(parent)
		},
		ObjectCreator: func() concept.Object {
			return instance.Variable.Object.New()
		},
		ContinueCreator: func(tag concept.String) *interrupt.Continue {
			return instance.Interrupt.Continue.New(tag)
		},
		BreakCreator: func(tag concept.String) *interrupt.Break {
			return instance.Interrupt.Break.New(tag)
		},
		ArrayCreator: func() *variable.Array {
			return instance.Variable.Array.New()
		},
	})
	return instance
}
