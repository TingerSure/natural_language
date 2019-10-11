package code_block

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/closure"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"strings"
)

type CodeBlock struct {
	flow []concept.Expression
}

func (c *CodeBlock) Size() int {
	return len(c.flow)
}

const (
	CodeBlockFlowNoneSize     = 0
	CodeBlockFlowOnlyOneSize  = 1
	CodeBlockFlowOnlyOneIndex = 0
)

func (c *CodeBlock) ToStringSimplify(prefix string) string {
	if c.Size() == CodeBlockFlowNoneSize {
		return ""
	}
	if c.Size() == CodeBlockFlowOnlyOneSize {
		return c.flow[CodeBlockFlowOnlyOneIndex].ToString(prefix)
	}
	return c.ToString(prefix)
}

func (c *CodeBlock) ToString(prefix string) string {
	flowPrefix := fmt.Sprintf("%v\t", prefix)
	flowToStrings := make([]string, 0, c.Size())
	for _, flow := range c.flow {
		flowToStrings = append(flowToStrings, fmt.Sprintf("%v%v", flowPrefix, flow.ToString(flowPrefix)))
	}
	return fmt.Sprintf("{\n%v\n%v}", strings.Join(flowToStrings, "\n"), prefix)
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
		_, suspend := step.Exec(space)
		if !nl_interface.IsNil(suspend) {
			return space, suspend
		}
	}
	return space, nil
}

func NewCodeBlock() *CodeBlock {
	return &CodeBlock{}
}
