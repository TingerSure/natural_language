package sandbox

import (
	// "fmt"
	// "github.com/TingerSure/natural_language/library/nl_interface"
)

type CodeBlock struct {
	flow []Expression
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
		if suspend != nil {
			return space, suspend
		}
	}
	return space, nil
}

func NewCodeBlock() *CodeBlock {
	return &CodeBlock{}
}
