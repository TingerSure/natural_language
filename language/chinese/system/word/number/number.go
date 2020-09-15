package number

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
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
	*adaptor.SourceAdaptor
}

func (p *Number) GetName() string {
	return NumberName
}

func (p *Number) GetWords(sentence string) []*tree.Vocabulary {
	value := numberTemplate.FindString(sentence)
	if value != "" {
		return []*tree.Vocabulary{tree.NewVocabulary(value, p)}
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
						value, err := strconv.ParseFloat(treasure.GetContext(), 64)
						if err != nil {
							panic(err)
						}
						return p.Libs.Sandbox.Index.ConstIndex.New(p.Libs.Sandbox.Variable.Number.New(value))
					},
					Content: treasure,
					Types:   phrase_type.Number,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewNumber(param *adaptor.SourceAdaptorParam) *Number {
	return (&Number{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
