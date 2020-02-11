package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

const (
	HasMethodContent = "object"
	HasMethodKey     = "key"
	HasMethodExist   = "exist"
)

var (
	HasMethodObjectErrorException = interrupt.NewException("type error", "HasMethodObjectErrorException")
	HasMethodKeyErrorException    = interrupt.NewException("type error", "HasMethodKeyErrorException")
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

			return variable.NewParamWithInit(map[string]concept.Variable{
				HasMethodExist: variable.NewBool(object.HasMethod(key.Value())),
			}), nil
		},
		[]string{
			HasMethodContent,
			HasMethodKey,
		},
		[]string{
			HasMethodExist,
		},
	)
}
