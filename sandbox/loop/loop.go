package loop

import (
	"errors"
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
)

const (
	LoopPrepare = 0
	LoopRunning = 1
	LoopStopped = 2
)

type Loop struct {
	events            chan *Event
	state             int
	interruptListener func(concept.Interrupt)
	closeListener     func()
}

func (l *Loop) Append(event *Event) {
	go func() {
		l.events <- event
	}()
}

func (l *Loop) OnClose(listener func()) {
	l.closeListener = listener
}
func (l *Loop) OnInterrupt(listener func(concept.Interrupt)) {
	l.interruptListener = listener
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
			suspend := event.Exec()
			if !nl_interface.IsNil(suspend) {
				if suspend.InterruptType() == interrupt.EndInterruptType {
					break
				}
				if l.interruptListener != nil {
					l.interruptListener(suspend)
				}
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
