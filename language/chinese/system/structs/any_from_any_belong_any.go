package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	AnyFromAnyBelongAnyName string = "structs.any.any_belong_any"
)

var (
	anyFromAnyBelongAnyList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Any,
		phrase_type.AuxiliaryBelong,
		phrase_type.Any,
	}
)

type AnyFromAnyBelongAny struct {
	*adaptor.SourceAdaptor
	GetFieldValue   concept.String
	GetFieldKey     concept.String
	GetFieldContent concept.String
	GetField        concept.Function
}

func (p *AnyFromAnyBelongAny) GetStructRules() []*tree.StructRule {
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
					Size:  len(anyFromAnyBelongAnyList),
					Types: phrase_type.Any,
					From:  p.GetName(),
				})
			},
			Types: anyFromAnyBelongAnyList,
			From:  p.GetName(),
		}),
	}
}

func (p *AnyFromAnyBelongAny) GetName() string {
	return AnyFromAnyBelongAnyName
}

func NewAnyFromAnyBelongAny(param *adaptor.SourceAdaptorParam) *AnyFromAnyBelongAny {
	instance := (&AnyFromAnyBelongAny{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	libObject := instance.Libs.GetLibraryPage("system", "object")

	instance.GetFieldValue = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldValue"))
	instance.GetFieldKey = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldKey"))
	instance.GetFieldContent = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldContent"))
	instance.GetField = libObject.GetFunction(instance.Libs.Sandbox.Variable.String.New("GetField"))

	return instance
}
