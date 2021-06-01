package loop

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Event struct {
	index concept.Pipe
	space concept.Pool
}

func (e *Event) Exec() concept.Interrupt {
	resault, suspend := e.index.Get(e.space)
	if nl_interface.IsNil(suspend) {
		e.space.AddExtempore(e.index, resault)
		return nil
	}
	return suspend
}

func NewEvent(index concept.Pipe, space concept.Pool) *Event {

	return &Event{
		index: index,
		space: space,
	}
}
