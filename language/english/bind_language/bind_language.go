package bind_language

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/language/english/bind_language/library"
	"github.com/TingerSure/natural_language/language/english/bind_language/sandbox"
)

func BindLanguage(libs *runtime.LibraryManager, language string) {
	sandbox.VariableBindLanguage(libs, language)
	sandbox.IndexBindLanguage(libs, language)
	sandbox.ExpressionBindLanguage(libs, language)
	library.SystemBindLanguage(libs, language)
}
