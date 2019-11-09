package adaptor

import (
	"github.com/TingerSure/natural_language/tree"
)

type Adaptor struct{}

func (*Adaptor) GetWords(string) []*tree.Word {
	return nil
}

func (*Adaptor) GetVocabularyRules() []*tree.VocabularyRule {
	return nil
}

func (*Adaptor) GetStructRules() []*tree.StructRule {
	return nil
}
