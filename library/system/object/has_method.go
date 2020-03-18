package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	HasMethodContent = variable.NewString("object")
	HasMethodKey     = variable.NewString("key")
	HasMethodExist   = variable.NewString("exist")
)

var (
	HasMethodObjectErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasMethodObjectErrorException"))
	HasMethodKeyErrorException    = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasMethodKeyErrorException"))
)

var (
	HasMethod *variable.SystemFunction = nil
)

func init() {
	HasMethod = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(HasMethodContent))
			if !ok {
				return nil, HasMethodObjectErrorException.Copy().AddStack(HasMethod)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(HasMethodKey))
			if !ok {
				return nil, HasMethodKeyErrorException.Copy().AddStack(HasMethod)
			}

			return variable.NewParam().Set(HasMethodExist, variable.NewBool(object.HasMethod(key))), nil
		},
		[]concept.String{
			HasMethodContent,
			HasMethodKey,
		},
		[]concept.String{
			HasMethodExist,
		},
	)
}
