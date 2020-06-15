package pronoun

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	ItName string = "word.pronoun.it"

	ItCharactor string = "它"
)

var (
	itPronounWords []*tree.Word = []*tree.Word{
		tree.NewWord(ItCharactor),
	}
)

type It struct {
	*adaptor.SourceAdaptor
	ItIndex concept.Index
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
						return p.ItIndex
					},
					Content: treasure,
					Types:   phrase_type.Any,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewIt(param *adaptor.SourceAdaptorParam) *It {
	instance := (&It{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	page := instance.Libs.GetLibraryPage("system", "pronoun")
	instance.ItIndex = page.GetIndex(instance.Libs.Sandbox.Variable.String.New("It"))
	return instance
}
