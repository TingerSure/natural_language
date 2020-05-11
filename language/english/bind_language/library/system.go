package library

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/english/bind_language/library/system"
)

func SystemBindLanguage(libs *tree.LibraryManager, language string) {
	system.OperatorBindLanguage(libs, language)
	system.ObjectBindlanguage(libs, language)
	system.QuestionBindLanguage(libs, language)
}
