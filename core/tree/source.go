package tree

type Source interface {
	GetName() string
	GetWords(string) []*Vocabulary
	GetVocabularyRules() []*VocabularyRule
	GetStructRules() []*StructRule
	GetPriorityRules() []*PriorityRule
	GetDutyRules() []*DutyRule
	GetPhraseTypes() []*PhraseType
}
