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

	instance.Interrupt = NewInterruptCreator(&InterruptCreatorParam{})

	instance.Closure = closure.NewClosureCreator(&closure.ClosureCreatorParam{
		EmptyCreator: func() concept.Null {
			return instance.Variable.Null.New()
		},
		ExceptionCreator: func(name string, message string) concept.Exception {
			return instance.Variable.Exception.NewOriginal(name, message)
		},
	})

	instance.CodeBlock = code_block.NewCodeBlockCreator(&code_block.CodeBlockCreatorParam{
		ClosureCreator: func(parent concept.Closure) concept.Closure {
			return instance.Closure.New(parent)
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
		ClosureCreator: func(parent concept.Closure) concept.Closure {
			return instance.Closure.New(parent)
		},
	})

	instance.Expression = NewExpressionCreator(&ExpressionCreatorParam{
		FunctionCreator: func(parent concept.Closure) *variable.Function {
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
		EndCreator: func() *interrupt.End {
			return instance.Interrupt.End.New()
		},
		ClosureCreator: func(parent concept.Closure) concept.Closure {
			return instance.Closure.New(parent)
		},
		ObjectCreator: func() concept.Object {
			return instance.Variable.Object.New()
		},
	})
	return instance
}
