package system

import (
	"github.com/TingerSure/natural_language/core/tree"
)

func NumberBindlanguage(libs *tree.LibraryManager, language string) {
	instance := libs.GetLibraryPage("system", "number")
	//TODO
	NumberValue := instance.GetConst(libs.Sandbox.Variable.String.New("NumberValue"))
	NumberClassValue := instance.GetConst(libs.Sandbox.Variable.String.New("NumberClassValue"))

	NumberValue.SetLanguage(language, "值")
	NumberClassValue.SetLanguage(language, "值")

}
