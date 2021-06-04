package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Diversion struct {
	rules     []*tree.DutyRule
	rootSpace concept.Pool
}

func NewDiversion(rootSpace concept.Pool) *Diversion {
	return &Diversion{
		rootSpace: rootSpace,
	}
}

func (d *Diversion) Match(value tree.Phrase) (string, error) {
	wanted := value.Index().Anticipate(d.rootSpace)
	for _, rule := range d.rules {
		if rule.Match(wanted) {
			return rule.Create(wanted), nil
		}
	}
	return "", fmt.Errorf("This Phrase has no DutyRule to match!\nphrase : %v\nanticipate : %v", value.ToString(), wanted.ToString(""))
}

func (a *Diversion) AddRule(rule *tree.DutyRule) {
	if rule == nil {
		return
	}
	a.rules = append(a.rules, rule)
}

func (a *Diversion) RemoveRule(need func(rule *tree.DutyRule) bool) {
	for index := 0; index < len(a.rules); index++ {
		rule := a.rules[index]
		if need(rule) {
			a.rules = append(a.rules[:index], a.rules[index+1:]...)
		}
	}
}
