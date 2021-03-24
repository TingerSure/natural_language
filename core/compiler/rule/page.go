package rule

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
)

var (
	SemanticRules = []*semantic.Rule{
		semantic.NewRule(PolymerizePageGroupEmpty, func(phrase grammar.Phrase, rule *grammar.Rule, page *semantic.FilePage) error {
			//TODO
			return nil
		}),
		semantic.NewRule(PolymerizePageGroup, func(phrase grammar.Phrase, rule *grammar.Rule, page *semantic.FilePage) error {
			//TODO
			return nil
		}),
		semantic.NewRule(PolymerizePageItemArrayStart, func(phrase grammar.Phrase, rule *grammar.Rule, page *semantic.FilePage) error {
			//TODO
			return nil
		}),
		semantic.NewRule(PolymerizePageItemArray, func(phrase grammar.Phrase, rule *grammar.Rule, page *semantic.FilePage) error {
			//TODO
			return nil
		}),
		semantic.NewRule(PolymerizePageItemFromImportGroup, func(phrase grammar.Phrase, rule *grammar.Rule, page *semantic.FilePage) error {
			//TODO
			return nil
		}),
		semantic.NewRule(PolymerizePageImportGroup, func(phrase grammar.Phrase, rule *grammar.Rule, page *semantic.FilePage) error {
			//TODO
			return nil
		}),
	}
)
