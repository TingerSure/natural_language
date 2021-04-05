package english

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/english/bind_language"
)

func BindRule(libs *tree.LibraryManager, chineseName string) {

}

func BindLanguage(libs *tree.LibraryManager, englishName string) {
	bind_language.BindLanguage(libs, englishName)
}
