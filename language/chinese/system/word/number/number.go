package number

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"

	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"regexp"
	"strconv"
)

const (
	NumberName string = "word.number"
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
		return []*tree.Word{tree.NewWord(value)}
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
					Types:   phrase_type.Number,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewNumber(libs *tree.LibraryManager) *Number {
	return (&Number{})
}
