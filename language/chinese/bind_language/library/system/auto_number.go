package system

import (
	"github.com/TingerSure/natural_language/core/runtime"
)

func AutoNumberBindlanguage(libs *runtime.LibraryManager, language string) {
	instance := libs.GetLibraryPage("system", "auto_number")
	//TODO
	AutoNumberValue := instance.GetConst(libs.Sandbox.Variable.String.New("AutoNumberValue"))
	AutoNumberClassValue := instance.GetConst(libs.Sandbox.Variable.String.New("AutoNumberClassValue"))

	AutoNumberValue.SetLanguage(language, "值")
	AutoNumberClassValue.SetLanguage(language, "值")

}
