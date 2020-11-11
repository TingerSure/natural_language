package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type DutyRuleParam struct {
	Create func(concept.Variable) string
	Match  func(concept.Variable) bool
	From   string
}

type DutyRule struct {
	param *DutyRuleParam
}

func (r *DutyRule) GetFrom() string {
	return r.param.From
}

func (r *DutyRule) Match(treasure concept.Variable) bool {
	return r.param.Match(treasure)
}

func (r *DutyRule) Create(treasure concept.Variable) string {
	return r.param.Create(treasure)
}

func NewDutyRule(param *DutyRuleParam) *DutyRule {
	if param.Match == nil {
		panic("no match function in this duty rule!")
	}
	if param.Create == nil {
		panic("no create function in this duty rule!")
	}
	return &DutyRule{
		param: param,
	}
}
