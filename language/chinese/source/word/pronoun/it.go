package pronoun

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/matcher"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	ItName      string = "word.pronoun.it"
	ItType      int    = word_types.Pronoun
	ItCharactor string = "它"
)

var (
	itPronounWords []*tree.Word = []*tree.Word{
		tree.NewWord(ItCharactor, ItType),
	}
)

type It struct {
	adaptor.SourceAdaptor
}

func (p *It) GetName() string {
	return ItName
}

func (p *It) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(itPronounWords, sentence)
}

func (p *It) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewSearchIndex([]concept.Matcher{
							matcher.NewSystemMatcher(func(concept.Variable) bool {
								return true
							}),
						})
					},
					Content: treasure,
					Types:   phrase_types.Any,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewIt() *It {
	return (&It{})
}
