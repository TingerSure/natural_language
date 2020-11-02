package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	EntityFromEntityBelongNounName string = "structs.entity.entity_belong_noun"
)

var (
	EntityFromEntityBelongNounList []string = []string{
		phrase_type.EntityName,
		phrase_type.AuxiliaryBelongName,
		phrase_type.NounName,
	}
)

type EntityFromEntityBelongNoun struct {
	*adaptor.SourceAdaptor
	GetFieldValue   concept.String
	GetFieldKey     concept.String
	GetFieldContent concept.String
	GetField        concept.Function
}

func (p *EntityFromEntityBelongNoun) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								p.Libs.Sandbox.Index.ConstIndex.New(p.GetField),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Index{
									p.GetFieldContent: phrase[0].Index(),
									p.GetFieldKey:     phrase[2].Index(),
								}),
							),
							p.GetFieldValue,
						)
					},
					Size:  len(EntityFromEntityBelongNounList),
					Types: phrase_type.EntityName,
					From:  p.GetName(),
				})
			},
			Types: EntityFromEntityBelongNounList,
			From:  p.GetName(),
		}),
	}
}

func (p *EntityFromEntityBelongNoun) GetName() string {
	return EntityFromEntityBelongNounName
}

func NewEntityFromEntityBelongNoun(param *adaptor.SourceAdaptorParam) *EntityFromEntityBelongNoun {
	instance := (&EntityFromEntityBelongNoun{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	libObject := instance.Libs.GetLibraryPage("system", "object")

	instance.GetFieldValue = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldValue"))
	instance.GetFieldKey = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldKey"))
	instance.GetFieldContent = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldContent"))
	instance.GetField = libObject.GetFunction(instance.Libs.Sandbox.Variable.String.New("GetField"))

	return instance
}
