package chinese

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system"
)

const (
	ChineseName = "chinese"
)

func NewChinese(libs *tree.LibraryManager) *tree.Language {
	chinese := tree.NewLanguage()
	chinese.SetName(ChineseName)
	chinese.SetPackage("system", system.NewSystem(libs))
	return chinese

}
