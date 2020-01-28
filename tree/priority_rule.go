package tree

type PriorityRuleParam struct {
	Match   func(Phrase, Phrase) bool
	Chooser func(Phrase, Phrase) int
}

type PriorityRule struct {
	match   func(Phrase, Phrase) bool
	chooser func(Phrase, Phrase) int
}

func (p *PriorityRule) Match(left, right Phrase) bool {
	return p.match(left, right)
}

func (p *PriorityRule) Choose(left, right Phrase) int {
	return p.chooser(left, right)
}

func NewPriorityRule(param *PriorityRuleParam) *PriorityRule {
	return &PriorityRule{
		match:   param.Match,
		chooser: param.Chooser,
	}
}
