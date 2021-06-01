package sandbox

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/loop"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type SandboxParam struct {
	OnError   func(error)
	OnClose   func()
	OnPrint   func(concept.Variable)
	Root      concept.Pool
	EventSize int
}

type Sandbox struct {
	eventLoop *loop.Loop
	param     *SandboxParam
}

func (s *Sandbox) GetEventLoop() *loop.Loop {
	return s.eventLoop
}

func (s *Sandbox) Exec(index concept.Pipe) {
	s.eventLoop.Append(loop.NewEvent(index, s.param.Root))
}

func (s *Sandbox) Start() error {
	return s.eventLoop.Start()
}

func (s *Sandbox) Stop() error {
	return s.eventLoop.Stop()
}

func NewSandbox(param *SandboxParam) *Sandbox {
	box := (&Sandbox{
		eventLoop: loop.NewLoop(param.EventSize),
		param:     param,
	})

	box.eventLoop.OnClose(param.OnClose)
	box.eventLoop.OnInterrupt(func(suspend concept.Interrupt) {
		switch suspend.InterruptType() {
		case variable.ExceptionInterruptType:
			param.OnError(errors.New(suspend.(concept.Exception).ToString("")))
		default:
			param.OnError(fmt.Errorf("Illegel interrupt in the root loop: %v.", suspend.InterruptType()))
		}
	})

	return box
}
