package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	EntityName string = "types.entity"
)

type Entity struct {
	*adaptor.SourceAdaptor
}

func (p *Entity) GetName() string {
	return EntityName
}

func (p *Entity) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: EntityName,
			From: EntityName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: AnyName,
				},
			},
		}),
	}
}

func NewEntity(param *adaptor.SourceAdaptorParam) *Entity {
	instance := (&Entity{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	return instance
}
