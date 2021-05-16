package rule

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
)

/*
	===========terminal=========
	LeftParenthesis
	RightParenthesis
	LeftBracket
	RightBracket
	LeftBrace
	RightBrace
	LeftArrow
	RightArrow
	Space
	Equal
	Colon
	Semicolon
	Dot
	Comma
	Number
	Identifier
	End
	String
	Page
	Import
	Public
	Private
	Class
	Require
	Provide
	Return
	Function
	Get
	Set
	True
	False
	Null
*/

var (
	SymbolLeftParenthesis  = grammar.NewTerminal(TypeLeftParenthesis, KeyLeftParenthesis)
	SymbolRightParenthesis = grammar.NewTerminal(TypeRightParenthesis, KeyRightParenthesis)
	SymbolLeftBracket      = grammar.NewTerminal(TypeLeftBracket, KeyLeftBracket)
	SymbolRightBracket     = grammar.NewTerminal(TypeRightBracket, KeyRightBracket)
	SymbolLeftBrace        = grammar.NewTerminal(TypeLeftBrace, KeyLeftBrace)
	SymbolRightBrace       = grammar.NewTerminal(TypeRightBrace, KeyRightBrace)
	SymbolLeftArrow        = grammar.NewTerminal(TypeLeftArrow, KeyLeftArrow)
	SymbolRightArrow       = grammar.NewTerminal(TypeRightArrow, KeyRightArrow)
	SymbolSpace            = grammar.NewTerminal(TypeSpace, KeySpace)
	SymbolEqual            = grammar.NewTerminal(TypeEqual, KeyEqual)
	SymbolColon            = grammar.NewTerminal(TypeColon, KeyColon)
	SymbolSemicolon        = grammar.NewTerminal(TypeSemicolon, KeySemicolon)
	SymbolDot              = grammar.NewTerminal(TypeDot, KeyDot)
	SymbolComma            = grammar.NewTerminal(TypeComma, KeyComma)
	SymbolNumber           = grammar.NewTerminal(TypeNumber, KeyNumber)
	SymbolIdentifier       = grammar.NewTerminal(TypeIdentifier, KeyIdentifier)
	SymbolEnd              = grammar.NewTerminal(TypeEnd, KeyEnd)
	SymbolString           = grammar.NewTerminal(TypeString, KeyString)
	SymbolComment          = grammar.NewTerminal(TypeComment, KeyComment)
	SymbolPage             = grammar.NewTerminal(TypePage, KeyPage)
	SymbolImport           = grammar.NewTerminal(TypeImport, KeyImport)
	SymbolPublic           = grammar.NewTerminal(TypePublic, KeyPublic)
	SymbolPrivate          = grammar.NewTerminal(TypePrivate, KeyPrivate)
	SymbolClass            = grammar.NewTerminal(TypeClass, KeyClass)
	SymbolRequire          = grammar.NewTerminal(TypeRequire, KeyRequire)
	SymbolProvide          = grammar.NewTerminal(TypeProvide, KeyProvide)
	SymbolReturn           = grammar.NewTerminal(TypeReturn, KeyReturn)
	SymbolFunction         = grammar.NewTerminal(TypeFunction, KeyFunction)
	SymbolGet              = grammar.NewTerminal(TypeGet, KeyGet)
	SymbolSet              = grammar.NewTerminal(TypeSet, KeySet)
	SymbolTrue             = grammar.NewTerminal(TypeTrue, KeyTrue)
	SymbolFalse            = grammar.NewTerminal(TypeFalse, KeyFalse)
	SymbolNull             = grammar.NewTerminal(TypeNull, KeyNull)
)

const (
	TypePageGroup = iota
	TypePageItemList
	TypePageItemArray
	TypePageItem
	TypeImportGroup
	TypePublicGroup
	TypePrivateGroup
	TypeIndex
	TypeIndexList
	TypeIndexArray
	TypeKeyValue
	TypeKeyValueList
	TypeKeyValueArray
	TypeBool
	TypeObject
	TypeFunctionGroup
	TypeVariable
	TypeExpression
	TypeExpression1
	TypeExpression2
	TypeExpression3
	TypeExpressionList
	TypeExpressionArray
	TypeParam
	TypeParamList
	TypeParamArray
	TypeKey
	TypeKeyList
	TypeKeyArray
)

