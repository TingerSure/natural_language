package lexer

import (
	"github.com/TingerSure/natural_language/tree"
)

type LexerInstance struct {
	sentence     string
	vocabularies []*tree.Vocabulary
	index        int
}

func (l *LexerInstance) ToString() string {
	var toString string = ""
	l.Reset()
	for vocabulary := l.Next(); vocabulary != nil; vocabulary = l.Next() {
        toString += vocabulary.GetWord()
		toString += " "
        // toString += " ( "
        // if vocabulary.GetSource()!= nil {
        //     toString += vocabulary.GetSource().GetName()
        // }else{
        //     toString += "nil"
        // }
		// toString += " ) "

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

func (l *LexerInstance) GetSentence(sentence string) string {
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
	l.vocabularies = append(l.vocabularies, vocabulary)
}

func (l *LexerInstance) init() *LexerInstance {
	return l
}

func NewLexerInstance() *LexerInstance {
	return (&LexerInstance{}).init()
}
