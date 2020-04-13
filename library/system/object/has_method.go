package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	HasMethodContentName = "object"
	HasMethodKeyName     = "key"
	HasMethodExistName   = "exist"

	HasMethodObjectErrorExceptionTemplate = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasMethodObjectErrorException"))
	HasMethodKeyErrorExceptionTemplate    = interrupt.NewException(variable.NewString("type error"), variable.NewString("HasMethodKeyErrorException"))
)

func initHasMethod(instance *Object) {
	HasMethodContent := variable.NewString(HasMethodContentName)
	HasMethodKey := variable.NewString(HasMethodKeyName)
	HasMethodExist := variable.NewString(HasMethodExistName)

	HasMethodObjectErrorException := HasMethodObjectErrorExceptionTemplate.Copy()
	HasMethodKeyErrorException := HasMethodKeyErrorExceptionTemplate.Copy()

	var HasMethod concept.Function
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

	instance.SetException(variable.NewString("HasMethodObjectErrorException"), HasMethodObjectErrorException)
	instance.SetException(variable.NewString("HasMethodKeyErrorException"), HasMethodKeyErrorException)

	instance.SetConst(variable.NewString("HasMethodContent"), HasMethodContent)
	instance.SetConst(variable.NewString("HasMethodKey"), HasMethodKey)
	instance.SetConst(variable.NewString("HasMethodExist"), HasMethodExist)

	instance.SetFunction(variable.NewString("HasMethod"), HasMethod)

}
