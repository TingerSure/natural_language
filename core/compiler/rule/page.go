package rule

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"strconv"
)

var (
	SemanticRules = []*semantic.Rule{

		semantic.NewRule(PolymerizePageGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolPageGroup -> SymbolPageItemList
			page := context.GetLibraryManager().Sandbox.Variable.Page.New()

			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			for cursor, item := range items {
				importIndex, yes := index.IndexFamilyInstance.IsImportIndex(item)
				if yes {
					exception := page.SetImport(
						context.GetLibraryManager().Sandbox.Variable.String.New(importIndex.Name()),
						importIndex,
					)
					if !nl_interface.IsNil(exception) {
						return nil, nil, exception.AddExceptionLine(lines[cursor])
					}
					continue
				}
				publicIndex, yes := index.IndexFamilyInstance.IsPublicIndex(item)
				if yes {
					exception := page.SetPublic(
						context.GetLibraryManager().Sandbox.Variable.String.New(publicIndex.Name()),
						publicIndex,
					)
					if !nl_interface.IsNil(exception) {
						return nil, nil, exception.AddExceptionLine(lines[cursor])
					}
					continue
				}
				privateIndex, yes := index.IndexFamilyInstance.IsPrivateIndex(item)
				if yes {
					exception := page.SetPrivate(
						context.GetLibraryManager().Sandbox.Variable.String.New(privateIndex.Name()),
						privateIndex,
					)
					if !nl_interface.IsNil(exception) {
						return nil, nil, exception.AddExceptionLine(lines[cursor])
					}
					continue
				}
				return nil, nil, fmt.Errorf("Unsupported pipe to be set to page: %v\n%v", item.ToString(""), lines[cursor].ToString())
			}
			var line *lexer.Line

			if len(lines) == 0 {
				line = lexer.NewLine(context.Path(), context.Content())
			} else {
				line = lexer.NewLineList(lines)
			}

			instance := context.GetLibraryManager().Sandbox.Index.ConstIndex.New(page)
			instance.SetLine(line)
			return []concept.Pipe{instance}, []*lexer.Line{line}, nil
		}),
		semantic.NewRule(PolymerizePageItemList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolPageItemList -> SymbolPageItemArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolPageItemList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizePageItemArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolPageItemArray -> SymbolPageItem
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolPageItemArray -> SymbolPageItemArray SymbolPageItem
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizeClassGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolClassGroup -> SymbolClass SymbolLeftBrace SymbolClassItemList SymbolRightBrace
			items, prelines, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			newClass := context.GetLibraryManager().Sandbox.Expression.NewClass.New()
			newClass.SetItems(items)
			lines := make([]concept.Line, len(prelines))
			for cursor, line := range prelines {
				lines[cursor] = line
			}
			newClass.SetLines(lines)
			return []concept.Pipe{newClass}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(3).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeClassItemList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolClassItemList -> SymbolClassItemArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolClassItemList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizeClassItemArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolClassItemArray -> SymbolClassItem
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolClassItemArray -> SymbolClassItemArray SymbolClassItem
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizePageItemFromImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolPageItem -> SymbolImportGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemFromPublicGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolPageItem -> SymbolPublicGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemFromPrivateGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolPageItem -> SymbolPrivateGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemFromProvideGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolClassItem -> SymbolProvideGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemFromRequireGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolClassItem -> SymbolRequireGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolImportGroup -> SymbolImport SymbolIdentifier SymbolString
			path := context.FormatSymbolString(phrase.GetChild(2).GetToken().Value())
			pageIndex, err := context.GetPage(path)
			if err != nil {
				return nil, nil, err
			}
			page, exception := pageIndex.Get(nil)
			if !nl_interface.IsNil(exception) {
				return nil, nil, fmt.Errorf(
					"Page index error: \"%v\"(\"%v\") is not an index without closure, cannot be used as a page index.\n%v",
					path,
					pageIndex.Type(),
					phrase.GetChild(0).GetLine().ToString(),
				)
			}

			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.ImportIndex.New(
						phrase.GetChild(1).GetToken().Value(),
						path,
						page,
					),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(2).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizePublicGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolPublicGroup -> SymbolPublic SymbolIdentifier SymbolEqual SymbolIndex
			name := phrase.GetChild(1).GetToken().Value()
			indexes, _, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.PublicIndex.New(name, indexes[0]),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(3).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizePrivateGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolPrivateGroup -> SymbolPrivate SymbolIdentifier SymbolEqual SymbolIndex
			name := phrase.GetChild(1).GetToken().Value()
			indexes, _, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.PrivateIndex.New(name, indexes[0]),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(3).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeProvideGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolProvideGroup -> SymbolProvide SymbolIdentifier SymbolEqual SymbolFunction
			name := phrase.GetChild(1).GetToken().Value()
			indexes, _, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.ProvideIndex.New(name, indexes[0]),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(3).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeRequireGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolRequireGroup -> SymbolRequire SymbolIdentifier SymbolEqual SymbolDefineFunctionGroup
			name := phrase.GetChild(1).GetToken().Value()
			indexes, _, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.RequireIndex.New(name, indexes[0]),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(3).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeExpressionFromIdentifier, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolIdentifier
			bubble := context.GetLibraryManager().Sandbox.Index.BubbleIndex.New(
				context.GetLibraryManager().Sandbox.Variable.String.New(
					phrase.GetChild(0).GetToken().Value(),
				),
			)
			bubble.SetLine(phrase.GetChild(0).GetLine())
			return []concept.Pipe{bubble}, []*lexer.Line{
				phrase.GetChild(0).GetLine(),
			}, nil
		}),
		semantic.NewRule(PolymerizeExpressionFromVariable, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolVariable
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexFromExpression, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolIndex -> SymbolExpressionCeil
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromNumber, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolNumber
			value, err := strconv.ParseFloat(phrase.GetChild(0).GetToken().Value(), 64)
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewNumber.New(value),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromString, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolString
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewString.New(
						context.FormatSymbolString(
							phrase.GetChild(0).GetToken().Value(),
						),
					),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromBool, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolBool
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolObject
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolFunctionGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromClassGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolClassGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeBoolFromTrue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolBool -> SymbolTrue
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewBool.New(true),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeBoolFromFalse, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolBool -> SymbolFalse
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewBool.New(false),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromNull, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolVariable -> SymbolNull
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewNull.New(),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolArray -> SymbolLeftBracket SymbolIndexList SymbolRightBracket
			items, _, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, nil, err
			}
			newArray := context.GetLibraryManager().Sandbox.Expression.NewArray.New()
			newArray.SetItems(items)
			return []concept.Pipe{newArray}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(2).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolObject -> SymbolLeftBrace SymbolKeyValueList SymbolRightBrace
			fields, preLines, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, nil, err
			}
			newObject := context.GetLibraryManager().Sandbox.Expression.NewObject.New()
			lines := make([]concept.Line, len(preLines))
			for cursor, line := range preLines {
				lines[cursor] = line
			}
			err = newObject.SetKeyValue(fields, lines)
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{newObject}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(2).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolFunctionGroup -> SymbolFunction SymbolLeftParenthesis SymbolKeyList SymbolRightParenthesis SymbolKeyList SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			params, preParamLines, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			returns, preReturnLines, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, nil, err
			}
			steps, _, err := context.Deal(phrase.GetChild(6))
			if err != nil {
				return nil, nil, err
			}
			newFunction := context.GetLibraryManager().Sandbox.Expression.NewFunction.New()
			paramLines := make([]concept.Line, len(preParamLines))
			for cursor, paramLine := range preParamLines {
				paramLines[cursor] = paramLine
			}
			returnLines := make([]concept.Line, len(preReturnLines))
			for cursor, returnLine := range preReturnLines {
				returnLines[cursor] = returnLine
			}
			err = newFunction.SetParams(params, paramLines)
			if err != nil {
				return nil, nil, err
			}
			err = newFunction.SetReturns(returns, returnLines)
			if err != nil {
				return nil, nil, err
			}
			newFunction.SetSteps(steps)
			return []concept.Pipe{newFunction}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(7).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeDefineFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolDefineFunctionGroup -> SymbolFunction SymbolLeftParenthesis SymbolKeyList SymbolRightParenthesis SymbolKeyList
			params, preParamLines, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			returns, preReturnLines, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, nil, err
			}
			newFunction := context.GetLibraryManager().Sandbox.Expression.NewDefineFunction.New()
			paramLines := make([]concept.Line, len(preParamLines))
			for cursor, paramLine := range preParamLines {
				paramLines[cursor] = paramLine
			}
			returnLines := make([]concept.Line, len(preReturnLines))
			for cursor, returnLine := range preReturnLines {
				returnLines[cursor] = returnLine
			}
			err = newFunction.SetParams(params, paramLines)
			if err != nil {
				return nil, nil, err
			}
			err = newFunction.SetReturns(returns, returnLines)
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{newFunction}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(4).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeIndexList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolIndexList -> SymbolIndexArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolIndexList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizeIndexArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolIndexArray -> SymbolIndex
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolIndexArray -> SymbolIndexArray SymbolComma SymbolIndex
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizeKeyValue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyValue -> SymbolIdentifier SymbolColon SymbolIndex
			indexes, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.KeyValueIndex.New(
						context.GetLibraryManager().Sandbox.Variable.String.New(
							phrase.GetChild(0).GetToken().Value(),
						),
						indexes[0],
					),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(2).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeKeyValueList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolKeyValueList -> SymbolKeyValueArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyValueListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolKeyValueList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizeKeyValueArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyValueArray -> SymbolKeyValue
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyValueArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyValueArray -> SymbolKeyValueArray SymbolComma SymbolKeyValue
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizeKeyKey, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyKey -> SymbolIdentifier SymbolColon SymbolIdentifier
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.KeyKeyIndex.New(
						context.GetLibraryManager().Sandbox.Variable.String.New(
							phrase.GetChild(0).GetToken().Value(),
						),
						context.GetLibraryManager().Sandbox.Variable.String.New(
							phrase.GetChild(2).GetToken().Value(),
						),
					),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(2).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeKeyKeyList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolKeyKeyList -> SymbolKeyKeyArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyKeyListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolKeyKeyList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizeKeyKeyArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyKeyArray -> SymbolKeyKey
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyKeyArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyKeyArray -> SymbolKeyKeyArray SymbolComma SymbolKeyKey
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizeMappingObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionCeil -> SymbolExpressionCeil SymbolRightArrow SymbolExpression1 SymbolLeftBrace SymbolKeyKeyList SymbolRightBrace
			object, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			class, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			mapping, preMappingLines, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, nil, err
			}
			mappingLines := make([]concept.Line, len(preMappingLines))
			for cursor, mappingLine := range preMappingLines {
				preMappingLines[cursor] = mappingLine
			}
			newMappingObject := context.GetLibraryManager().Sandbox.Expression.NewMappingObject.New()
			newMappingObject.SetObject(object[0])
			newMappingObject.SetClass(class[0])
			newMappingObject.SetMapping(mapping, mappingLines)
			newMappingObject.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{newMappingObject}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(5).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeMappingObjectAuto, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionCeil -> SymbolExpressionCeil SymbolRightArrow SymbolExpression1
			object, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			class, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			newMappingObject := context.GetLibraryManager().Sandbox.Expression.NewMappingObject.New()
			newMappingObject.SetObject(object[0])
			newMappingObject.SetClass(class[0])
			newMappingObject.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{newMappingObject}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(2).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeCallWithIndexArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolExpression1 SymbolLeftParenthesis SymbolIndexArray SymbolRightParenthesis
			funcs, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			params, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			newParam := context.GetLibraryManager().Sandbox.Expression.NewParam.New()
			newParam.SetArray(params)
			call := context.GetLibraryManager().Sandbox.Expression.Call.New(
				funcs[0],
				newParam,
			)
			call.SetCallLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{call}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(3).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeCallWithKeyValueList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolExpression1 SymbolLeftParenthesis SymbolKeyValueList SymbolRightParenthesis
			funcs, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			params, preLines, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			lines := make([]concept.Line, len(preLines))
			for cursor, line := range preLines {
				lines[cursor] = line
			}
			newParam := context.GetLibraryManager().Sandbox.Expression.NewParam.New()
			err = newParam.SetKeyValue(params, lines)
			if err != nil {
				return nil, nil, err
			}
			call := context.GetLibraryManager().Sandbox.Expression.Call.New(
				funcs[0],
				newParam,
			)
			call.SetCallLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{call}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(3).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeAssignment, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolExpressionCeil SymbolEqual SymbolExpressionCeil
			toes, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			froms, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			instance := context.GetLibraryManager().Sandbox.Expression.Assignment.New(froms[0], toes[0])
			instance.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{instance}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(2).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeAppend, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionCeil -> SymbolExpressionCeil SymbolLeftArrow SymbolExpression1
			array, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}

			instance := context.GetLibraryManager().Sandbox.Expression.Append.New(array[0], item[0])
			instance.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{instance}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(2).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeComponent, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolExpressionFloor SymbolDot SymbolIdentifier
			indexes, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}

			component := context.GetLibraryManager().Sandbox.Expression.Component.New(
				indexes[0],
				context.GetLibraryManager().Sandbox.Variable.String.New(
					phrase.GetChild(2).GetToken().Value(),
				),
			)
			component.SetFieldLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{component}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(2).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeIndexComponent, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolExpressionFloor SymbolLeftBracket SymbolExpressionCeil SymbolRightBracket
			indexes, _, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			field, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}

			component := context.GetLibraryManager().Sandbox.Expression.IndexComponent.New(
				indexes[0],
				field[0],
			)
			component.SetFieldLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{component}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(3).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeDefine, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolVar SymbolIdentifier
			instance := context.GetLibraryManager().Sandbox.Expression.Define.New(
				context.GetLibraryManager().Sandbox.Variable.String.New(
					phrase.GetChild(1).GetToken().Value(),
				),
				nil,
			)
			instance.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{instance}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(1).GetLine(),
			)}, nil
		}),

		semantic.NewRule(PolymerizeDefineAndInit, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolVar SymbolIdentifier SymbolEqual SymbolExpressionCeil
			defaultValue, _, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, nil, err
			}
			instance := context.GetLibraryManager().Sandbox.Expression.Define.New(
				context.GetLibraryManager().Sandbox.Variable.String.New(
					phrase.GetChild(1).GetToken().Value(),
				),
				defaultValue[0],
			)
			instance.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{instance}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(3).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeParentheses, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpressionFloor -> SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis
			target, _, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, nil, err
			}
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.Parenthesis.New(target[0]),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(2).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeIf, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolIf SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			steps, _, err := context.Deal(phrase.GetChild(5))
			if err != nil {
				return nil, nil, err
			}
			eif := context.GetLibraryManager().Sandbox.Expression.If.New()
			eif.SetCondition(condition[0])
			eif.Primary().AddStep(steps...)
			eif.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{eif}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(6).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeIfElse, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolIf SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace SymbolElse SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			primarySteps, _, err := context.Deal(phrase.GetChild(5))
			if err != nil {
				return nil, nil, err
			}
			secondarySteps, _, err := context.Deal(phrase.GetChild(9))
			if err != nil {
				return nil, nil, err
			}
			eif := context.GetLibraryManager().Sandbox.Expression.If.New()
			eif.SetCondition(condition[0])
			eif.Primary().AddStep(primarySteps...)
			eif.Secondary().AddStep(secondarySteps...)
			eif.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{eif}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(10).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeFor, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolFor SymbolLeftParenthesis SymbolExpressionIndependentList SymbolSemicolon SymbolExpressionCeil SymbolSemicolon SymbolExpressionIndependentList SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			initSteps, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			condition, _, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, nil, err
			}
			endSteps, _, err := context.Deal(phrase.GetChild(6))
			if err != nil {
				return nil, nil, err
			}
			bodySteps, _, err := context.Deal(phrase.GetChild(9))
			if err != nil {
				return nil, nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetCondition(condition[0])
			efor.Init().AddStep(initSteps...)
			efor.End().AddStep(endSteps...)
			efor.Body().AddStep(bodySteps...)
			efor.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{efor}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(10).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeForTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolIdentifier SymbolColon SymbolFor SymbolLeftParenthesis SymbolExpressionIndependentList SymbolSemicolon SymbolExpressionCeil SymbolSemicolon SymbolExpressionIndependentList SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			initSteps, _, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, nil, err
			}
			condition, _, err := context.Deal(phrase.GetChild(6))
			if err != nil {
				return nil, nil, err
			}
			endSteps, _, err := context.Deal(phrase.GetChild(8))
			if err != nil {
				return nil, nil, err
			}
			bodySteps, _, err := context.Deal(phrase.GetChild(11))
			if err != nil {
				return nil, nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetTag(context.GetLibraryManager().Sandbox.Variable.String.New(
				phrase.GetChild(0).GetToken().Value(),
			))
			efor.SetCondition(condition[0])
			efor.Init().AddStep(initSteps...)
			efor.End().AddStep(endSteps...)
			efor.Body().AddStep(bodySteps...)
			efor.SetLine(phrase.GetChild(3).GetLine())
			return []concept.Pipe{efor}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(12).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeWhile, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolFor SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, _, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			bodySteps, _, err := context.Deal(phrase.GetChild(5))
			if err != nil {
				return nil, nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetCondition(condition[0])
			efor.Body().AddStep(bodySteps...)
			efor.SetLine(phrase.GetChild(1).GetLine())
			return []concept.Pipe{efor}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(6).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeWhileTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolIdentifier SymbolColon SymbolFor SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, _, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, nil, err
			}
			bodySteps, _, err := context.Deal(phrase.GetChild(7))
			if err != nil {
				return nil, nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetTag(context.GetLibraryManager().Sandbox.Variable.String.New(
				phrase.GetChild(0).GetToken().Value(),
			))
			efor.SetCondition(condition[0])
			efor.Body().AddStep(bodySteps...)
			efor.SetLine(phrase.GetChild(3).GetLine())
			return []concept.Pipe{efor}, []*lexer.Line{lexer.NewLineFromTo(
				phrase.GetChild(0).GetLine(),
				phrase.GetChild(8).GetLine(),
			)}, nil
		}),
		semantic.NewRule(PolymerizeExpression1FromFloor, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpression1 -> SymbolExpressionFloor
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionCeilFrom1, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpressionCeil -> SymbolExpression1
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentFromCeil, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpressionIndependent -> SymbolExpressionCeil
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpression, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpression -> SymbolExpressionIndependent SymbolSemicolon
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpressionIndependentList -> SymbolExpressionIndependentArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpressionIndependentList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizeExpressionIndependentArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependentArray -> SymbolExpressionIndependent
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependentArray -> SymbolExpressionIndependentArray SymbolComma SymbolExpressionIndependent
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizeExpressionList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpressionList -> SymbolExpressionArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolExpressionList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizeExpressionArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionArray -> SymbolExpression
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionArray -> SymbolExpressionArray SymbolExpression
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizeKey, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKey -> SymbolIdentifier
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Index.KeyIndex.New(
						context.GetLibraryManager().Sandbox.Variable.String.New(
							phrase.GetChild(0).GetToken().Value(),
						),
					),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeKeyList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolKeyList -> SymbolKeyArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			// SymbolKeyList ->
			return []concept.Pipe{}, []*lexer.Line{}, nil
		}),
		semantic.NewRule(PolymerizeKeyArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyArray -> SymbolKey
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolKeyArray -> SymbolKeyArray SymbolComma SymbolKey
			items, lines, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, nil, err
			}
			item, line, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, nil, err
			}
			return append(items, item...), append(lines, line...), nil
		}),
		semantic.NewRule(PolymerizeContinue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolContinue
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewContinue.New(
						context.GetLibraryManager().Sandbox.Variable.String.New(""),
					),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeContinueTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolContinue SymbolIdentifier
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewContinue.New(
						context.GetLibraryManager().Sandbox.Variable.String.New(
							phrase.GetChild(1).GetToken().Value(),
						),
					),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(1).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeBreak, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolBreak
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewBreak.New(
						context.GetLibraryManager().Sandbox.Variable.String.New(""),
					),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeBreakTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolBreak SymbolIdentifier
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewBreak.New(
						context.GetLibraryManager().Sandbox.Variable.String.New(
							phrase.GetChild(1).GetToken().Value(),
						),
					),
				}, []*lexer.Line{lexer.NewLineFromTo(
					phrase.GetChild(0).GetLine(),
					phrase.GetChild(1).GetLine(),
				)}, nil
		}),
		semantic.NewRule(PolymerizeReturn, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionIndependent -> SymbolReturn
			return []concept.Pipe{
					context.GetLibraryManager().Sandbox.Expression.NewReturn.New(),
				}, []*lexer.Line{
					phrase.GetChild(0).GetLine(),
				}, nil
		}),
		semantic.NewRule(PolymerizeThis, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolThis
			this := context.GetLibraryManager().Sandbox.Index.ThisIndex.New()
			this.SetLine(phrase.GetChild(0).GetLine())
			return []concept.Pipe{this}, []*lexer.Line{
				phrase.GetChild(0).GetLine(),
			}, nil
		}),
		semantic.NewRule(PolymerizeSelf, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, []*lexer.Line, error) {
			//SymbolExpressionFloor -> SymbolSelf
			self := context.GetLibraryManager().Sandbox.Index.SelfIndex.New()
			self.SetLine(phrase.GetChild(0).GetLine())
			return []concept.Pipe{self}, []*lexer.Line{
				phrase.GetChild(0).GetLine(),
			}, nil
		}),
	}
)

func init() {
	if len(SemanticRules) < len(GrammarRules) {
		panic("Rule error: The lengths of SemanticRules and GrammarRules are not equal.")
	}
}
