package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type PriorityRuleParam struct {
	Match   func(Phrase, Phrase) (bool, concept.Exception)
	Chooser func(Phrase, Phrase) (*PriorityResult, concept.Exception)
	From    string
}

type PriorityRule struct {
	param *PriorityRuleParam
}

func (r *PriorityRule) GetFrom() string {
	return r.param.From
}

func (p *PriorityRule) Match(left, right Phrase) (bool, concept.Exception) {
	return p.param.Match(left, right)
}

func (p *PriorityRule) Choose(left, right Phrase) (*PriorityResult, concept.Exception) {
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
