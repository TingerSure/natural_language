package adaptor

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type SourceAdaptor struct {
	Libs *tree.LibraryManager
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
