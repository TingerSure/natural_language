package set

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	Is    = "是"
	Equal = "等于"

	SetName string = "word.verb.set"
)

var (
	SetCharactors = []*tree.Word{
		tree.NewWord(Is),
		tree.NewWord(Equal),
	}
)

type Set struct {
	*adaptor.SourceAdaptor
	Is    concept.String
	Equal concept.String
}

func (s *Set) GetName() string {
	return SetName
}

func (s *Set) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(SetCharactors, sentence)
}
func (s *Set) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == s
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				set := s.Is
				if treasure.GetWord().GetContext() == Equal {
					set = s.Equal
				}

				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return s.Libs.Sandbox.Index.ConstIndex.New(set)
					},
					Content: treasure,
					Types:   phrase_type.Set,
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
	set.Is.SetLanguage(param.Language, Is)
	set.Equal.SetLanguage(param.Language, Equal)
	return set
}
