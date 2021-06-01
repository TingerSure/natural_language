package brackets

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	LeftCharactor        = "("
	LeftName      string = "word.brackets.left"
)

type BracketsLeft struct {
	*adaptor.SourceAdaptor
	LeftIndex concept.String
	instances []*tree.Vocabulary
}

func (s *BracketsLeft) GetName() string {
	return LeftName
}

func (s *BracketsLeft) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(s.instances, sentence)
}
func (s *BracketsLeft) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == s
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Pipe {
						return s.Libs.Sandbox.Index.ConstIndex.New(s.LeftIndex.Clone())
					},
					Content: treasure,
					Types:   phrase_type.BracketsLeftName,
					From:    s.GetName(),
				})
			}, From: s.GetName(),
		}),
	}
}

func NewBracketsLeft(param *adaptor.SourceAdaptorParam) *BracketsLeft {
	left := (&BracketsLeft{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
		LeftIndex:     param.Libs.Sandbox.Variable.String.New(LeftCharactor),
	})

	left.instances = []*tree.Vocabulary{
		tree.NewVocabulary(LeftCharactor, left),
	}
	return left
}
