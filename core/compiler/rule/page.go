package rule

import (
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

		semantic.NewRule(PolymerizePageGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolPageGroup -> SymbolPageItemList
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
				publicIndex, yes := index.IndexFamilyInstance.IsPublicIndex(item)
				if yes {
					err := page.SetPublic(
						context.GetLibraryManager().Sandbox.Variable.String.New(publicIndex.Name()),
						publicIndex,
					)
					if !nl_interface.IsNil(err) {
						return nil, err
					}
					continue
				}
				privateIndex, yes := index.IndexFamilyInstance.IsPrivateIndex(item)
				if yes {
					err := page.SetPrivate(
						context.GetLibraryManager().Sandbox.Variable.String.New(privateIndex.Name()),
						privateIndex,
					)
					if !nl_interface.IsNil(err) {
						return nil, err
					}
					continue
				}
				return nil, fmt.Errorf("Unsupport index to be set: %v", item.ToString(""))
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.ConstIndex.New(page),
			}, nil
		}),
		semantic.NewRule(PolymerizePageItemList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolPageItemList -> SymbolPageItemArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolPageItemList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizePageItemArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolPageItemArray -> SymbolPageItem
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
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
		semantic.NewRule(PolymerizeClassGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolClassGroup -> SymbolClass SymbolLeftBrace SymbolClassItemList SymbolRightBrace
			items, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			newClass := context.GetLibraryManager().Sandbox.Expression.NewClass.New()
			newClass.SetLines(items)
			return []concept.Pipe{newClass}, nil
		}),
		semantic.NewRule(PolymerizeClassItemList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolClassItemList -> SymbolClassItemArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolClassItemList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizeClassItemArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolClassItemArray -> SymbolClassItem
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolClassItemArray -> SymbolClassItemArray SymbolClassItem
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
		semantic.NewRule(PolymerizePageItemFromImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolPageItem -> SymbolImportGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemFromPublicGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolPageItem -> SymbolPublicGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemFromPrivateGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolPageItem -> SymbolPrivateGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemFromProvideGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolClassItem -> SymbolProvideGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemFromRequireGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolClassItem -> SymbolRequireGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolImportGroup -> SymbolImport SymbolIdentifier SymbolString
			path := context.FormatSymbolString(phrase.GetChild(2).GetToken().Value())
			pageIndex, err := context.GetPage(path)
			if err != nil {
				return nil, err
			}
			page, exception := pageIndex.Get(nil)
			if !nl_interface.IsNil(exception) {
				return nil, fmt.Errorf(
					"Page index error: \"%v\"(\"%v\") is not an index without closure, cannot be used as a page index.",
					path,
					pageIndex.Type(),
				)
			}

			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.ImportIndex.New(
					phrase.GetChild(1).GetToken().Value(),
					path,
					page,
				),
			}, nil
		}),
		semantic.NewRule(PolymerizePublicGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolPublicGroup -> SymbolPublic SymbolIdentifier SymbolEqual SymbolIndex
			name := phrase.GetChild(1).GetToken().Value()
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.PublicIndex.New(name, indexes[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizePrivateGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolPrivateGroup -> SymbolPrivate SymbolIdentifier SymbolEqual SymbolIndex
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.PrivateIndex.New(
					phrase.GetChild(1).GetToken().Value(),
					indexes[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeProvideGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolProvideGroup -> SymbolProvide SymbolIdentifier SymbolEqual SymbolFunction
			name := phrase.GetChild(1).GetToken().Value()
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.ProvideIndex.New(name, indexes[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizeRequireGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolRequireGroup -> SymbolRequire SymbolIdentifier SymbolEqual SymbolDefineFunctionGroup
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.RequireIndex.New(
					phrase.GetChild(1).GetToken().Value(),
					indexes[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeExpressionFromIdentifier, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolIdentifier
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.BubbleIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(0).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeExpressionFromVariable, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolVariable
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexFromExpression, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolIndex -> SymbolExpressionCeil
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromNumber, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolNumber
			value, err := strconv.ParseFloat(phrase.GetChild(0).GetToken().Value(), 64)
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewNumber.New(value),
			}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromString, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolString
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewString.New(
					context.FormatSymbolString(
						phrase.GetChild(0).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromBool, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolBool
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolObject
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolFunctionGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromClassGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolClassGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeBoolFromTrue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolBool -> SymbolTrue
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewBool.New(true),
			}, nil
		}),
		semantic.NewRule(PolymerizeBoolFromFalse, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolBool -> SymbolFalse
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewBool.New(false),
			}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromNull, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolVariable -> SymbolNull
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewNull.New(),
			}, nil
		}),
		semantic.NewRule(PolymerizeArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolArray -> SymbolLeftBracket SymbolIndexList SymbolRightBracket
			items, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, err
			}
			newArray := context.GetLibraryManager().Sandbox.Expression.NewArray.New()
			newArray.SetItems(items)
			return []concept.Pipe{newArray}, nil
		}),
		semantic.NewRule(PolymerizeObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolObject -> SymbolLeftBrace SymbolKeyValueList SymbolRightBrace
			fields, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, err
			}
			newObject := context.GetLibraryManager().Sandbox.Expression.NewObject.New()
			newObject.SetKeyValue(fields)
			return []concept.Pipe{newObject}, nil
		}),
		semantic.NewRule(PolymerizeFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolFunctionGroup -> SymbolFunction SymbolLeftParenthesis SymbolKeyList SymbolRightParenthesis SymbolKeyList SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			params, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			returns, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, err
			}
			steps, err := context.Deal(phrase.GetChild(6))
			if err != nil {
				return nil, err
			}
			newFunction := context.GetLibraryManager().Sandbox.Expression.NewFunction.New()
			newFunction.SetParams(params)
			newFunction.SetReturns(returns)
			newFunction.SetSteps(steps)
			return []concept.Pipe{newFunction}, nil
		}),
		semantic.NewRule(PolymerizeDefineFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolDefineFunctionGroup -> SymbolFunction SymbolLeftParenthesis SymbolKeyList SymbolRightParenthesis SymbolKeyList
			params, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			returns, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, err
			}
			newFunction := context.GetLibraryManager().Sandbox.Expression.NewDefineFunction.New()
			newFunction.SetParams(params)
			newFunction.SetReturns(returns)
			return []concept.Pipe{newFunction}, nil
		}),
		semantic.NewRule(PolymerizeIndexList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolIndexList -> SymbolIndexArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolIndexList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizeIndexArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolIndexArray -> SymbolIndex
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
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
		semantic.NewRule(PolymerizeKeyValue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolKeyValue -> SymbolIdentifier SymbolColon SymbolIndex
			indexes, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.KeyValueIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(0).GetToken().Value(),
					),
					indexes[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeKeyValueList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolKeyValueList -> SymbolKeyValueArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyValueListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolKeyValueList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizeKeyValueArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolKeyValueArray -> SymbolKeyValue
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyValueArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
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
		semantic.NewRule(PolymerizeKeyKey, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
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
			}, nil
		}),
		semantic.NewRule(PolymerizeKeyKeyList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolKeyKeyList -> SymbolKeyKeyArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyKeyListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolKeyKeyList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizeKeyKeyArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolKeyKeyArray -> SymbolKeyKey
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyKeyArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolKeyKeyArray -> SymbolKeyKeyArray SymbolComma SymbolKeyKey
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
		semantic.NewRule(PolymerizeMappingObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionCeil -> SymbolExpressionCeil SymbolRightArrow SymbolExpression1 SymbolLeftBrace SymbolKeyKeyList SymbolRightBrace
			object, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			class, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			mapping, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, err
			}
			newMappingObject := context.GetLibraryManager().Sandbox.Expression.NewMappingObject.New()
			newMappingObject.SetObject(object[0])
			newMappingObject.SetClass(class[0])
			newMappingObject.SetMapping(mapping)
			return []concept.Pipe{newMappingObject}, nil
		}),
		semantic.NewRule(PolymerizeMappingObjectAuto, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionCeil -> SymbolExpressionCeil SymbolRightArrow SymbolExpression1
			object, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			class, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			newMappingObject := context.GetLibraryManager().Sandbox.Expression.NewMappingObject.New()
			newMappingObject.SetObject(object[0])
			newMappingObject.SetClass(class[0])
			return []concept.Pipe{newMappingObject}, nil
		}),
		semantic.NewRule(PolymerizeCallWithIndexArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolExpression1 SymbolLeftParenthesis SymbolIndexArray SymbolRightParenthesis
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
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Call.New(
					funcs[0],
					newParam,
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeCallWithKeyValueList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolExpression1 SymbolLeftParenthesis SymbolKeyValueList SymbolRightParenthesis
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
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Call.New(
					funcs[0],
					newParam,
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeAssignment, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolExpressionCeil SymbolEqual SymbolExpressionCeil
			toes, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			froms, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Assignment.New(froms[0], toes[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizeAppend, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionCeil -> SymbolExpressionCeil SymbolLeftArrow SymbolExpression1
			array, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			item, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Append.New(array[0], item[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizeComponent, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolExpressionFloor SymbolDot SymbolIdentifier
			indexes, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Component.New(
					indexes[0],
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(2).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeIndexComponent, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolExpressionFloor SymbolLeftBracket SymbolExpressionCeil SymbolRightBracket
			indexes, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			field, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.IndexComponent.New(
					indexes[0],
					field[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeDefine, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolVar SymbolIdentifier
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Define.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
					nil,
				),
			}, nil
		}),

		semantic.NewRule(PolymerizeDefineAndInit, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolVar SymbolIdentifier SymbolEqual SymbolExpressionCeil
			defaultValue, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Define.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
					defaultValue[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeParentheses, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpressionFloor -> SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis
			target, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, err
			}
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.Parenthesis.New(target[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizeIf, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolIf SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			steps, err := context.Deal(phrase.GetChild(5))
			if err != nil {
				return nil, err
			}
			eif := context.GetLibraryManager().Sandbox.Expression.If.New()
			eif.SetCondition(condition[0])
			eif.Primary().AddStep(steps...)
			return []concept.Pipe{eif}, nil
		}),
		semantic.NewRule(PolymerizeIfElse, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolIf SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace SymbolElse SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			primarySteps, err := context.Deal(phrase.GetChild(5))
			if err != nil {
				return nil, err
			}
			secondarySteps, err := context.Deal(phrase.GetChild(9))
			if err != nil {
				return nil, err
			}
			eif := context.GetLibraryManager().Sandbox.Expression.If.New()
			eif.SetCondition(condition[0])
			eif.Primary().AddStep(primarySteps...)
			eif.Secondary().AddStep(secondarySteps...)
			return []concept.Pipe{eif}, nil
		}),
		semantic.NewRule(PolymerizeFor, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolFor SymbolLeftParenthesis SymbolExpressionIndependentList SymbolSemicolon SymbolExpressionCeil SymbolSemicolon SymbolExpressionIndependentList SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			initSteps, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			condition, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, err
			}
			endSteps, err := context.Deal(phrase.GetChild(6))
			if err != nil {
				return nil, err
			}
			bodySteps, err := context.Deal(phrase.GetChild(9))
			if err != nil {
				return nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetCondition(condition[0])
			efor.Init().AddStep(initSteps...)
			efor.End().AddStep(endSteps...)
			efor.Body().AddStep(bodySteps...)
			return []concept.Pipe{efor}, nil
		}),
		semantic.NewRule(PolymerizeForTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolIdentifier SymbolColon SymbolFor SymbolLeftParenthesis SymbolExpressionIndependentList SymbolSemicolon SymbolExpressionCeil SymbolSemicolon SymbolExpressionIndependentList SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			initSteps, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, err
			}
			condition, err := context.Deal(phrase.GetChild(6))
			if err != nil {
				return nil, err
			}
			endSteps, err := context.Deal(phrase.GetChild(8))
			if err != nil {
				return nil, err
			}
			bodySteps, err := context.Deal(phrase.GetChild(11))
			if err != nil {
				return nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetTag(context.GetLibraryManager().Sandbox.Variable.String.New(
				phrase.GetChild(0).GetToken().Value(),
			))
			efor.SetCondition(condition[0])
			efor.Init().AddStep(initSteps...)
			efor.End().AddStep(endSteps...)
			efor.Body().AddStep(bodySteps...)
			return []concept.Pipe{efor}, nil
		}),
		semantic.NewRule(PolymerizeWhile, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolFor SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			bodySteps, err := context.Deal(phrase.GetChild(5))
			if err != nil {
				return nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetCondition(condition[0])
			efor.Body().AddStep(bodySteps...)
			return []concept.Pipe{efor}, nil
		}),
		semantic.NewRule(PolymerizeWhileTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolIdentifier SymbolColon SymbolFor SymbolLeftParenthesis SymbolExpressionCeil SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
			condition, err := context.Deal(phrase.GetChild(4))
			if err != nil {
				return nil, err
			}
			bodySteps, err := context.Deal(phrase.GetChild(7))
			if err != nil {
				return nil, err
			}
			efor := context.GetLibraryManager().Sandbox.Expression.For.New()
			efor.SetTag(context.GetLibraryManager().Sandbox.Variable.String.New(
				phrase.GetChild(0).GetToken().Value(),
			))
			efor.SetCondition(condition[0])
			efor.Body().AddStep(bodySteps...)
			return []concept.Pipe{efor}, nil
		}),
		semantic.NewRule(PolymerizeExpression1FromFloor, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpression1 -> SymbolExpressionFloor
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionCeilFrom1, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpressionCeil -> SymbolExpression1
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentFromCeil, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpressionIndependent -> SymbolExpressionCeil
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpression, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpression -> SymbolExpressionIndependent SymbolSemicolon
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpressionIndependentList -> SymbolExpressionIndependentArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpressionIndependentList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizeExpressionIndependentArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependentArray -> SymbolExpressionIndependent
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependentArray -> SymbolExpressionIndependentArray SymbolComma SymbolExpressionIndependent
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
		semantic.NewRule(PolymerizeExpressionList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpressionList -> SymbolExpressionArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolExpressionList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizeExpressionArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionArray -> SymbolExpression
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionArray -> SymbolExpressionArray SymbolExpression
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
		semantic.NewRule(PolymerizeKey, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolKey -> SymbolIdentifier
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.KeyIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(0).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeKeyList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolKeyList -> SymbolKeyArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			// SymbolKeyList ->
			return []concept.Pipe{}, nil
		}),
		semantic.NewRule(PolymerizeKeyArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolKeyArray -> SymbolKey
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolKeyArray -> SymbolKeyArray SymbolComma SymbolKey
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
		semantic.NewRule(PolymerizeContinue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolContinue
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewContinue.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(""),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeContinueTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolContinue SymbolIdentifier
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewContinue.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeBreak, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolBreak
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewBreak.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(""),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeBreakTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolBreak SymbolIdentifier
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewBreak.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeReturn, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionIndependent -> SymbolReturn
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Expression.NewReturn.New(),
			}, nil
		}),
		semantic.NewRule(PolymerizeThis, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolThis
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.ThisIndex.New(),
			}, nil
		}),
		semantic.NewRule(PolymerizeSelf, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Pipe, error) {
			//SymbolExpressionFloor -> SymbolSelf
			return []concept.Pipe{
				context.GetLibraryManager().Sandbox.Index.SelfIndex.New(),
			}, nil
		}),
	}
)

func init() {
	if len(SemanticRules) < len(GrammarRules) {
		panic("Rule error: The lengths of SemanticRules and GrammarRules are not equal.")
	}
}
