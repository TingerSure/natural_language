package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

const (
	HasFieldContent = "object"
	HasFieldKey     = "key"
	HasFieldExist   = "exist"
)

var (
	HasFieldObjectErrorException = interrupt.NewException("type error", "HasFieldObjectErrorException")
	HasFieldKeyErrorException    = interrupt.NewException("type error", "HasFieldKeyErrorException")
)

var (
	HasField *variable.SystemFunction = nil
)

func init() {
	HasField = variable.NewSystemFunction(
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(HasFieldContent))
			if !ok {
				return nil, HasFieldObjectErrorException.Copy().AddStack(HasField)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(HasFieldKey))
			if !ok {
				return nil, HasFieldKeyErrorException.Copy().AddStack(HasField)
			}

			return variable.NewParamWithInit(map[string]concept.Variable{
				HasFieldExist: variable.NewBool(object.HasField(key.Value())),
			}), nil
		},
		[]string{
			HasFieldContent,
			HasFieldKey,
		},
		[]string{
			HasFieldExist,
		},
	)
}
