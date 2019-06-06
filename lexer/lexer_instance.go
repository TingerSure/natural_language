package lexer

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/word"
)

type LexerInstance struct {
	sentence     string
	vocabularies []*tree.Vocabulary
	index        int
}

func (l *LexerInstance) ValidLength() int {
	var valid int = 0
	for _, vocabulary := range l.vocabularies {
		if vocabulary.GetSource() == nil {
			return valid
		}
		valid += vocabulary.GetWord().Len()
	}
	return valid
}

func (l *LexerInstance) HasNull() bool {
	for _, vocabulary := range l.vocabularies {
		if vocabulary.GetSource() == nil {
			return true
		}
	}
	return false
}

func (l *LexerInstance) ToString() string {
	var toString string = ""
	l.Reset()
	for vocabulary := l.Next(); vocabulary != nil; vocabulary = l.Next() {
		toString += vocabulary.GetWord().GetContext()
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

func (l *LexerInstance) Next() *tree.Vocabulary {
	if l.IsEnd() {
		return nil
	}
	now := l.vocabularies[l.index]
	l.index++
	return now
}

func (l *LexerInstance) SetSentence(sentence string) {
	l.sentence = sentence
}

func (l *LexerInstance) GetSentence() string {
	return l.sentence
}

func (l *LexerInstance) Copy() *LexerInstance {
	newInstance := NewLexerInstance()
	newInstance.sentence = l.sentence
	for i := 0; i < len(l.vocabularies); i++ {
		newInstance.vocabularies = append(newInstance.vocabularies, l.vocabularies[i])
	}
	return newInstance
}

func (l *LexerInstance) IsEnd() bool {
	return (l.index >= len(l.vocabularies))
}

func (l *LexerInstance) Reset() {
	l.index = 0
}

func (l *LexerInstance) AddVocabulary(vocabulary *tree.Vocabulary) {
	if len(l.vocabularies) != 0 {
		var last *tree.Vocabulary = l.vocabularies[len(l.vocabularies)-1]
		if last.GetSource() == nil && vocabulary.GetSource() == nil {
			l.vocabularies[len(l.vocabularies)-1] = tree.NewVocabulary(word.NewUnknownWord(last.GetWord().GetContext()+vocabulary.GetWord().GetContext()), nil)
			return
		}
	}
	l.vocabularies = append(l.vocabularies, vocabulary)
}

func (l *LexerInstance) init() *LexerInstance {
	return l
}

func NewLexerInstance() *LexerInstance {
	return (&LexerInstance{}).init()
}
