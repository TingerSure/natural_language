package sandbox

import (
	"errors"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/loop"
)

type SandboxParam struct {
	OnError   func(error)
	OnClose   func()
	EventSize int
}

type Sandbox struct {
	eventLoop *loop.Loop
	param     *SandboxParam
}

func (s *Sandbox) GetEventLoop() *loop.Loop {
	return s.eventLoop
}

func (s *Sandbox) Exec(funcs concept.Function, param concept.Param) {
	s.eventLoop.Append(loop.NewEvent(funcs, param))
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
	box.eventLoop.OnException(func(exception concept.Exception) {
		param.OnError(errors.New(exception.ToString("")))
	})

	return box
}
