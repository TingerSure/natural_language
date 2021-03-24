package semantic

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
)

type Rule struct {
	source *grammar.Rule
	deal   func(grammar.Phrase, *Context, *FilePage) error
}

func NewRule(source *grammar.Rule, deal func(grammar.Phrase, *Context, *FilePage) error) *Rule {
	return &Rule{
		source: source,
		deal:   deal,
	}
}

func (r *Rule) GetSource() *grammar.Rule {
	return r.source
}

func (r *Rule) Deal(phrase grammar.Phrase, context *Context, page *FilePage) error {
	return r.deal(phrase, context, page)
}
