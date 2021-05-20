package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type InterruptCreator struct {
	Return   *interrupt.ReturnCreator
	Continue *interrupt.ContinueCreator
	Break    *interrupt.BreakCreator
}

type InterruptCreatorParam struct {
}

func NewInterruptCreator(param *InterruptCreatorParam) *InterruptCreator {
	instance := &InterruptCreator{}
	instance.Break = interrupt.NewBreakCreator(&interrupt.BreakCreatorParam{})
	instance.Continue = interrupt.NewContinueCreator(&interrupt.ContinueCreatorParam{})
	instance.Return = interrupt.NewReturnCreator(&interrupt.ReturnCreatorParam{})
	return instance
}
