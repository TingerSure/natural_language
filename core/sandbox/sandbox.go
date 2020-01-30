package sandbox

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/loop"
	"github.com/TingerSure/natural_language/library/std"
)

type sandboxStd struct {
	param *SandboxParam
}

func (s *sandboxStd) Print(value concept.Variable) {
	s.param.OnPrint(value)
}

func (s *sandboxStd) Error(value concept.Variable) {
	s.param.OnError(errors.New(value.ToString("")))
}

func newSandboxStd(param *SandboxParam) *sandboxStd {
	return &sandboxStd{
		param: param,
	}
}

type SandboxParam struct {
	OnError   func(error)
	OnClose   func()
	OnPrint   func(concept.Variable)
	Root      concept.Closure
	EventSize int
}

type Sandbox struct {
	eventLoop *loop.Loop
	param     *SandboxParam
}

func (s *Sandbox) GetEventLoop() *loop.Loop {
	return s.eventLoop
}

func (s *Sandbox) Exec(index concept.Index) {
	s.eventLoop.Append(loop.NewEvent(index, s.param.Root))
}

func (s *Sandbox) Start() error {
	return s.eventLoop.Start()
}

func (s *Sandbox) Stop() error {
	return s.eventLoop.Start()
}

func NewSandbox(param *SandboxParam) *Sandbox {
	box := (&Sandbox{
		eventLoop: loop.NewLoop(param.EventSize),
		param:     param,
	})

	box.eventLoop.OnClose(param.OnClose)
	box.eventLoop.OnInterrupt(func(suspend concept.Interrupt) {
		switch suspend.InterruptType() {
		case interrupt.ExceptionInterruptType:
			param.OnError(errors.New(suspend.(concept.Exception).ToString("")))
		default:
			param.OnError(errors.New(fmt.Sprintf("Illegel interrupt in the root loop: %v.", suspend.InterruptType())))
		}
	})

	std.Std = newSandboxStd(param)

	return box
}
