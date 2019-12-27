package code_block

import (
	"fmt"
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/closure"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"strings"
)

type CodeBlock struct {
	flow []concept.Index
}

func (c *CodeBlock) Size() int {
	return len(c.flow)
}

const (
	CodeBlockFlowNoneSize     = 0
	CodeBlockFlowOnlyOneSize  = 1
	CodeBlockFlowOnlyOneIndex = 0
)

func (c *CodeBlock) Iterate(onIndex func(concept.Index) bool) bool {
	for _, index := range c.flow {
		if onIndex(index) || index.SubCodeBlockIterate(onIndex) {
			return true
		}
	}
	return false
}

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
func (c *CodeBlock) AddStep(step ...concept.Index) {
	c.flow = append(c.flow, step...)
}

func (c *CodeBlock) AddSteps(steps []concept.Index) {
	c.flow = append(c.flow, steps...)
}

func (f *CodeBlock) Exec(
	parent concept.Closure,
	returnBubble bool,
	init func(concept.Closure) concept.Interrupt,
) (concept.Closure, concept.Interrupt) {

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
		value, suspend := step.Get(space)
		if !nl_interface.IsNil(suspend) {
			return space, suspend
		}
		space.AddExtempore(step, value)
	}
	return space, nil
}

func NewCodeBlock() *CodeBlock {
	return &CodeBlock{}
}
