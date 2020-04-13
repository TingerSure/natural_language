package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	CreateContentName = "object"
)

func initCreate(instance *Object) {

	CreateContent := variable.NewString(CreateContentName)

	instance.SetConst(variable.NewString("CreateContent"), CreateContent)

	instance.SetFunction(variable.NewString("Create"), variable.NewSystemFunction(
		func(_ concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			return variable.NewParam().Set(CreateContent, variable.NewObject()), nil
		},
		[]concept.String{},
		[]concept.String{
			CreateContent,
		},
	))
}
