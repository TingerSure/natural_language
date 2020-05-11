package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func ObjectBindlanguage(libs *tree.LibraryManager, language string) {
	instance := libs.GetLibraryPage("system", "object")

	// GetFieldObjectErrorException := instance.GetException(variable.NewString("GetFieldObjectErrorException"))
	// GetFieldKeyErrorException := instance.GetException(variable.NewString("GetFieldKeyErrorException"))
	// GetFieldKeyNotExistException := instance.GetException(variable.NewString("GetFieldKeyNotExistException"))

	GetFieldContent := instance.GetConst(variable.NewString("GetFieldContent"))
	GetFieldKey := instance.GetConst(variable.NewString("GetFieldKey"))
	GetFieldValue := instance.GetConst(variable.NewString("GetFieldValue"))

	GetField := instance.GetFunction(variable.NewString("GetField"))

	GetFieldContent.SetLanguage(language, "object")
	GetFieldKey.SetLanguage(language, "field")
	GetFieldValue.SetLanguage(language, "value")
	GetField.Name().SetLanguage(language, "get field")
	GetField.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		content := param.Get(GetFieldContent).(concept.ToString)
		key := param.Get(GetFieldKey).(concept.ToString)
		return fmt.Sprintf("%v of %v", key.ToLanguage(language), content.ToLanguage(language))

	})
}