const (
	KeyPageGroup       = "page_group"
	KeyPageItemList    = "page_item_list"
	KeyPageItemArray   = "page_item_array"
	KeyPageItem        = "page_item"
	KeyImportGroup     = "import_group"
	KeyPublicGroup     = "public_group"
	KeyPrivateGroup    = "private_group"
	KeyIndex           = "index"
	KeyIndexList       = "index_list"
	KeyIndexArray      = "index_array"
	KeyKeyValue        = "key_value"
	KeyKeyValueList    = "key_value_list"
	KeyKeyValueArray   = "key_value_array"
	KeyBool            = "bool"
	KeyObject          = "object"
	KeyFunctionGroup   = "function_group"
	KeyVariable        = "variable"
	KeyExpression      = "expression"
	KeyExpression1     = "expression_1"
	KeyExpression2     = "expression_2"
	KeyExpression3     = "expression_3"
	KeyExpressionList  = "expression_list"
	KeyExpressionArray = "expression_array"
	KeyKey             = "param"
	KeyKeyList         = "param_list"
	KeyKeyArray        = "param_array"
)

var (
	SymbolPageGroup       = grammar.NewNonterminal(TypePageGroup, KeyPageGroup)
	SymbolPageItemList    = grammar.NewNonterminal(TypePageItemList, KeyPageItemList)
	SymbolPageItemArray   = grammar.NewNonterminal(TypePageItemArray, KeyPageItemArray)
	SymbolPageItem        = grammar.NewNonterminal(TypePageItem, KeyPageItem)
	SymbolImportGroup     = grammar.NewNonterminal(TypeImportGroup, KeyImportGroup)
	SymbolPublicGroup     = grammar.NewNonterminal(TypePublicGroup, KeyPublicGroup)
	SymbolPrivateGroup    = grammar.NewNonterminal(TypePrivateGroup, KeyPrivateGroup)
	SymbolIndex           = grammar.NewNonterminal(TypeIndex, KeyIndex)
	SymbolIndexList       = grammar.NewNonterminal(TypeIndexList, KeyIndexList)
	SymbolIndexArray      = grammar.NewNonterminal(TypeIndexArray, KeyIndexArray)
	SymbolKeyValue        = grammar.NewNonterminal(TypeKeyValue, KeyKeyValue)
	SymbolKeyValueList    = grammar.NewNonterminal(TypeKeyValueList, KeyKeyValueList)
	SymbolKeyValueArray   = grammar.NewNonterminal(TypeKeyValueArray, KeyKeyValueArray)
	SymbolBool            = grammar.NewNonterminal(TypeBool, KeyBool)
	SymbolObject          = grammar.NewNonterminal(TypeObject, KeyObject)
	SymbolFunctionGroup   = grammar.NewNonterminal(TypeFunctionGroup, KeyFunctionGroup)
	SymbolVariable        = grammar.NewNonterminal(TypeVariable, KeyVariable)
	SymbolExpression      = grammar.NewNonterminal(TypeExpression, KeyExpression)
	SymbolExpression1     = grammar.NewNonterminal(TypeExpression1, KeyExpression1)
	SymbolExpression2     = grammar.NewNonterminal(TypeExpression2, KeyExpression2)
	SymbolExpression3     = grammar.NewNonterminal(TypeExpression3, KeyExpression3)
	SymbolExpressionList  = grammar.NewNonterminal(TypeExpressionList, KeyExpressionList)
	SymbolExpressionArray = grammar.NewNonterminal(TypeExpressionArray, KeyExpressionArray)
	SymbolKey             = grammar.NewNonterminal(TypeKey, KeyKey)
	SymbolKeyList         = grammar.NewNonterminal(TypeKeyList, KeyKeyList)
	SymbolKeyArray        = grammar.NewNonterminal(TypeKeyArray, KeyKeyArray)
)

