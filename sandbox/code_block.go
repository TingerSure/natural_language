package sandbox

import (
	"errors"
)

type CodeBlock struct {
	flow         []Expression
	init         func(*Closure)
	returnBubble bool
}

func (f *CodeBlock) AddStep(step Expression) {
	f.flow = append(f.flow, step)
}

func (f *CodeBlock) Exec(parent *Closure, flow []Expression) (*Closure, bool, error) {
	if parent == nil && f.returnBubble {
		return nil, false, errors.New("No parent to bubble.")
	}
	space := NewClosure(parent)
	defer func() {
		if f.returnBubble {
			parent.MergeReturn(space)
		}
		space.Clear()
	}()
	for _, step := range flow {
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

func NewCodeBlock(init func(*Closure), returnBubble bool) *CodeBlock {
	return &CodeBlock{
		init:         init,
		returnBubble: returnBubble,
	}
}
