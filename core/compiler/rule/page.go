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
				return nil, errors.New(fmt.Sprintf("Unsupport index to be set: %v", item.ToString("")))
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ConstIndex.New(page),
			}, nil
		}),
		semantic.NewRule(PolymerizePageItemList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolPageItemList -> SymbolPageItemArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolPageItemList ->
			return []concept.Index{}, nil
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
		semantic.NewRule(PolymerizeClassGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolClassGroup -> SymbolClass SymbolLeftBrace SymbolClassItemList SymbolRightBrace
			items, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			newClass := context.GetLibraryManager().Sandbox.Expression.NewClass.New()
			newClass.SetLines(items)
			return []concept.Index{newClass}, nil
		}),
		semantic.NewRule(PolymerizeClassItemList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolClassItemList -> SymbolClassItemArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolClassItemList ->
			return []concept.Index{}, nil
		}),
		semantic.NewRule(PolymerizeClassItemArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolClassItemArray -> SymbolClassItem
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
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
		semantic.NewRule(PolymerizePageItemFromImportGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPageItem -> SymbolImportGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemFromPublicGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPageItem -> SymbolPublicGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizePageItemFromPrivateGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPageItem -> SymbolPrivateGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemFromProvideGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolClassItem -> SymbolProvideGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeClassItemFromRequireGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolClassItem -> SymbolRequireGroup
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
		semantic.NewRule(PolymerizePublicGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPublicGroup -> SymbolPublic SymbolIdentifier SymbolEqual SymbolIndex
			name := phrase.GetChild(1).GetToken().Value()
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.PublicIndex.New(name, indexes[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizePrivateGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolPrivateGroup -> SymbolPrivate SymbolIdentifier SymbolEqual SymbolIndex
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.PrivateIndex.New(
					phrase.GetChild(1).GetToken().Value(),
					indexes[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeProvideGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolProvideGroup -> SymbolProvide SymbolIdentifier SymbolEqual SymbolFunction
			name := phrase.GetChild(1).GetToken().Value()
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ProvideIndex.New(name, indexes[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizeRequireGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolRequireGroup -> SymbolRequire SymbolIdentifier SymbolEqual SymbolDefineFunctionGroup
			indexes, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.RequireIndex.New(
					phrase.GetChild(1).GetToken().Value(),
					indexes[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeExpressionFromIdentifier, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpression2 -> SymbolIdentifier
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.BubbleIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(0).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeExpressionFromVariable, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpression2 -> SymbolVariable
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexFromExpression, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolIndex -> SymbolExpression1
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromNumber, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVariable -> SymbolNumber
			value, err := strconv.ParseFloat(phrase.GetChild(0).GetToken().Value(), 64)
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewNumber.New(value),
			}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromString, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVariable -> SymbolString
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewString.New(
					context.FormatSymbolString(
						phrase.GetChild(0).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromBool, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVariable -> SymbolBool
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVariable -> SymbolObject
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVariable -> SymbolFunctionGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeVariableFromClassGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVariable -> SymbolClassGroup
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeBoolFromTrue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolBool -> SymbolTrue
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewBool.New(true),
			}, nil
		}),
		semantic.NewRule(PolymerizeBoolFromFalse, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolBool -> SymbolFalse
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewBool.New(false),
			}, nil
		}),
		semantic.NewRule(PolymerizeVariableFromNull, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolVariable -> SymbolNull
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewNull.New(),
			}, nil
		}),
		semantic.NewRule(PolymerizeObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolObject -> SymbolLeftBrace SymbolKeyValueList SymbolRightBrace
			fields, err := context.Deal(phrase.GetChild(1))
			if err != nil {
				return nil, err
			}
			newObject := context.GetLibraryManager().Sandbox.Expression.NewObject.New()
			newObject.SetKeyValue(fields)
			return []concept.Index{newObject}, nil
		}),
		semantic.NewRule(PolymerizeFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
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
			return []concept.Index{newFunction}, nil
		}),
		semantic.NewRule(PolymerizeDefineFunctionGroup, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
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
			return []concept.Index{newFunction}, nil
		}),
		semantic.NewRule(PolymerizeIndexList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolIndexList -> SymbolIndexArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeIndexListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolIndexList ->
			return []concept.Index{}, nil
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
		semantic.NewRule(PolymerizeKeyValueList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolKeyValueList -> SymbolKeyValueArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyValueListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolKeyValueList ->
			return []concept.Index{}, nil
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
		semantic.NewRule(PolymerizeKeyKey, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolKeyKey -> SymbolIdentifier SymbolColon SymbolIdentifier
			return []concept.Index{
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
		semantic.NewRule(PolymerizeKeyKeyList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolKeyKeyList -> SymbolKeyKeyArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyKeyListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolKeyKeyList ->
			return []concept.Index{}, nil
		}),
		semantic.NewRule(PolymerizeKeyKeyArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolKeyKeyArray -> SymbolKeyKey
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyKeyArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
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
		semantic.NewRule(PolymerizeMappingObject, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolObject -> SymbolIndex SymbolRightArrow SymbolIndex SymbolLeftBrace SymbolKeyKeyList SymbolRightBrace
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
			return []concept.Index{newMappingObject}, nil
		}),
		semantic.NewRule(PolymerizeCallWithIndexArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpression1 -> SymbolExpression1 SymbolLeftParenthesis SymbolIndexArray SymbolRightParenthesis
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
		semantic.NewRule(PolymerizeCallWithKeyValueList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpression1 -> SymbolExpression1 SymbolLeftParenthesis SymbolKeyValueList SymbolRightParenthesis
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
		semantic.NewRule(PolymerizeAssignment, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolExpression1 SymbolEqual SymbolExpression1
			toes, err := context.Deal(phrase.GetChild(0))
			if err != nil {
				return nil, err
			}
			froms, err := context.Deal(phrase.GetChild(2))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.Assignment.New(froms[0], toes[0]),
			}, nil
		}),
		semantic.NewRule(PolymerizeComponent, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpression1 -> SymbolExpression1 SymbolDot SymbolIdentifier
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

		semantic.NewRule(PolymerizeDefine, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolVar SymbolIdentifier
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.Define.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
					nil,
				),
			}, nil
		}),

		semantic.NewRule(PolymerizeDefineAndInit, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolVar SymbolIdentifier SymbolEqual SymbolExpression1
			defaultValue, err := context.Deal(phrase.GetChild(3))
			if err != nil {
				return nil, err
			}
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.Define.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
					defaultValue[0],
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeParentheses, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpression2 -> SymbolLeftParenthesis SymbolExpression1 SymbolRightParenthesis
			return context.Deal(phrase.GetChild(1))
		}),
		semantic.NewRule(PolymerizeIf, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolIf SymbolLeftParenthesis SymbolExpression1 SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
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
			return []concept.Index{eif}, nil
		}),
		semantic.NewRule(PolymerizeIfElse, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolIf SymbolLeftParenthesis SymbolExpression1 SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace SymbolElse SymbolLeftBrace SymbolExpressionList SymbolRightBrace
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
			return []concept.Index{eif}, nil
		}),
		semantic.NewRule(PolymerizeFor, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolFor SymbolLeftParenthesis SymbolExpressionIndependentList SymbolSemicolon SymbolExpression1 SymbolSemicolon SymbolExpressionIndependentList SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
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
			return []concept.Index{efor}, nil
		}),
		semantic.NewRule(PolymerizeWhile, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolFor SymbolLeftParenthesis SymbolExpression1 SymbolRightParenthesis SymbolLeftBrace SymbolExpressionList SymbolRightBrace
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
			return []concept.Index{efor}, nil
		}),
		semantic.NewRule(PolymerizeExpression2To1, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpression1 -> SymbolExpression2
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpression1ToIndependent, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpressionIndependent -> SymbolExpression1
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpression, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpression -> SymbolExpressionIndependent SymbolSemicolon
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpressionIndependentList -> SymbolExpressionIndependentArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpressionIndependentList ->
			return []concept.Index{}, nil
		}),
		semantic.NewRule(PolymerizeExpressionIndependentArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependentArray -> SymbolExpressionIndependent
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionIndependentArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
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
		semantic.NewRule(PolymerizeExpressionList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpressionList -> SymbolExpressionArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolExpressionList ->
			return []concept.Index{}, nil
		}),
		semantic.NewRule(PolymerizeExpressionArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionArray -> SymbolExpression
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeExpressionArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
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
		semantic.NewRule(PolymerizeKey, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolKey -> SymbolIdentifier
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.KeyIndex.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(0).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeKeyList, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolKeyList -> SymbolKeyArray
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyListEmpty, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			// SymbolKeyList ->
			return []concept.Index{}, nil
		}),
		semantic.NewRule(PolymerizeKeyArrayStart, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolKeyArray -> SymbolKey
			return context.Deal(phrase.GetChild(0))
		}),
		semantic.NewRule(PolymerizeKeyArray, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
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
		semantic.NewRule(PolymerizeContinue, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolContinue
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewContinue.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(""),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeContinueTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolContinue SymbolIdentifier
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewContinue.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeBreak, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolBreak
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewBreak.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(""),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeBreakTag, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolBreak SymbolIdentifier
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewBreak.New(
					context.GetLibraryManager().Sandbox.Variable.String.New(
						phrase.GetChild(1).GetToken().Value(),
					),
				),
			}, nil
		}),
		semantic.NewRule(PolymerizeReturn, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpressionIndependent -> SymbolReturn
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Expression.NewReturn.New(),
			}, nil
		}),
		semantic.NewRule(PolymerizeThis, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpression2 -> SymbolThis
			return []concept.Index{
				context.GetLibraryManager().Sandbox.Index.ThisIndex.New(),
			}, nil
		}),
		semantic.NewRule(PolymerizeSelf, func(phrase grammar.Phrase, context *semantic.Context) ([]concept.Index, error) {
			//SymbolExpression2 -> SymbolSelf
			return []concept.Index{
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
