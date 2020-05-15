package adaptor

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/tree"
)

type SourceAdaptorParam struct {
	Libs     *runtime.LibraryManager
	Language string
}

type SourceAdaptor struct {
	Libs     *runtime.LibraryManager
	Language string
}

func NewSourceAdaptor(param *SourceAdaptorParam) *SourceAdaptor {
	return &SourceAdaptor{
		Libs:     param.Libs,
		Language: param.Language,
	}
}

func (*SourceAdaptor) GetWords(string) []*tree.Word {
	return nil
}

func (*SourceAdaptor) GetVocabularyRules() []*tree.VocabularyRule {
	return nil
}

func (*SourceAdaptor) GetStructRules() []*tree.StructRule {
	return nil
}

func (*SourceAdaptor) GetPriorityRules() []*tree.PriorityRule {
	return nil
}
