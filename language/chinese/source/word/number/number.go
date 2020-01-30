package number

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/core/tree/phrase_types"
	"github.com/TingerSure/natural_language/core/tree/word_types"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"regexp"
	"strconv"
)

const (
	NumberName string = "word.number"
	numberType int    = word_types.Number
)

var (
	numberTemplate *regexp.Regexp = regexp.MustCompile("^(-?\\d+)(\\.\\d+)?")
)

type Number struct {
	adaptor.SourceAdaptor
}

func (p *Number) GetName() string {
	return NumberName
}

func (p *Number) GetWords(sentence string) []*tree.Word {
	value := numberTemplate.FindString(sentence)
	if value != "" {
		return []*tree.Word{tree.NewWord(value, numberType)}
	}
	return nil
}

func (p *Number) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						value, err := strconv.ParseFloat(treasure.GetWord().GetContext(), 64)
						if err != nil {
							panic(err)
						}
						return index.NewConstIndex(variable.NewNumber(value))
					},
					Content: treasure,
					Types:   phrase_types.Number,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewNumber() *Number {
	return (&Number{})
}
