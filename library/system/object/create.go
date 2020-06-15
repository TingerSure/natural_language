package object

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

var (
	CreateContentName = "object"
)

func initCreate(libs *runtime.LibraryManager, instance *Object) {

	CreateContent := libs.Sandbox.Variable.String.New(CreateContentName)

	instance.SetConst(libs.Sandbox.Variable.String.New("CreateContent"), CreateContent)

	instance.SetFunction(libs.Sandbox.Variable.String.New("Create"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("Create"),
		func(_ concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			return libs.Sandbox.Variable.Param.New().Set(CreateContent, libs.Sandbox.Variable.Object.New()), nil
		},
		[]concept.String{},
		[]concept.String{
			CreateContent,
		},
	))
}
