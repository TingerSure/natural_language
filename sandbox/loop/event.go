package loop

import (
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

var (
	eventDefaultParam = variable.NewParam()
)

type Event struct {
	param concept.Param
	exec  concept.Function
}

func (e *Event) Exec() concept.Exception {
	_, exception := e.exec.Exec(e.param, nil)
	return exception
}

func NewEvent(exec concept.Function, param concept.Param) *Event {
	if nl_interface.IsNil(param) {
		param = eventDefaultParam
	}
	return &Event{
		param: param,
		exec:  exec,
	}
}
