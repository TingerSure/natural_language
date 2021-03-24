package rule

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
)

var (
	SemanticRules = []*semantic.Rule{
		semantic.NewRule(PolymerizePageGroupEmpty, func(phrase grammar.Phrase, context *semantic.Context, page *semantic.FilePage) error {
			// SymbolPageGroup -> SymbolPage SymbolIdentifier SymbolLeftBrace SymbolRightBrace
			page.SetName(phrase.GetChild(1).GetToken().Value())
			return nil
		}),
		semantic.NewRule(PolymerizePageGroup, func(phrase grammar.Phrase, context *semantic.Context, page *semantic.FilePage) error {
			// SymbolPageGroup -> SymbolPage SymbolIdentifier SymbolLeftBrace SymbolPageItemArray SymbolRightBrace
			page.SetName(phrase.GetChild(1).GetToken().Value())
			return context.Deal(phrase.GetChild(3), context, page)
		}),
		semantic.NewRule(PolymerizePageItemArrayStart, func(phrase grammar.Phrase, context *semantic.Context, page *semantic.FilePage) error {
			// SymbolPageItemArray -> SymbolPageItem
			return context.Deal(phrase.GetChild(0), context, page)
		}),
		semantic.NewRule(PolymerizePageItemArray, func(phrase grammar.Phrase, context *semantic.Context, page *semantic.FilePage) error {
			// SymbolPageItemArray -> SymbolPageItemArray SymbolPageItem
			err := context.Deal(phrase.GetChild(0), context, page)
			if err != nil {
				return err
			}
			return context.Deal(phrase.GetChild(1), context, page)
		}),
		semantic.NewRule(PolymerizePageItemFromImportGroup, func(phrase grammar.Phrase, context *semantic.Context, page *semantic.FilePage) error {
			//SymbolPageItem -> SymbolImportGroup
			return context.Deal(phrase.GetChild(0), context, page)
		}),
		semantic.NewRule(PolymerizePageImportGroup, func(phrase grammar.Phrase, context *semantic.Context, page *semantic.FilePage) error {
			//SymbolImportGroup -> SymbolImport SymbolIdentifier SymbolString SymbolSemicolon
			importGroup, err := context.GetImport(phrase.GetChild(2).GetToken().Value())
			if err != nil {
				return err
			}
			page.AddImport(phrase.GetChild(1).GetToken().Value(), importGroup)
			return nil
		}),
	}
)
