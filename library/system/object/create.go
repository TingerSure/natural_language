package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	CreateContentName = "object"
)

func initCreate(libs *tree.LibraryManager, instance *Object) {

	CreateContent := libs.Sandbox.Variable.String.New(CreateContentName)

	instance.SetConst(libs.Sandbox.Variable.String.New("CreateContent"), CreateContent)

	anticipateObject := libs.Sandbox.Variable.Object.New()

	instance.SetFunction(libs.Sandbox.Variable.String.New("Create"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("Create"),
		func(_ concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			return libs.Sandbox.Variable.Param.New().Set(CreateContent, libs.Sandbox.Variable.Object.New()), nil
		},
		func(_ concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(CreateContent, anticipateObject)
		},
		[]concept.String{},
		[]concept.String{
			CreateContent,
		},
	))
}