var (
	PolymerizePageGroup                 = grammar.NewRule(SymbolPageGroup, SymbolPageItemList)
	PolymerizePageItemList              = grammar.NewRule(SymbolPageItemList, SymbolPageItemArray)
	PolymerizePageItemListEmpty         = grammar.NewRule(SymbolPageItemList)
	PolymerizePageItemArrayStart        = grammar.NewRule(SymbolPageItemArray, SymbolPageItem)
	PolymerizePageItemArray             = grammar.NewRule(SymbolPageItemArray, SymbolPageItemArray, SymbolPageItem)
	PolymerizePageItemFromImportGroup   = grammar.NewRule(SymbolPageItem, SymbolImportGroup)
	PolymerizePageItemFromPublicGroup   = grammar.NewRule(SymbolPageItem, SymbolPublicGroup)
	PolymerizePageItemFromPrivateGroup  = grammar.NewRule(SymbolPageItem, SymbolPrivateGroup)
	PolymerizeImportGroup               = grammar.NewRule(SymbolImportGroup, SymbolImport, SymbolIdentifier, SymbolString)
	PolymerizePublicGroup               = grammar.NewRule(SymbolPublicGroup, SymbolPublic, SymbolIdentifier, SymbolEqual, SymbolIndex)
	PolymerizePrivateGroup              = grammar.NewRule(SymbolPrivateGroup, SymbolPrivate, SymbolIdentifier, SymbolEqual, SymbolIndex)
	PolymerizeIndexFromIdentifier       = grammar.NewRule(SymbolExpression3, SymbolIdentifier)
	PolymerizeIndexFromVariable         = grammar.NewRule(SymbolExpression3, SymbolVariable)
	PolymerizeIndexFromExpression       = grammar.NewRule(SymbolIndex, SymbolExpression1)
	PolymerizeVariableFromNumber        = grammar.NewRule(SymbolVariable, SymbolNumber)
	PolymerizeVariableFromString        = grammar.NewRule(SymbolVariable, SymbolString)
	PolymerizeVariableFromBool          = grammar.NewRule(SymbolVariable, SymbolBool)
	PolymerizeVariableFromObject        = grammar.NewRule(SymbolVariable, SymbolObject)
	PolymerizeVariableFromFunctionGroup = grammar.NewRule(SymbolVariable, SymbolFunctionGroup)
	PolymerizeBoolFromTrue              = grammar.NewRule(SymbolBool, SymbolTrue)
	PolymerizeBoolFromFalse             = grammar.NewRule(SymbolBool, SymbolFalse)
	PolymerizeVariableFromNull          = grammar.NewRule(SymbolVariable, SymbolNull)
	PolymerizeObject                    = grammar.NewRule(SymbolObject, SymbolLeftBrace, SymbolKeyValueList, SymbolRightBrace)
	PolymerizeFunctionGroup             = grammar.NewRule(SymbolFunctionGroup, SymbolFunction, SymbolLeftParenthesis, SymbolKeyList, SymbolRightParenthesis, SymbolKeyList, SymbolLeftBrace, SymbolExpressionList, SymbolRightBrace)
	PolymerizeIndexList                 = grammar.NewRule(SymbolIndexList, SymbolIndexArray)
	PolymerizeIndexListEmpty            = grammar.NewRule(SymbolIndexList)
	PolymerizeIndexArrayStart           = grammar.NewRule(SymbolIndexArray, SymbolIndex)
	PolymerizeIndexArray                = grammar.NewRule(SymbolIndexArray, SymbolIndexArray, SymbolComma, SymbolIndex)
	PolymerizeKeyValue                  = grammar.NewRule(SymbolKeyValue, SymbolIdentifier, SymbolColon, SymbolIndex)
	PolymerizeKeyValueList              = grammar.NewRule(SymbolKeyValueList, SymbolKeyValueArray)
	PolymerizeKeyValueListEmpty         = grammar.NewRule(SymbolKeyValueList)
	PolymerizeKeyValueArrayStart        = grammar.NewRule(SymbolKeyValueArray, SymbolKeyValue)
	PolymerizeKeyValueArray             = grammar.NewRule(SymbolKeyValueArray, SymbolKeyValueArray, SymbolComma, SymbolKeyValue)
	PolymerizeCallWithIndexArray        = grammar.NewRule(SymbolExpression2, SymbolExpression2, SymbolLeftParenthesis, SymbolIndexArray, SymbolRightParenthesis)
	PolymerizeCallWithKeyValueList      = grammar.NewRule(SymbolExpression2, SymbolExpression2, SymbolLeftParenthesis, SymbolKeyValueList, SymbolRightParenthesis)
	PolymerizeAssignment                = grammar.NewRule(SymbolExpression1, SymbolExpression2, SymbolEqual, SymbolExpression2)
	PolymerizeComponent                 = grammar.NewRule(SymbolExpression2, SymbolExpression2, SymbolDot, SymbolIdentifier)
	PolymerizeParentheses               = grammar.NewRule(SymbolExpression3, SymbolLeftParenthesis, SymbolExpression1, SymbolRightParenthesis)
	PolymerizeExpression3               = grammar.NewRule(SymbolExpression2, SymbolExpression3)
	PolymerizeExpression2               = grammar.NewRule(SymbolExpression1, SymbolExpression2)
	PolymerizeExpression1               = grammar.NewRule(SymbolExpression, SymbolExpression1, SymbolSemicolon)
	PolymerizeExpressionList            = grammar.NewRule(SymbolExpressionList, SymbolExpressionArray)
	PolymerizeExpressionListEmpty       = grammar.NewRule(SymbolExpressionList)
	PolymerizeExpressionArrayStart      = grammar.NewRule(SymbolExpressionArray, SymbolExpression)
	PolymerizeExpressionArray           = grammar.NewRule(SymbolExpressionArray, SymbolExpressionArray, SymbolExpression)
	PolymerizeKey                       = grammar.NewRule(SymbolKey, SymbolIdentifier)
	PolymerizeKeyList                   = grammar.NewRule(SymbolKeyList, SymbolKeyArray)
	PolymerizeKeyListEmpty              = grammar.NewRule(SymbolKeyList)
	PolymerizeKeyArrayStart             = grammar.NewRule(SymbolKeyArray, SymbolKey)
	PolymerizeKeyArray                  = grammar.NewRule(SymbolKeyArray, SymbolKeyArray, SymbolComma, SymbolKey)
)

