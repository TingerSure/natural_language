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
	ResultName      string = "word.pronoun.result"
	ResultType      int    = word_types.Pronoun
	ResultCharactor string = "结果"
)

var (
	resultPronounWords []*tree.Word = []*tree.Word{
		tree.NewWord(ResultCharactor, ResultType),
	}
)

type Result struct {
	adaptor.SourceAdaptor
}

func (p *Result) GetName() string {
	return ResultName
}

func (p *Result) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(resultPronounWords, sentence)
}

func (p *Result) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewResaultIndex([]concept.Matcher{
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

func NewResult() *Result {
	return (&Result{})
}
