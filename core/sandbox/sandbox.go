package sandbox

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/loop"
)

type SandboxParam struct {
	OnError   func(error)
	OnClose   func()
	OnPrint   func(concept.Variable)
	OnExec    func(concept.Function, concept.Variable)
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

func (s *Sandbox) Exec(index concept.Function, line concept.Line) {
	s.eventLoop.Append(loop.NewEvent(index, line))
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
	box.eventLoop.OnExec(func(pipe concept.Function, value concept.Variable, err error) {
		if err != nil {
			param.OnError(err)
			return
		}
		param.OnExec(pipe, value)
	})
	return box
}
