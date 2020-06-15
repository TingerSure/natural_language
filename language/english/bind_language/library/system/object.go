package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

func ObjectBindlanguage(libs *runtime.LibraryManager, language string) {
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
	GetField.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		content := param.Get(GetFieldContent).(concept.ToString)
		key := param.Get(GetFieldKey).(concept.ToString)
		return fmt.Sprintf("the %v of %v", key.ToLanguage(language), content.ToLanguage(language))

	})
}
