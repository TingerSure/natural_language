package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	HasFieldContent = variable.NewString("object")
	HasFieldKey     = variable.NewString("key")
	HasFieldExist   = variable.NewString("exist")
)

var (
	HasFieldObjectErrorException = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasFieldObjectErrorException"))
	HasFieldKeyErrorException    = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasFieldKeyErrorException"))
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

			return variable.NewParam().Set(HasFieldExist, variable.NewBool(object.HasField(key))), nil
		},
		[]concept.String{
			HasFieldContent,
			HasFieldKey,
		},
		[]concept.String{
			HasFieldExist,
		},
	)
}
