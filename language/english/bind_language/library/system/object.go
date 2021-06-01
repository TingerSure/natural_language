package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

func ObjectBindlanguage(libs *tree.LibraryManager, language string) {
	instance := libs.GetLibraryPage("system", "object")

	// GetFieldObjectErrorException := instance.GetException(libs.Sandbox.Variable.String.New("GetFieldObjectErrorException"))
	// GetFieldKeyErrorException := instance.GetException(libs.Sandbox.Variable.String.New("GetFieldKeyErrorException"))
	// GetFieldKeyNotExistException := instance.GetException(libs.Sandbox.Variable.String.New("GetFieldKeyNotExistException"))

	GetFieldContent := instance.GetConst(libs.Sandbox.Variable.String.New("GetFieldContent"))
	GetFieldKey := instance.GetConst(libs.Sandbox.Variable.String.New("GetFieldKey"))
	GetFieldValue := instance.GetConst(libs.Sandbox.Variable.String.New("GetFieldValue"))

	GetField := instance.GetFunction(libs.Sandbox.Variable.String.New("GetField"))

	GetFieldContent.SetLanguage(language, "object")
	GetFieldKey.SetLanguage(language, "field")
	GetFieldValue.SetLanguage(language, "value")
	GetField.Name().SetLanguage(language, "get field")
	GetField.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *nl_interface.Mapping) string {
		content := param.Get(GetFieldContent).(concept.ToString)
		key := param.Get(GetFieldKey).(concept.ToString)
		return fmt.Sprintf("the %v of %v", key.ToLanguage(language), content.ToLanguage(language))

	})
}
