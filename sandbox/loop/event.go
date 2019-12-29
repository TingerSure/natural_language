package loop

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type Event struct {
	index concept.Index
	space concept.Closure
}

func (e *Event) Exec() concept.Interrupt {
	resault, suspend := e.index.Get(e.space)
	if nl_interface.IsNil(suspend) {
		e.space.AddExtempore(e.index, resault)
		return nil
	}
	return suspend
}

func NewEvent(index concept.Index, space concept.Closure) *Event {

	return &Event{
		index: index,
		space: space,
	}
}
