package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	HasFieldContentName = "object"
	HasFieldKeyName     = "key"
	HasFieldExistName   = "exist"

	HasFieldObjectErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasFieldObjectErrorException"))
	HasFieldKeyErrorExceptionTemplate    = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasFieldKeyErrorException"))
)

func initHasField(instance *Object) {

	HasFieldContent := variable.NewString(HasFieldContentName)
	HasFieldKey := variable.NewString(HasFieldKeyName)
	HasFieldExist := variable.NewString(HasFieldExistName)

	HasFieldObjectErrorException := HasFieldObjectErrorExceptionTemplate.Copy()
	HasFieldKeyErrorException := HasFieldKeyErrorExceptionTemplate.Copy()

	var HasField concept.Function
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

	instance.SetException(variable.NewString("HasFieldObjectErrorException"), HasFieldObjectErrorException)
	instance.SetException(variable.NewString("HasFieldKeyErrorException"), HasFieldKeyErrorException)

	instance.SetConst(variable.NewString("HasFieldContent"), HasFieldContent)
	instance.SetConst(variable.NewString("HasFieldKey"), HasFieldKey)
	instance.SetConst(variable.NewString("HasFieldExist"), HasFieldExist)

	instance.SetFunction(variable.NewString("HasField"), HasField)

}
