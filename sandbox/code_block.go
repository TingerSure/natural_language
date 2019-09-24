package sandbox

import (
	"errors"
)

type CodeBlock struct {
	flow []Expression
}

func (c *CodeBlock) AddStep(step Expression) {
	c.flow = append(c.flow, step)
}

func (f *CodeBlock) Exec(parent *Closure, returnBubble bool, init func(*Closure) error) (*Closure, bool, error) {
	if parent == nil && returnBubble {
		return nil, false, errors.New("No parent to bubble.")
	}
	space := NewClosure(parent)
	defer func() {
		if returnBubble {
			parent.MergeReturn(space)
		}
	}()
	if init != nil {
		err := init(space)
		if err != nil {
			return space, false, err
		}
	}
	for _, step := range f.flow {
		keep, err := step.Exec(space)
		if err != nil {
			return space, false, err
		}
		if !keep {
			return space, false, nil
		}
	}
	return space, true, nil
}

func NewCodeBlock() *CodeBlock {
	return &CodeBlock{}
}
