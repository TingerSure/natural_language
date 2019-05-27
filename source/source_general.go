package source

import (
	"strings"
)

type SourceGeneral struct {
	name  string
	words []string
}

func (s *SourceGeneral) SetName(name string) {
	s.name = name
}

func (s *SourceGeneral) GetName() string {
	return s.name
}

func (s *SourceGeneral) AddWord(word string) *SourceGeneral {
	s.words = append(s.words, word)
	return s
}

func (s *SourceGeneral) GetWords(firstCharacter string) []string {
	var words []string

	for _, word := range s.words {
		if strings.Index(word, firstCharacter) == 0 {
			words = append(words, word)
		}
	}
	return words
}

func NewSourceGeneral() *SourceGeneral {
	return &SourceGeneral{}
}
