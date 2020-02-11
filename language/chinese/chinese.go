package chinese

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system"
)

func NewChinese(libs *tree.LibraryManager) *tree.Language {
	chinese := tree.NewLanguage()
	chinese.SetPackage("system", system.NewSystem(libs))
	return chinese

}
