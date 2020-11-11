package parser

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/tree"
)

type Diversion struct {
	rules     []*tree.DutyRule
	rootSpace *closure.Closure
}

func NewDiversion(rootSpace *closure.Closure) *Diversion {
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
	return "", errors.New(fmt.Sprintf("This Phrase has no DutyRule to match!\nphrase : %v\nanticipate : %v", value, wanted))
}

func (a *Diversion) AddRule(rules []*tree.DutyRule) {
	if rules == nil {
		return
	}
	a.rules = append(a.rules, rules...)
}

func (a *Diversion) RemoveRule(need func(rule *tree.DutyRule) bool) {
	for index := 0; index < len(a.rules); index++ {
		rule := a.rules[index]
		if need(rule) {
			a.rules = append(a.rules[:index], a.rules[index+1:]...)
		}
	}
}
