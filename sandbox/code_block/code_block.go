package code_block

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/closure"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type CodeBlock struct {
	flow []concept.Expression
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
func (c *CodeBlock) AddStep(step concept.Expression) {
	c.flow = append(c.flow, step)
}

func (f *CodeBlock) Exec(parent concept.Closure, returnBubble bool, init func(concept.Closure) concept.Interrupt) (concept.Closure, concept.Interrupt) {

	if parent == nil && returnBubble {
		returnBubble = false
	}

	space := closure.NewClosure(parent)
	defer func() {
		if returnBubble {
			parent.MergeReturn(space)
		}
	}()

	if init != nil {
		suspend := init(space)
		if !nl_interface.IsNil(suspend) {
			return space, suspend
		}
	}
	for _, step := range f.flow {
		suspend := step.Exec(space)
		if !nl_interface.IsNil(suspend) {
			return space, suspend
		}
	}
	return space, nil
}

func NewCodeBlock() *CodeBlock {
	return &CodeBlock{}
}
