package chinese

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

func NewChinese(libs *tree.LibraryManager, chineseName string) tree.Library {
	chinese := tree.NewLibraryAdaptor()
	chinese.SetPage("system", system.NewSystem(&adaptor.SourceAdaptorParam{
		Libs:     libs,
		Language: chineseName,
	}))
	return chinese
}
