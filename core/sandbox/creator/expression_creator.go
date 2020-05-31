package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
)

type InterruptCreator struct {
	Call *interrupt.CallCreator
}

type InterruptCreatorParam struct {
	CodeBlockCreator func() *code_block.CodeBlock
	ExceptionCreator func(string) concept.Exception
}

func NewInterruptCreator(param *InterruptCreatorParam) *InterruptCreator {
	instance := &InterruptCreator{}
	instance.Call = interrupt.NewCallCreator(&interrupt.CallCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})
	return instance
}
