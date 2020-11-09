package object

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	HasMethodContentName = "object"
	HasMethodKeyName     = "key"
	HasMethodExistName   = "exist"
)

func initHasMethod(libs *runtime.LibraryManager, instance *Object) {
	HasMethodContent := libs.Sandbox.Variable.String.New(HasMethodContentName)
	HasMethodKey := libs.Sandbox.Variable.String.New(HasMethodKeyName)
	HasMethodExist := libs.Sandbox.Variable.String.New(HasMethodExistName)

	HasMethodObjectErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "HasMethodObjectErrorException")
	HasMethodKeyErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "HasMethodKeyErrorException")

	var HasMethod concept.Function
	HasMethod = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("HasMethod"),
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(HasMethodContent))
			if !ok {
				return nil, HasMethodObjectErrorException.Copy().AddStack(HasMethod)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(HasMethodKey))
			if !ok {
				return nil, HasMethodKeyErrorException.Copy().AddStack(HasMethod)
			}

			return libs.Sandbox.Variable.Param.New().Set(HasMethodExist, libs.Sandbox.Variable.Bool.New(object.HasMethod(key))), nil
		},
		func(_ concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(HasMethodExist, libs.Sandbox.Variable.Bool.New(false))
		},
		[]concept.String{
			HasMethodContent,
			HasMethodKey,
		},
		[]concept.String{
			HasMethodExist,
		},
	)

	instance.SetException(libs.Sandbox.Variable.String.New("HasMethodObjectErrorException"), HasMethodObjectErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("HasMethodKeyErrorException"), HasMethodKeyErrorException)

	instance.SetConst(libs.Sandbox.Variable.String.New("HasMethodContent"), HasMethodContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("HasMethodKey"), HasMethodKey)
	instance.SetConst(libs.Sandbox.Variable.String.New("HasMethodExist"), HasMethodExist)

	instance.SetFunction(libs.Sandbox.Variable.String.New("HasMethod"), HasMethod)

}
