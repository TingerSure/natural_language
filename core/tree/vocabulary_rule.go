package tree

type VocabularyRuleParam struct {
	Create func(treasure *Vocabulary) Phrase
	Match  func(treasure *Vocabulary) bool
	From   string
}

type VocabularyRule struct {
	param *VocabularyRuleParam
}

func (r *VocabularyRule) GetFrom() string {
	return r.param.From
}

func (r *VocabularyRule) Match(treasure *Vocabulary) bool {
	return r.param.Match(treasure)
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
