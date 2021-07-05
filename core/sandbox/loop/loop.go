package loop

import (
	"errors"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	LoopPrepare = 0
	LoopRunning = 1
	LoopStopped = 2
)

type Loop struct {
	events        chan *Event
	state         int
	execListener  func(concept.Function, concept.Variable, error)
	closeListener func()
}

func (l *Loop) Append(event *Event) {
	go func() {
		l.events <- event
	}()
}

func (l *Loop) OnClose(listener func()) {
	l.closeListener = listener
}

func (l *Loop) OnExec(listener func(concept.Function, concept.Variable, error)) {
	l.execListener = listener
}

func (l *Loop) Start() error {
	if l.state != LoopPrepare {
		return errors.New("Start failed. The Loop had been Started yet.")
	}
	l.state = LoopRunning
	go func() {
		for {
			event, ok := <-l.events
			if !ok {
				break
			}
			value, err := event.Exec()
			if l.execListener != nil {
				l.execListener(event.GetPipe(), value, err)
			}
		}
		if l.closeListener != nil {
			l.closeListener()
		}
	}()
	return nil
}

func (l *Loop) Stop() error {
	if l.state != LoopRunning {
		return errors.New("Start failed. The Loop is not running now.")
	}
	l.state = LoopStopped
	close(l.events)
	return nil
}

func NewLoop(chcheSize int) *Loop {
	return &Loop{
		state:  LoopPrepare,
		events: make(chan *Event, chcheSize),
	}
}
