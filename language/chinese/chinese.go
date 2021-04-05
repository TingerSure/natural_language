package chinese

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/bind_language"
	"github.com/TingerSure/natural_language/language/chinese/system"
)

func BindRule(libs *tree.LibraryManager, chineseName string) {
	system.BindRule(libs, chineseName)
}

func BindLanguage(libs *tree.LibraryManager, chineseName string) {
	bind_language.BindLanguage(libs, chineseName)
}
