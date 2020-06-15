package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	CreateContentName = "object"
)

func initCreate(instance *Object) {

	CreateContent := libs.Sandbox.Variable.String.New(CreateContentName)

	instance.SetConst(libs.Sandbox.Variable.String.New("CreateContent"), CreateContent)

	instance.SetFunction(libs.Sandbox.Variable.String.New("Create"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("Create"),
		func(_ concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			return libs.Sandbox.Variable.Param.New().Set(CreateContent, variable.NewObject()), nil
		},
		[]concept.String{},
		[]concept.String{
			CreateContent,
		},
	))
}
