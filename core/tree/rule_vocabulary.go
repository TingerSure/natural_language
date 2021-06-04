package tree

type VocabularyRuleParam struct {
	Create func(*Vocabulary) Phrase
	Words  []*Vocabulary
	Match  string
	From   string
}

type VocabularyRule struct {
	param *VocabularyRuleParam
}

func (r *VocabularyRule) GetFrom() string {
	return r.param.From
}

func (r *VocabularyRule) Match() string {
	return r.param.Match
}

func (r *VocabularyRule) Words() []*Vocabulary {
	return r.param.Words
}

func (r *VocabularyRule) Create(treasure *Vocabulary) Phrase {
	return r.param.Create(treasure)
}

func NewVocabularyRule(param *VocabularyRuleParam) *VocabularyRule {
	if param.Create == nil {
		panic("no create function in this vocabulary rule!")
	}
	return &VocabularyRule{
		param: param,
	}
}
