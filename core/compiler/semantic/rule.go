package semantic

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
)

type Rule struct {
	source *grammar.Rule
	deal   func(grammar.Phrase, *grammar.Rule, *FilePage) error
}

func NewRule(source *grammar.Rule, deal func(grammar.Phrase, *grammar.Rule, *FilePage) error) *Rule {
	return &Rule{
		source: source,
		deal:   deal,
	}
}

func (r *Rule) GetSource() *grammar.Rule {
	return r.source
}

func (r *Rule) Deal(phrase grammar.Phrase, page *FilePage) error {
	return r.deal(phrase, r.source, page)
}
