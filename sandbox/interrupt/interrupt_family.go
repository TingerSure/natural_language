package interrupt

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type InterruptFamily struct {
}

func (v *InterruptFamily) IsException(value concept.Interrupt) (*Exception, bool) {
	if value == nil {
		return nil, false
	}
	if value.InterruptType() == ExceptionInterruptType {
		exception, yes := value.(*Exception)
		return exception, yes
	}
	return nil, false
}

func (v *InterruptFamily) IsEnd(value concept.Interrupt) (*End, bool) {
	if value == nil {
		return nil, false
	}
	if value.InterruptType() == EndInterruptType {
		end, yes := value.(*End)
		return end, yes
	}
	return nil, false
}

func newInterruptFamily() *InterruptFamily {
	return &InterruptFamily{}
}

var (
	InterruptFamilyInstance *InterruptFamily = newInterruptFamily()
)
