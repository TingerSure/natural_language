package brackets

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	RightCharactor        = ")"
	RightName      string = "word.brackets.right"
)

type BracketsRight struct {
	*adaptor.SourceAdaptor
	RightIndex concept.String
	instances  []*tree.Vocabulary
}

func (s *BracketsRight) GetName() string {
	return RightName
}

func (s *BracketsRight) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(s.instances, sentence)
}
func (s *BracketsRight) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == s
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return s.Libs.Sandbox.Index.ConstIndex.New(s.RightIndex.Clone())
					},
					Content: treasure,
					Types:   phrase_type.BracketsRight,
					From:    s.GetName(),
				})
			}, From: s.GetName(),
		}),
	}
}

func NewBracketsRight(param *adaptor.SourceAdaptorParam) *BracketsRight {
	right := (&BracketsRight{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
		RightIndex:    param.Libs.Sandbox.Variable.String.New(RightCharactor),
	})

	right.instances = []*tree.Vocabulary{
		tree.NewVocabulary(RightCharactor, right),
	}
	return right
}
