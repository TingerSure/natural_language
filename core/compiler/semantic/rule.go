package semantic

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Rule struct {
	source *grammar.Rule
	deal   func(grammar.Phrase, *Context) ([]concept.Pipe, error)
}

func NewRule(source *grammar.Rule, deal func(grammar.Phrase, *Context) ([]concept.Pipe, error)) *Rule {
	return &Rule{
		source: source,
		deal:   deal,
	}
}

func (r *Rule) GetSource() *grammar.Rule {
	return r.source
}

func (r *Rule) Deal(phrase grammar.Phrase, context *Context) ([]concept.Pipe, error) {
	return r.deal(phrase, context)
}
