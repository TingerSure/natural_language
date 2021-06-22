package loop

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Event struct {
	index concept.Function
	space concept.Pool
	line  concept.Line
}

func (e *Event) Exec() concept.Exception {
	resault, exception := e.index.Exec(nil, nil)
	if !nl_interface.IsNil(exception) {
		return exception.AddExceptionLine(e.line)
	}
	if nl_interface.IsNil(resault) {
		return nil
	}
	value := resault.GetOriginal("value")
	if nl_interface.IsNil(value) {
		return nil
	}
	if value.IsNull() {
		return nil
	}
	e.space.AddExtempore(e.index, value)
	return nil
}

func NewEvent(index concept.Function, line concept.Line, space concept.Pool) *Event {
	return &Event{
		index: index,
		space: space,
		line:  line,
	}
}
