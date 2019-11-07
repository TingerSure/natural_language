package tree

type Source interface {
	GetName() string
	GetWords(string) []*Word
	GetVocabularyRules() []*VocabularyRule
	GetStructRules() []*StructRule
}
