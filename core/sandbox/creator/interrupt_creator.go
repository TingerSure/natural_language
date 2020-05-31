package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type InterruptCreator struct {
	Exception *interrupt.ExceptionCreator
}

type InterruptCreatorParam struct {
	StringCreator func(string) concept.String
}

func NewInterruptCreator(param *InterruptCreatorParam) *InterruptCreator {
	instance := &InterruptCreator{}
	instance.Exception = interrupt.NewExceptionCreator(&interrupt.ExceptionCreatorParam{
		StringCreator: param.StringCreator,
	})
	return instance
}
