package set

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	IsCharactor    = "是"
	EqualCharactor = "等于"

	SetName string = "word.verb.set"
)

type Set struct {
	*adaptor.SourceAdaptor
	Is        concept.String
	Equal     concept.String
	instances []*tree.Vocabulary
}

func (s *Set) GetName() string {
	return SetName
}

func (s *Set) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(s.instances, sentence)
}
func (s *Set) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == s
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				set := s.Is
				if treasure.GetContext() == EqualCharactor {
					set = s.Equal
				}

				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Index {
						return s.Libs.Sandbox.Index.ConstIndex.New(set)
					},
					Content: treasure,
					Types:   phrase_type.SetName,
					From:    s.GetName(),
				})
			}, From: s.GetName(),
		}),
	}
}

func NewSet(param *adaptor.SourceAdaptorParam) *Set {
	set := (&Set{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})

	setPage := set.Libs.GetLibraryPage("system", "set")

	set.Is = setPage.GetConst(set.Libs.Sandbox.Variable.String.New("Is"))
	set.Equal = setPage.GetConst(set.Libs.Sandbox.Variable.String.New("Equal"))
	set.Is.SetLanguage(param.Language, IsCharactor)
	set.Equal.SetLanguage(param.Language, EqualCharactor)
	set.instances = []*tree.Vocabulary{
		tree.NewVocabulary(EqualCharactor, set),
		tree.NewVocabulary(IsCharactor, set),
	}
	return set
}
