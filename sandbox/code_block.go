package sandbox

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
)

type CodeBlock struct {
	flow []Expression
}

func (c *CodeBlock) Size() int {
	return len(c.flow)
}

func (c *CodeBlock) ToString(prefix string) string {
	flowPrefix := fmt.Sprintf("%v\t", prefix)
	flowToStrings := ""
	for _, flow := range c.flow {
		flowToStrings = fmt.Sprintf("%v\n%v", flowToStrings, flow.ToString(flowPrefix))
	}
	return fmt.Sprintf("{%v\n%v}", flowToStrings, prefix)
}
func (c *CodeBlock) AddStep(step Expression) {
	c.flow = append(c.flow, step)
}

func (f *CodeBlock) Exec(parent *Closure, returnBubble bool, init func(*Closure) Interrupt) (*Closure, Interrupt) {

	if parent == nil && returnBubble {
		returnBubble = false
	}

	space := NewClosure(parent)
	defer func() {
		if returnBubble {
			parent.MergeReturn(space)
		}
	}()

	if init != nil {
		suspend := init(space)
		if suspend != nil {
			return space, suspend
		}
	}
	for _, step := range f.flow {
		suspend := step.Exec(space)
		fmt.Printf("%+v\n", !nl_interface.IsNil(suspend))
		if !nl_interface.IsNil(suspend) {
			fmt.Printf("%v\n", suspend.InterruptType())
			return space, suspend
		}
	}
	return space, nil
}

func NewCodeBlock() *CodeBlock {
	return &CodeBlock{}
}
