package tree

type VocabularyRule struct {
	logic func(treasure *Vocabulary) *Phrase
	from  string
}

func (r *VocabularyRule) GetFrom() string {
	return r.from
}

func (r *VocabularyRule) Logic(treasure *Vocabulary) *Phrase {
	return r.logic(treasure)
}

func NewVocabularyRule(
	logic func(treasure *Vocabulary) *Phrase,
	from string,
) *VocabularyRule {
	if logic == nil {
		panic("no logic function in this vocabulary rule!")
	}
	return &VocabularyRule{
		logic: logic,
		from:  from,
	}
}
