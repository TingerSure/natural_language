package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type InterruptCreator struct {
	Exception *interrupt.ExceptionCreator
	End       *interrupt.EndCreator
	Continue  *interrupt.ContinueCreator
	Break     *interrupt.BreakCreator
}

type InterruptCreatorParam struct {
	StringCreator func(string) concept.String
}

func NewInterruptCreator(param *InterruptCreatorParam) *InterruptCreator {
	instance := &InterruptCreator{}
	instance.Break = interrupt.NewBreakCreator(&interrupt.BreakCreatorParam{})
	instance.Continue = interrupt.NewContinueCreator(&interrupt.ContinueCreatorParam{})
	instance.End = interrupt.NewEndCreator(&interrupt.EndCreatorParam{})
	instance.Exception = interrupt.NewExceptionCreator(&interrupt.ExceptionCreatorParam{
		StringCreator: param.StringCreator,
	})
	return instance
}
