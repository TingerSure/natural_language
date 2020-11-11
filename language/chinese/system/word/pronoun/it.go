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

type It struct {
	*adaptor.SourceAdaptor
	ItIndex   concept.Index
	instances []*tree.Vocabulary
}

func (p *It) GetName() string {
	return ItName
}

func (p *It) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
}

func (p *It) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Index {
						return p.ItIndex
					},
					Content: treasure,
					Types:   phrase_type.PronounPersonalName,
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
	instance.instances = []*tree.Vocabulary{
		tree.NewVocabulary(ItCharactor, instance),
	}
	return instance
}
