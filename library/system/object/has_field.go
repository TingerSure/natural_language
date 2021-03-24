package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	HasFieldContentName = "object"
	HasFieldKeyName     = "key"
	HasFieldExistName   = "exist"
)

func initHasField(libs *tree.LibraryManager, instance *Object) {

	HasFieldContent := libs.Sandbox.Variable.String.New(HasFieldContentName)
	HasFieldKey := libs.Sandbox.Variable.String.New(HasFieldKeyName)
	HasFieldExist := libs.Sandbox.Variable.String.New(HasFieldExistName)

	HasFieldObjectErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "HasFieldObjectErrorException")
	HasFieldKeyErrorException := libs.Sandbox.Interrupt.Exception.NewOriginal("type error", "HasFieldKeyErrorException")

	var HasField concept.Function
	HasField = libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("HasField"),
		func(input concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			object, ok := variable.VariableFamilyInstance.IsObjectHome(input.Get(HasFieldContent))
			if !ok {
				return nil, HasFieldObjectErrorException.Copy().AddStack(HasField)
			}

			key, ok := variable.VariableFamilyInstance.IsString(input.Get(HasFieldKey))
			if !ok {
				return nil, HasFieldKeyErrorException.Copy().AddStack(HasField)
			}

			return libs.Sandbox.Variable.Param.New().Set(HasFieldExist, libs.Sandbox.Variable.Bool.New(object.HasField(key))), nil
		},
		func(_ concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(HasFieldExist, libs.Sandbox.Variable.Bool.New(false))
		},
		[]concept.String{
			HasFieldContent,
			HasFieldKey,
		},
		[]concept.String{
			HasFieldExist,
		},
	)

	instance.SetException(libs.Sandbox.Variable.String.New("HasFieldObjectErrorException"), HasFieldObjectErrorException)
	instance.SetException(libs.Sandbox.Variable.String.New("HasFieldKeyErrorException"), HasFieldKeyErrorException)

	instance.SetConst(libs.Sandbox.Variable.String.New("HasFieldContent"), HasFieldContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("HasFieldKey"), HasFieldKey)
	instance.SetConst(libs.Sandbox.Variable.String.New("HasFieldExist"), HasFieldExist)

	instance.SetFunction(libs.Sandbox.Variable.String.New("HasField"), HasField)

}
