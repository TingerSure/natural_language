package chinese

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system"
)

const (
	ChineseName = "chinese"
)

func NewChinese(libs *tree.LibraryManager) tree.Library {
	chinese := tree.NewLibraryAdaptor()
	chinese.SetPage("system", system.NewSystem(libs))
	return chinese
}
