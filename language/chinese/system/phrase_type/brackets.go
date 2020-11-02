package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	BracketsName      string = "types.brackets"
	BracketsLeftName  string = "types.brackets.left"
	BracketsRightName string = "types.brackets.right"
)

type Brackets struct {
	*adaptor.SourceAdaptor
}

func (p *Brackets) GetName() string {
	return BracketsName
}

func (p *Brackets) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    BracketsLeftName,
			From:    BracketsName,
			Parents: nil,
		}),
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name:    BracketsRightName,
			From:    BracketsName,
			Parents: nil,
		}),
	}
}

func NewBrackets(param *adaptor.SourceAdaptorParam) *Brackets {
	instance := (&Brackets{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
