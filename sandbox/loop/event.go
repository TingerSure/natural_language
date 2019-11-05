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
	param           concept.Param
	exec            concept.Function
	resaultListener func(value concept.Variable)
}

func (e *Event) Exec() concept.Exception {
	resault, exception := e.exec.Exec(e.param, nil)
	if e.resaultListener != nil {
		e.resaultListener(resault)
	}
	return exception
}

func NewEvent(exec concept.Function, param concept.Param, resaultListener func(value concept.Variable)) *Event {
	if nl_interface.IsNil(param) {
		param = eventDefaultParam
	}
	return &Event{
		param:           param,
		exec:            exec,
		resaultListener: resaultListener,
	}
}