var (
	GrammarRules = []*grammar.Rule{
		PolymerizePageGroup,
		PolymerizePageItemList,
		PolymerizePageItemListEmpty,
		PolymerizePageItemArrayStart,
		PolymerizePageItemArray,
		PolymerizePageItemFromImportGroup,
		PolymerizePageItemFromPublicGroup,
		PolymerizePageItemFromPrivateGroup,
		PolymerizeImportGroup,
		PolymerizePublicGroup,
		PolymerizePrivateGroup,
		PolymerizeIndexFromIdentifier,
		PolymerizeIndexFromVariable,
		PolymerizeIndexFromExpression,
		PolymerizeVariableFromNumber,
		PolymerizeVariableFromString,
		PolymerizeVariableFromBool,
		PolymerizeVariableFromObject,
		PolymerizeVariableFromFunctionGroup,
		PolymerizeBoolFromTrue,
		PolymerizeBoolFromFalse,
		PolymerizeVariableFromNull,
		PolymerizeObject,
		PolymerizeFunctionGroup,
		PolymerizeIndexList,
		PolymerizeIndexListEmpty,
		PolymerizeIndexArrayStart,
		PolymerizeIndexArray,
		PolymerizeKeyValue,
		PolymerizeKeyValueList,
		PolymerizeKeyValueListEmpty,
		PolymerizeKeyValueArrayStart,
		PolymerizeKeyValueArray,
		PolymerizeCallWithIndexArray,
		PolymerizeCallWithKeyValueList,
		PolymerizeAssignment,
		PolymerizeComponent,
		PolymerizeParentheses,
		PolymerizeExpression3,
		PolymerizeExpression2,
		PolymerizeExpression1,
		PolymerizeExpressionList,
		PolymerizeExpressionListEmpty,
		PolymerizeExpressionArrayStart,
		PolymerizeExpressionArray,
		PolymerizeKey,
		PolymerizeKeyList,
		PolymerizeKeyListEmpty,
		PolymerizeKeyArrayStart,
		PolymerizeKeyArray,
	}

	GrammarEnd = SymbolEnd

	GrammarGlobal = SymbolPageGroup
)
