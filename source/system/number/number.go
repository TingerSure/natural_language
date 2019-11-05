package number

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/sandbox/variable"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
	"strconv"
)

const (
	numberName string = "system.number"
	numberType int    = word_types.Number
)

var (
	NumberCharactors = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	numberWords []*tree.Word = func() []*tree.Word {
		words := []*tree.Word{}
		for _, character := range NumberCharactors {
			words = append(words, tree.NewWord(character, numberType))
		}
		return words
	}()
)

type Number struct {
}

func (p *Number) GetName() string {
	return numberName
}

func (p *Number) GetWords(firstCharacter string) []*tree.Word {
	return tree.WordsFilter(numberWords, firstCharacter)
}

func (p *Number) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				value, err := strconv.ParseFloat(treasure.GetWord().GetContext(), 64)
				if err != nil {
					panic(err)
				}
				return index.NewConstIndex(variable.NewNumber(value))
			}, treasure, phrase_types.Number)
		}, p.GetName()),
	}
}

func (p *Number) GetStructRules() []*tree.StructRule {
	return nil
}

func NewNumber() *Number {
	return (&Number{})
}
