package rule

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"strconv"
)

var (
	SemanticRules = []*semantic.Rule{

		semantic.NewRule(PolymerizePageGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolPageGroup -> SymbolPageItemArray
			page := context.GetLibraryManager().Sandbox.Variable.Page.New()

			items, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			for _, item := range items {
				importIndex, yes := index.IndexFamilyInstance.IsImportIndex(item)
				if yes {
					exception := page.SetImport(context.GetLibraryManager().Sandbox.Variable.String.New(importIndex.Name()), importIndex)
					if nl_interface.IsNil(exception) {
						return nil, exception
					}
					continue
				}
				exportIndex, yes := index.IndexFamilyInstance.IsExportIndex(item)
				if yes {
					exception := page.SetExport(context.GetLibraryManager().Sandbox.Variable.String.New(exportIndex.Name()), exportIndex)
					if nl_interface.IsNil(exception) {
						return nil, exception
					}
					continue
				}
				// TODO varIndex
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ConstIndex.New(page),
			}, nil
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
		semantic.NewRule(PolymerizeImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolImportGroup -> SymbolImport SymbolIdentifier SymbolString
			path := context.FormatSymbolString(phrase.GetChild(2).GetToken().Value())
			name := phrase.GetChild(1).GetToken().Value()
			pageIndex, err := context.GetPage(path)
			if err != nil {
				return nil, err
			}
			page, exception := pageIndex.Get(nil)
			if !nl_interface.IsNil(exception) {
				return nil, errors.New(fmt.Sprintf("Page index error: \"%v\"(\"%v\") is not an index without closure, cannot be used as a page index.", path, pageIndex.Type()))
			}

			return []concept.Index{context.GetLibraryManager().Sandbox.Index.ImportIndex.New(name, path, page)}, nil
		}),
		semantic.NewRule(PolymerizeExportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExportGroup -> SymbolExport SymbolIdentifier SymbolIndex
			name := phrase.GetChild(1).GetToken().Value()
			indexes, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Index{context.GetLibraryManager().Sandbox.Index.ExportIndex.New(name, indexes[0])}, nil
		}),
		semantic.NewRule(PolymerizeIndexFromNumber, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolNumber
			value, err := strconv.ParseFloat(phrase.GetChild(0).GetToken().Value(), 64)
			if err != nil {
				return nil, err
			}
			return []concept.Index{context.GetLibraryManager().Sandbox.Index.ConstIndex.New(context.GetLibraryManager().Sandbox.Variable.Number.New(value))}, nil
		}),
		semantic.NewRule(PolymerizeIndexFromString, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolString
			return []concept.Index{context.GetLibraryManager().Sandbox.Index.ConstIndex.New(context.GetLibraryManager().Sandbox.Variable.String.New(context.FormatSymbolString(phrase.GetChild(0).GetToken().Value())))}, nil
		}),
	}
)
