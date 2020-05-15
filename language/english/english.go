package english

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/english/bind_language"
)

func NewEnglish(libs *runtime.LibraryManager, englishName string) tree.Library {
	return tree.NewLibraryAdaptor()
}

func EnglishBindLanguage(libs *runtime.LibraryManager, englishName string) {
	bind_language.BindLanguage(libs, englishName)
}
