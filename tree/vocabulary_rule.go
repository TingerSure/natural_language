package tree

type VocabularyRule struct {
	create func(treasure *Vocabulary) Phrase
	match  func(treasure *Vocabulary) bool
	from   string
}

func (r *VocabularyRule) GetFrom() string {
	return r.from
}

func (r *VocabularyRule) Match(treasure *Vocabulary) bool {
	return r.match(treasure)
}

func (r *VocabularyRule) Create(treasure *Vocabulary) Phrase {
	return r.create(treasure)
}

func NewVocabularyRule(
	match func(treasure *Vocabulary) bool,
	create func(treasure *Vocabulary) Phrase,
	from string,
) *VocabularyRule {
	if create == nil {
		panic("no create function in this vocabulary rule!")
	}
	return &VocabularyRule{
		create: create,
		match:  match,
		from:   from,
	}
}
