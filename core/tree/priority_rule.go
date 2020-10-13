package tree

type PriorityRuleParam struct {
	Match   func(Phrase, Phrase) bool
	Chooser func(Phrase, Phrase) int
	From    string
}

type PriorityRule struct {
	param *PriorityRuleParam
}

func (r *PriorityRule) GetFrom() string {
	return r.param.From
}

func (p *PriorityRule) Match(left, right Phrase) bool {
	return p.param.Match(left, right)
}

func (p *PriorityRule) Choose(left, right Phrase) int {
	return p.param.Chooser(left, right)
}

func NewPriorityRule(param *PriorityRuleParam) *PriorityRule {
	if param.Match == nil {
		panic("no match function in this priority rule!")
	}
	if param.Chooser == nil {
		panic("no chooser function in this priority rule!")
	}
	return &PriorityRule{
		param: param,
	}
}
