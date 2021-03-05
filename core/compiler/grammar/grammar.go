package grammar

import (
	"errors"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

const (
	VirtualGlobalTypes = -1
)

type Grammar struct {
	rules  []*Rule
	global *Nonterminal
}

func NewGrammar() *Grammar {
	return &Grammar{}
}

func (g *Grammar) AddRule(rule *Rule) {
	g.rules = append(g.rules, rule)
}

func (g *Grammar) SetGlobal(global *Nonterminal) {
	g.global = global
}

func (g *Grammar) Build() error {
	err := g.format()
	if err != nil {
		return err
	}
	// TODO create goto/action table
	return nil
}

func (g *Grammar) format() error {
	err := g.formatGlobal()
	if err != nil {
		return err
	}
	// TODO format more
	return nil
}

func (g *Grammar) formatGlobal() error {
	if g.global == nil {
		return errors.New("Global missed.")
	}
	resultCount := 0
	fromCount := 0
	for _, rule := range g.rules {
		if g.global.Equal(rule.GetResult()) {
			resultCount++
		}
		for index := 0; index < rule.Size(); index++ {
			if g.global.Equal(rule.GetFrom(index)) {
				fromCount++
			}
		}
	}
	if resultCount < 1 {
		return errors.New("Rule missed which result to global")
	}
	if resultCount > 1 || fromCount > 0 {
		virtualGlobal := NewNonterminal(VirtualGlobalTypes)
		virtualRule := NewRule(virtualGlobal, g.global)
		g.rules = append(g.rules, virtualRule)
		g.global = virtualGlobal
	}
	return nil
}

func (g *Grammar) Read([]*lexer.Token) (tree.Page, error) {
	return nil, nil
}
