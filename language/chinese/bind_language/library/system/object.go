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

	GetFieldContent.SetLanguage(language, "对象")
	GetFieldKey.SetLanguage(language, "属性")
	GetFieldValue.SetLanguage(language, "值")
	GetField.Name().SetLanguage(language, "取值")
	GetField.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		content := param.Get(GetFieldContent).(concept.ToString)
		key := param.Get(GetFieldKey).(concept.ToString)
		return fmt.Sprintf("%v的%v", content.ToLanguage(language), key.ToLanguage(language))

	})
}
