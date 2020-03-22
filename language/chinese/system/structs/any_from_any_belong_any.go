package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
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
	adaptor.SourceAdaptor
	libs            *tree.LibraryManager
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
						return expression.NewParamGet(
							expression.NewCall(
								index.NewConstIndex(p.GetField),
								expression.NewNewParamWithInit(map[concept.String]concept.Index{
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

func NewAnyFromAnyBelongAny(libs *tree.LibraryManager) *AnyFromAnyBelongAny {
	libObject := libs.GetLibraryPage("system", "object")
	return (&AnyFromAnyBelongAny{
		libs:            libs,
		GetFieldValue:   libObject.GetConst(variable.NewString("GetFieldValue")),
		GetFieldKey:     libObject.GetConst(variable.NewString("GetFieldKey")),
		GetFieldContent: libObject.GetConst(variable.NewString("GetFieldContent")),
		GetField:        libObject.GetFunction(variable.NewString("GetField")),
	})
}
