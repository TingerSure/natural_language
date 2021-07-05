package loop

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Event struct {
	pipe concept.Function
	line concept.Line
}

func (e *Event) GetPipe() concept.Function {
	return e.pipe
}

func (e *Event) Exec() (concept.Variable, error) {
	resault, exception := e.pipe.Exec(nil, nil)
	if !nl_interface.IsNil(exception) {
		return nil, exception.AddExceptionLine(e.line)
	}
	if nl_interface.IsNil(resault) {
		return nil, fmt.Errorf("Event pipe has no returns:\n%v", e.pipe.ToString(""))
	}
	value := resault.GetOriginal("value")
	return value, nil
}

func NewEvent(pipe concept.Function, line concept.Line) *Event {
	return &Event{
		pipe: pipe,
		line: line,
	}
}
