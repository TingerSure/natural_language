package tree

type Source interface {
	GetName() string
	GetWords(firstCharacter string) []*Word
	GetVocabularyRules() []*VocabularyRule
	GetStructRules() []*StructRule
}
