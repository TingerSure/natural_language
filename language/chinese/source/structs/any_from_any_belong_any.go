package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/source/phrase_type"
	"github.com/TingerSure/natural_language/library/object"
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
}

func (p *AnyFromAnyBelongAny) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return expression.NewParamGet(
							expression.NewCall(
								index.NewConstIndex(object.GetField),
								expression.NewNewParamWithInit(map[string]concept.Index{
									object.GetFieldContent: phrase[0].Index(),
									object.GetFieldKey:     phrase[2].Index(),
								}),
							),
							object.GetFieldValue,
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

func NewAnyFromAnyBelongAny() *AnyFromAnyBelongAny {
	return (&AnyFromAnyBelongAny{})
}
