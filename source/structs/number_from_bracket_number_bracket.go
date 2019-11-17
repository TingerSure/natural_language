package structs

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	NumberFromBracketNumberBracketName string = "structs.number.bracket_number_bracket"
)

var (
	NumberFromBracketNumberBracketList []string = []string{
		phrase_types.BracketsLeft,
		phrase_types.Number,
		phrase_types.BracketsRight,
	}
)

type NumberFromBracketNumberBracket struct {
	adaptor.SourceAdaptor
}

func (p *NumberFromBracketNumberBracket) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(func(phrase []tree.Phrase) concept.Index {
					return phrase[1].Index()
				}, len(NumberFromBracketNumberBracketList), phrase_types.Number, p.GetName())
			},
			Types: NumberFromBracketNumberBracketList,
			From:  p.GetName(),
		}),
	}
}

func (p *NumberFromBracketNumberBracket) GetName() string {
	return NumberFromBracketNumberBracketName
}

func NewNumberFromBracketNumberBracket() *NumberFromBracketNumberBracket {
	return (&NumberFromBracketNumberBracket{})
}
