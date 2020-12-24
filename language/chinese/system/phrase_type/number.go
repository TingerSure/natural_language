package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	NumberName string = "types.number"
)

type Number struct {
	*adaptor.SourceAdaptor
}

func (p *Number) GetName() string {
	return NumberName
}

func (p *Number) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: NumberName,
			From: NumberName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: EntityName,
					Rule: tree.NewPackageRule(&tree.PackageRuleParam{
						Pack: func(value tree.Phrase) tree.Phrase {
							return value
							//TODO
						},
						From: NumberName,
					}),
				},
			},
		}),
	}
}

func NewNumber(param *adaptor.SourceAdaptorParam) *Number {
	instance := &Number{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	}
	return instance
}
