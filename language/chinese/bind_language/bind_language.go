package bind_language

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/bind_language/library"
	"github.com/TingerSure/natural_language/language/chinese/bind_language/sandbox"
)

func BindLanguage(libs *tree.LibraryManager, language string) {
	sandbox.VariableBindLanguage(libs, language)
	sandbox.IndexBindLanguage(libs, language)
	sandbox.ExpressionBindLanguage(libs, language)
	library.SystemBindLanguage(libs, language)
}
