package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/creator"
	"github.com/TingerSure/natural_language/core/tree"
)

type Diversion struct {
	rules     []*tree.DutyRule
	rootSpace concept.Pool
	sandbox   *creator.SandboxCreator
}

func NewDiversion(rootSpace concept.Pool, sandbox *creator.SandboxCreator) *Diversion {
	return &Diversion{
		rootSpace: rootSpace,
		sandbox:   sandbox,
	}
}

func (d *Diversion) Match(value tree.Phrase) (string, error) {
	line := tree.NewLine("[diversion_match]", "")
	pipe, exception := value.Index()
	if !nl_interface.IsNil(exception) {
		exception.AddExceptionLine(line)
		return "", exception
	}
	param, exception := pipe.Exec(d.sandbox.Variable.Param.New(), nil)
	if !nl_interface.IsNil(exception) {
		exception.AddExceptionLine(line)
		return "", exception
	}
	wanted := param.Get(d.sandbox.Variable.String.New("value"))
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
