package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	EntityFromNumberName string = "structs.entity.number"
)

var (
	EntityFromNumberList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Number,
	}
)

type EntityFromNumber struct {
	*adaptor.SourceAdaptor
	NewAutoNumberContent concept.String
	NewAutoNumberObject  concept.String
	NewAutoNumber        concept.Function
}

func (p *EntityFromNumber) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								p.Libs.Sandbox.Index.ConstIndex.New(p.NewAutoNumber),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Index{
									p.NewAutoNumberContent: phrase[0].Index(),
								}),
							),
							p.NewAutoNumberObject,
						)
					},
					Size:  len(EntityFromNumberList),
					Types: phrase_type.Entity,
					From:  p.GetName(),
				})
			},
			Types: EntityFromNumberList,
			From:  p.GetName(),
		}),
	}
}

func (p *EntityFromNumber) GetName() string {
	return EntityFromNumberName
}

func NewEntityFromNumber(param *adaptor.SourceAdaptorParam) *EntityFromNumber {
	instance := (&EntityFromNumber{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	libObject := instance.Libs.GetLibraryPage("system", "auto_number")
	instance.NewAutoNumberContent = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("NewAutoNumberContent"))
	instance.NewAutoNumberObject = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("NewAutoNumberObject"))
	instance.NewAutoNumber = libObject.GetFunction(instance.Libs.Sandbox.Variable.String.New("NewAutoNumber"))
	return instance
}
