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

func (v *InterruptFamily) IsBreak(value concept.Interrupt) (*Break, bool) {
	if value == nil {
		return nil, false
	}
	if value.InterruptType() == BreakInterruptType {
		breaks, yes := value.(*Break)
		return breaks, yes
	}
	return nil, false
}

func (v *InterruptFamily) IsContinue(value concept.Interrupt) (*Continue, bool) {
	if value == nil {
		return nil, false
	}
	if value.InterruptType() == ContinueInterruptType {
		continues, yes := value.(*Continue)
		return continues, yes
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
