package rule

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	numberFromNumberOperatorNumberName string = "rule.number.number_number"
)

var (
	numberFromNumberOperatorNumberList []string = []string{
		phrase_types.Number,
		phrase_types.Operator,
		phrase_types.Number,
	}
)

type NumberFromNumberOperatorNumber struct {
}

func (p *NumberFromNumberOperatorNumber) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(func(phrase []tree.Phrase) concept.Index {

				return expression.NewAddition(phrase[0].Index(), phrase[2].Index())

			}, len(numberFromNumberOperatorNumberList), phrase_types.Number)
		}, numberFromNumberOperatorNumberList, p.GetName()),
	}
}

func (p *NumberFromNumberOperatorNumber) GetName() string {
	return numberFromNumberOperatorNumberName
}

func (p *NumberFromNumberOperatorNumber) GetWords(firstCharacter string) []*tree.Word {
	return nil
}

func (p *NumberFromNumberOperatorNumber) GetVocabularyRules() []*tree.VocabularyRule {
	return nil
}

func NewNumberFromNumberOperatorNumber() *NumberFromNumberOperatorNumber {
	return (&NumberFromNumberOperatorNumber{})
}
