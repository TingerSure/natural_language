package library

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/english/bind_language/library/system"
)

func SystemBindLanguage(libs *runtime.LibraryManager, language string) {
	system.OperatorBindLanguage(libs, language)
	system.ObjectBindlanguage(libs, language)
	system.QuestionBindLanguage(libs, language)
}
