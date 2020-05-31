package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type SandboxCreator struct {
	Variable   *VariableCreator
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
	})
	return instance
}
