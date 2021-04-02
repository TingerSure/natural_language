package rule

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

var (
	SemanticRules = []*semantic.Rule{
		semantic.NewRule(PolymerizePageGroupEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolPageGroup -> SymbolPage SymbolIdentifier SymbolLeftBrace SymbolRightBrace
			page := semantic.NewFilePage(context.GetLibraryManager())
			page.SetName(phrase.GetChild(1).GetToken().Value())
			return []concept.Index{page}, nil
		}),
		semantic.NewRule(PolymerizePageGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolPageGroup -> SymbolPage SymbolIdentifier SymbolLeftBrace SymbolPageItemArray SymbolRightBrace
			page := semantic.NewFilePage(context.GetLibraryManager())
			page.SetName(phrase.GetChild(1).GetToken().Value())

			// items, err := context.Deal(phrase.GetChild(3))
			// if err != nil {
			// 	return nil, err
			// }
			// for _, item := range items {
			// 	item = nil
			// 	// TODO
			// }

			return []concept.Index{page}, nil
		}),
		semantic.NewRule(PolymerizePageItemArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolPageItemArray -> SymbolPageItem
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolPageItemArray -> SymbolPageItemArray SymbolPageItem
			items, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			item, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, err
			}
			return append(items, item...), nil
		}),
		semantic.NewRule(PolymerizePageItemFromImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPageItem -> SymbolImportGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolImportGroup -> SymbolImport SymbolIdentifier SymbolString SymbolSemicolon
			path := context.FormatSymbolString(phrase.GetChild(2).GetToken().Value())
			name := phrase.GetChild(1).GetToken().Value()
			page, err := context.GetPage(path)
			if err != nil {
				return nil, err
			}

			return []concept.Index{context.GetLibraryManager().Sandbox.Index.ImportIndex.New(name, path, page)}, nil
		}),
	}
)
