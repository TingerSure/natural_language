package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/closure"
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
	Closure    *closure.ClosureCreator
	CodeBlock  *code_block.CodeBlockCreator
}

func NewSandboxCreator() *SandboxCreator {
	instance := &SandboxCreator{}

	instance.Interrupt = NewInterruptCreator(&InterruptCreatorParam{
		StringCreator: func(value string) concept.String {
			return instance.Variable.String.New(value)
		},
	})

	instance.Closure = closure.NewClosureCreator(&closure.ClosureCreatorParam{
		EmptyCreator: func() concept.Null {
			return instance.Variable.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Interrupt.Exception.NewOriginal(name, message)
		},
	})

	instance.CodeBlock = code_block.NewCodeBlockCreator(&code_block.CodeBlockCreatorParam{
		ClosureCreator: func(parent concept.Closure) concept.Closure {
			return instance.Closure.New(parent)
		},
	})

	instance.Index = NewIndexCreator(&IndexCreatorParam{
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Interrupt.Exception.NewOriginal(name, message)
		},
		NullCreator: func() concept.Null {
			return instance.Variable.Null.New()
		},
		StringCreator: func(value string) concept.String {
			return instance.Variable.String.New(value)
		},
		PreObjectFunctionCreator: func(function concept.Function, object concept.Object) *variable.PreObjectFunction {
			return instance.Variable.PreObjectFunction.New(function, object)
		},
	})

	instance.Variable = NewVariableCreator(&VariableCreatorParam{
		CodeBlockCreator: func() *code_block.CodeBlock {
			return instance.CodeBlock.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Interrupt.Exception.NewOriginal(name, message)
		},
	})

	instance.Expression = NewExpressionCreator(&ExpressionCreatorParam{
		CodeBlockCreator: func() *code_block.CodeBlock {
			return instance.CodeBlock.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Interrupt.Exception.NewOriginal(name, message)
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
		BoolCreator: func(value bool) *variable.Bool {
			return instance.Variable.Bool.New(value)
		},
		EndCreator: func() *interrupt.End {
			return instance.Interrupt.End.New()
		},
		ClosureCreator: func(parent concept.Closure) concept.Closure {
			return instance.Closure.New(parent)
		},
	})
	return instance
}
