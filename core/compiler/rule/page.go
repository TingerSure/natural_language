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
					err := page.SetImport(
						context.GetLibraryManager().Sandbox.Variable.String.New(importIndex.Name()),
						importIndex,
					)
					if !nl_interface.IsNil(err) {
						return nil, err
					}
					continue
				}
				exportIndex, yes := index.IndexFamilyInstance.IsExportIndex(item)
				if yes {
					err := page.SetExport(
						context.GetLibraryManager().Sandbox.Variable.String.New(exportIndex.Name()),
						exportIndex,
					)
					if !nl_interface.IsNil(err) {
						return nil, err
					}
					continue
				}
				varIndex, yes := index.IndexFamilyInstance.IsVarIndex(item)
				if yes {
					err := page.SetVar(
						context.GetLibraryManager().Sandbox.Variable.String.New(varIndex.Name()),
						varIndex,
					)
					if !nl_interface.IsNil(err) {
						return nil, err
					}
					continue
				}
				return nil, errors.New(fmt.Sprintf("Unsupport index to be set: %v", item.ToString("")))
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
		semantic.NewRule(PolymerizePageItemFromExportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPageItem -> SymbolExportGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemFromVarGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPageItem -> SymbolVarGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolImportGroup -> SymbolImport SymbolIdentifier SymbolString
			path := context.FormatSymbolString(phrase.GetChild(2).GetToken().Value())
			pageIndex, err := context.GetPage(path)
			if err != nil {
				return nil, err
			}
			page, exception := pageIndex.Get(nil)
			if !nl_interface.IsNil(exception) {
				return nil, errors.New(fmt.Sprintf(
					"Page index error: \"%v\"(\"%v\") is not an index without closure, cannot be used as a page index.",
					path,
					pageIndex.Type(),
				))
			}

			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ImportIndex.New(
					phrase.GetChild(1).GetToken().Value(),
					path,
					page,
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeExportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExportGroup -> SymbolExport SymbolIdentifier SymbolIndex
			name := phrase.GetChild(1).GetToken().Value()
			indexes, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ExportIndex.New(name, indexes[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizeVarGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVarGroup -> SymbolVar SymbolIdentifier SymbolIndex
			indexes, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.VarIndex.New(
					phrase.GetChild(1).GetToken().Value(),
					indexes[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeIndexFromIdentifier, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolIdentifier
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.BubbleIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(0).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeIndexFromNumber, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolNumber
			value, err := strconv.ParseFloat(phrase.GetChild(0).GetToken().Value(), 64)
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ConstIndex.New(
					context.GetLibraryManager().Sandbox.Variable.Number.New(value),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeIndexFromString, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolString
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ConstIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						context.FormatSymbolString(
							phrase.GetChild(0).GetToken().Value(),
						),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeIndexFromBool, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolBool
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeBoolFromTrue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolBool -> SymbolTrue
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ConstIndex.New(
					context.GetLibraryManager().Sandbox.Variable.Bool.New(true),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeBoolFromFalse, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolBool -> SymbolFalse
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ConstIndex.New(
					context.GetLibraryManager().Sandbox.Variable.Bool.New(false),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeIndexArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndexArray -> SymbolIndex
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndexArray -> SymbolIndexArray SymbolComma SymbolIndex
			items, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			item, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return append(items, item...), nil
		}),
		semantic.NewRule(PolymerizeKeyValue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolKeyValue -> SymbolIdentifier SymbolColon SymbolIndex
			indexes, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.KeyValueIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(0).GetToken().Value(),
					),
					indexes[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeKeyValueArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolKeyValueArray -> SymbolKeyValue
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyValueArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolKeyValueArray -> SymbolKeyValueArray SymbolComma SymbolKeyValue
			items, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			item, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return append(items, item...), nil
		}),
		semantic.NewRule(PolymerizeCallWithoutParam, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolIndex SymbolLeftParenthesis SymbolRightParenthesis
			indexes, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.Call.New(
					indexes[0],
					context.GetLibraryManager().Sandbox.Index.ConstIndex.New(
						context.GetLibraryManager().Sandbox.Variable.Param.New(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeCallWithIndexArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolIndex SymbolLeftParenthesis SymbolIndexArray SymbolRightParenthesis
			funcs, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			params, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			newParam := context.GetLibraryManager().Sandbox.Expression.NewParam.New()
			newParam.SetArray(params)
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.Call.New(
					funcs[0],
					newParam,
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeCallWithKeyValueArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolIndex SymbolLeftParenthesis SymbolKeyValueArray SymbolRightParenthesis
			funcs, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			params, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			newParam := context.GetLibraryManager().Sandbox.Expression.NewParam.New()
			newParam.SetKeyValue(params)
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.Call.New(
					funcs[0],
					newParam,
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeComponent, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolIndex SymbolDot SymbolIdentifier
			indexes, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.Component.New(
					indexes[0],
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(2).GetToken().Value(),
					),
				),
			}, nil
		}),
	}
)
