package lexer

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/tree"
)

type Flow struct {
	sentence     string
	vocabularies []*tree.Vocabulary
	index        int
}

func (l *Flow) Len() int {
	return len(l.vocabularies)
}

func (l *Flow) ValidLength() int {
	var valid int = 0
	for _, vocabulary := range l.vocabularies {
		if nl_interface.IsNil(vocabulary.GetSource()) {
			return valid
		}
		valid += vocabulary.Len()
	}
	return valid
}

func (l *Flow) HasNull() bool {
	for _, vocabulary := range l.vocabularies {
		if nl_interface.IsNil(vocabulary.GetSource()) {
			return true
		}
	}
	return false
}

func (l *Flow) ToString() string {
	var toString string = ""
	l.Reset()
	for vocabulary := l.Next(); vocabulary != nil; vocabulary = l.Next() {
		toString += vocabulary.GetContext()
		toString += "("
		if vocabulary.GetSource() != nil {
			toString += vocabulary.GetSource().GetName()
		} else {
			toString += "nil"
		}
		toString += ")"
		toString += " "
	}
	return toString
}

func (l *Flow) Index(index int) *tree.Vocabulary {
	if index < 0 || index >= l.Len() {
		return nil
	}
	return l.vocabularies[index]
}

func (l *Flow) Next() *tree.Vocabulary {
	if l.IsEnd() {
		return nil
	}
	now := l.vocabularies[l.index]
	l.index++
	return now
}

func (l *Flow) Peek() *tree.Vocabulary {
	if l.IsEnd() {
		return nil
	}
	return l.vocabularies[l.index]
}

func (l *Flow) SetSentence(sentence string) {
	l.sentence = sentence
}

func (l *Flow) GetSentence() string {
	return l.sentence
}

func (l *Flow) Copy() *Flow {
	newInstance := NewFlow()
	newInstance.sentence = l.sentence
	newInstance.vocabularies = make([]*tree.Vocabulary, len(l.vocabularies))
	copy(newInstance.vocabularies, l.vocabularies)
	newInstance.index = l.index
	return newInstance
}

func (l *Flow) IsEnd() bool {
	return (l.index >= len(l.vocabularies))
}

func (l *Flow) Reset() {
	l.index = 0
}

func (l *Flow) AddVocabulary(vocabulary *tree.Vocabulary) {
	if len(l.vocabularies) != 0 {
		var last *tree.Vocabulary = l.vocabularies[len(l.vocabularies)-1]
		if last.GetSource() == nil && vocabulary.GetSource() == nil {
			l.vocabularies[len(l.vocabularies)-1] = tree.NewVocabulary(last.GetContext()+vocabulary.GetContext(), nil)
			return
		}
	}
	l.vocabularies = append(l.vocabularies, vocabulary)
}

func (l *Flow) init() *Flow {
	return l
}

func NewFlow() *Flow {
	return (&Flow{}).init()
}
