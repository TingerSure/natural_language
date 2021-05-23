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
	Eof
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
	Var
	If
	Else
	For
	Continue
	Break
	This
	Self
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
	SymbolEof              = grammar.NewTerminal(TypeEof, KeyEof)
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
	SymbolVar              = grammar.NewTerminal(TypeVar, KeyVar)
	SymbolIf               = grammar.NewTerminal(TypeIf, KeyIf)
	SymbolElse             = grammar.NewTerminal(TypeElse, KeyElse)
	SymbolFor              = grammar.NewTerminal(TypeFor, KeyFor)
	SymbolContinue         = grammar.NewTerminal(TypeContinue, KeyContinue)
	SymbolBreak            = grammar.NewTerminal(TypeBreak, KeyBreak)
	SymbolThis             = grammar.NewTerminal(TypeThis, KeyThis)
	SymbolSelf             = grammar.NewTerminal(TypeSelf, KeySelf)
)

const (
	TypePageGroup = iota
	TypePageItemList
	TypePageItemArray
	TypePageItem
	TypeClassGroup
	TypeClassItemList
	TypeClassItemArray
	TypeClassItem
	TypeImportGroup
	TypePublicGroup
	TypePrivateGroup
	TypeProvideGroup
	TypeRequireGroup
	TypeIndex
	TypeIndexList
	TypeIndexArray
	TypeKeyValue
	TypeKeyValueList
	TypeKeyValueArray
	TypeKeyKey
	TypeKeyKeyList
	TypeKeyKeyArray
	TypeBool
	TypeObject
	TypeFunctionGroup
	TypeDefineFunctionGroup
	TypeVariable
	TypeExpression
	TypeExpressionIndependent
	TypeExpressionIndependentList
	TypeExpressionIndependentArray
	TypeExpressionCeil
	TypeExpression1
	TypeExpressionFloor
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
	KeyPageGroup                  = "page_group"
	KeyPageItemList               = "page_item_list"
	KeyPageItemArray              = "page_item_array"
	KeyPageItem                   = "page_item"
	KeyClassGroup                 = "class_group"
	KeyClassItemList              = "class_item_list"
	KeyClassItemArray             = "class_item_array"
	KeyClassItem                  = "class_item"
	KeyImportGroup                = "import_group"
	KeyPublicGroup                = "public_group"
	KeyPrivateGroup               = "private_group"
	KeyProvideGroup               = "provide_group"
	KeyRequireGroup               = "require_group"
	KeyIndex                      = "index"
	KeyIndexList                  = "index_list"
	KeyIndexArray                 = "index_array"
	KeyKeyValue                   = "key_value"
	KeyKeyValueList               = "key_value_list"
	KeyKeyValueArray              = "key_value_array"
	KeyKeyKey                     = "key_key"
	KeyKeyKeyList                 = "key_key_list"
	KeyKeyKeyArray                = "key_key_array"
	KeyBool                       = "bool"
	KeyObject                     = "object"
	KeyFunctionGroup              = "function_group"
	KeyDefineFunctionGroup        = "define_function_group"
	KeyVariable                   = "variable"
	KeyExpression                 = "expression"
	KeyExpressionIndependent      = "expression_independent"
	KeyExpressionIndependentList  = "expression_independent_list"
	KeyExpressionIndependentArray = "expression_independent_array"
	KeyExpressionCeil             = "expression_1"
	KeyExpression1                = "expression_2"
	KeyExpressionFloor            = "expression_3"
	KeyExpressionList             = "expression_list"
	KeyExpressionArray            = "expression_array"
	KeyKey                        = "param"
	KeyKeyList                    = "param_list"
	KeyKeyArray                   = "param_array"
)

var (
	SymbolPageGroup                  = grammar.NewNonterminal(TypePageGroup, KeyPageGroup)
	SymbolPageItemList               = grammar.NewNonterminal(TypePageItemList, KeyPageItemList)
	SymbolPageItemArray              = grammar.NewNonterminal(TypePageItemArray, KeyPageItemArray)
	SymbolPageItem                   = grammar.NewNonterminal(TypePageItem, KeyPageItem)
	SymbolClassGroup                 = grammar.NewNonterminal(TypeClassGroup, KeyClassGroup)
	SymbolClassItemList              = grammar.NewNonterminal(TypeClassItemList, KeyClassItemList)
	SymbolClassItemArray             = grammar.NewNonterminal(TypeClassItemArray, KeyClassItemArray)
	SymbolClassItem                  = grammar.NewNonterminal(TypeClassItem, KeyClassItem)
	SymbolImportGroup                = grammar.NewNonterminal(TypeImportGroup, KeyImportGroup)
	SymbolPublicGroup                = grammar.NewNonterminal(TypePublicGroup, KeyPublicGroup)
	SymbolPrivateGroup               = grammar.NewNonterminal(TypePrivateGroup, KeyPrivateGroup)
	SymbolProvideGroup               = grammar.NewNonterminal(TypeProvideGroup, KeyProvideGroup)
	SymbolRequireGroup               = grammar.NewNonterminal(TypeRequireGroup, KeyRequireGroup)
	SymbolIndex                      = grammar.NewNonterminal(TypeIndex, KeyIndex)
	SymbolIndexList                  = grammar.NewNonterminal(TypeIndexList, KeyIndexList)
	SymbolIndexArray                 = grammar.NewNonterminal(TypeIndexArray, KeyIndexArray)
	SymbolKeyValue                   = grammar.NewNonterminal(TypeKeyValue, KeyKeyValue)
	SymbolKeyValueList               = grammar.NewNonterminal(TypeKeyValueList, KeyKeyValueList)
	SymbolKeyValueArray              = grammar.NewNonterminal(TypeKeyValueArray, KeyKeyValueArray)
	SymbolKeyKey                     = grammar.NewNonterminal(TypeKeyKey, KeyKeyKey)
	SymbolKeyKeyList                 = grammar.NewNonterminal(TypeKeyKeyList, KeyKeyKeyList)
	SymbolKeyKeyArray                = grammar.NewNonterminal(TypeKeyKeyArray, KeyKeyKeyArray)
	SymbolBool                       = grammar.NewNonterminal(TypeBool, KeyBool)
	SymbolObject                     = grammar.NewNonterminal(TypeObject, KeyObject)
	SymbolFunctionGroup              = grammar.NewNonterminal(TypeFunctionGroup, KeyFunctionGroup)
	SymbolDefineFunctionGroup        = grammar.NewNonterminal(TypeDefineFunctionGroup, KeyDefineFunctionGroup)
	SymbolVariable                   = grammar.NewNonterminal(TypeVariable, KeyVariable)
	SymbolExpression                 = grammar.NewNonterminal(TypeExpression, KeyExpression)
	SymbolExpressionIndependent      = grammar.NewNonterminal(TypeExpressionIndependent, KeyExpressionIndependent)
	SymbolExpressionIndependentList  = grammar.NewNonterminal(TypeExpressionIndependentList, KeyExpressionIndependentList)
	SymbolExpressionIndependentArray = grammar.NewNonterminal(TypeExpressionIndependentArray, KeyExpressionIndependentArray)
	SymbolExpressionCeil             = grammar.NewNonterminal(TypeExpressionCeil, KeyExpressionCeil)
	SymbolExpression1                = grammar.NewNonterminal(TypeExpression1, KeyExpression1)
	SymbolExpressionFloor            = grammar.NewNonterminal(TypeExpressionFloor, KeyExpressionFloor)
	SymbolExpressionList             = grammar.NewNonterminal(TypeExpressionList, KeyExpressionList)
	SymbolExpressionArray            = grammar.NewNonterminal(TypeExpressionArray, KeyExpressionArray)
	SymbolKey                        = grammar.NewNonterminal(TypeKey, KeyKey)
	SymbolKeyList                    = grammar.NewNonterminal(TypeKeyList, KeyKeyList)
	SymbolKeyArray                   = grammar.NewNonterminal(TypeKeyArray, KeyKeyArray)
)

var (
	PolymerizePageGroup                       = grammar.NewRule(SymbolPageGroup, SymbolPageItemList)
	PolymerizePageItemList                    = grammar.NewRule(SymbolPageItemList, SymbolPageItemArray)
	PolymerizePageItemListEmpty               = grammar.NewRule(SymbolPageItemList)
	PolymerizePageItemArrayStart              = grammar.NewRule(SymbolPageItemArray, SymbolPageItem)
	PolymerizePageItemArray                   = grammar.NewRule(SymbolPageItemArray, SymbolPageItemArray, SymbolPageItem)
	PolymerizeClassGroup                      = grammar.NewRule(SymbolClassGroup, SymbolClass, SymbolLeftBrace, SymbolClassItemList, SymbolRightBrace)
	PolymerizeClassItemList                   = grammar.NewRule(SymbolClassItemList, SymbolClassItemArray)
	PolymerizeClassItemListEmpty              = grammar.NewRule(SymbolClassItemList)
	PolymerizeClassItemArrayStart             = grammar.NewRule(SymbolClassItemArray, SymbolClassItem)
	PolymerizeClassItemArray                  = grammar.NewRule(SymbolClassItemArray, SymbolClassItemArray, SymbolClassItem)
	PolymerizePageItemFromImportGroup         = grammar.NewRule(SymbolPageItem, SymbolImportGroup)
	PolymerizePageItemFromPublicGroup         = grammar.NewRule(SymbolPageItem, SymbolPublicGroup)
	PolymerizePageItemFromPrivateGroup        = grammar.NewRule(SymbolPageItem, SymbolPrivateGroup)
	PolymerizeClassItemFromProvideGroup       = grammar.NewRule(SymbolClassItem, SymbolProvideGroup)
	PolymerizeClassItemFromRequireGroup       = grammar.NewRule(SymbolClassItem, SymbolRequireGroup)
	PolymerizeImportGroup                     = grammar.NewRule(SymbolImportGroup, SymbolImport, SymbolIdentifier, SymbolString)
	PolymerizePublicGroup                     = grammar.NewRule(SymbolPublicGroup, SymbolPublic, SymbolIdentifier, SymbolEqual, SymbolIndex)
	PolymerizePrivateGroup                    = grammar.NewRule(SymbolPrivateGroup, SymbolPrivate, SymbolIdentifier, SymbolEqual, SymbolIndex)
	PolymerizeProvideGroup                    = grammar.NewRule(SymbolProvideGroup, SymbolProvide, SymbolIdentifier, SymbolEqual, SymbolIndex)
	PolymerizeRequireGroup                    = grammar.NewRule(SymbolRequireGroup, SymbolRequire, SymbolIdentifier, SymbolEqual, SymbolDefineFunctionGroup)
	PolymerizeExpressionFromIdentifier        = grammar.NewRule(SymbolExpressionFloor, SymbolIdentifier)
	PolymerizeExpressionFromVariable          = grammar.NewRule(SymbolExpressionFloor, SymbolVariable)
	PolymerizeIndexFromExpression             = grammar.NewRule(SymbolIndex, SymbolExpressionCeil)
	PolymerizeVariableFromNumber              = grammar.NewRule(SymbolVariable, SymbolNumber)
	PolymerizeVariableFromString              = grammar.NewRule(SymbolVariable, SymbolString)
	PolymerizeVariableFromBool                = grammar.NewRule(SymbolVariable, SymbolBool)
	PolymerizeVariableFromObject              = grammar.NewRule(SymbolVariable, SymbolObject)
	PolymerizeVariableFromFunctionGroup       = grammar.NewRule(SymbolVariable, SymbolFunctionGroup)
	PolymerizeVariableFromClassGroup          = grammar.NewRule(SymbolVariable, SymbolClassGroup)
	PolymerizeBoolFromTrue                    = grammar.NewRule(SymbolBool, SymbolTrue)
	PolymerizeBoolFromFalse                   = grammar.NewRule(SymbolBool, SymbolFalse)
	PolymerizeVariableFromNull                = grammar.NewRule(SymbolVariable, SymbolNull)
	PolymerizeObject                          = grammar.NewRule(SymbolObject, SymbolLeftBrace, SymbolKeyValueList, SymbolRightBrace)
	PolymerizeFunctionGroup                   = grammar.NewRule(SymbolFunctionGroup, SymbolFunction, SymbolLeftParenthesis, SymbolKeyList, SymbolRightParenthesis, SymbolKeyList, SymbolLeftBrace, SymbolExpressionList, SymbolRightBrace)
	PolymerizeDefineFunctionGroup             = grammar.NewRule(SymbolDefineFunctionGroup, SymbolFunction, SymbolLeftParenthesis, SymbolKeyList, SymbolRightParenthesis, SymbolKeyList)
	PolymerizeIndexList                       = grammar.NewRule(SymbolIndexList, SymbolIndexArray)
	PolymerizeIndexListEmpty                  = grammar.NewRule(SymbolIndexList)
	PolymerizeIndexArrayStart                 = grammar.NewRule(SymbolIndexArray, SymbolIndex)
	PolymerizeIndexArray                      = grammar.NewRule(SymbolIndexArray, SymbolIndexArray, SymbolComma, SymbolIndex)
	PolymerizeKeyValue                        = grammar.NewRule(SymbolKeyValue, SymbolIdentifier, SymbolColon, SymbolIndex)
	PolymerizeKeyValueList                    = grammar.NewRule(SymbolKeyValueList, SymbolKeyValueArray)
	PolymerizeKeyValueListEmpty               = grammar.NewRule(SymbolKeyValueList)
	PolymerizeKeyValueArrayStart              = grammar.NewRule(SymbolKeyValueArray, SymbolKeyValue)
	PolymerizeKeyValueArray                   = grammar.NewRule(SymbolKeyValueArray, SymbolKeyValueArray, SymbolComma, SymbolKeyValue)
	PolymerizeKeyKey                          = grammar.NewRule(SymbolKeyKey, SymbolIdentifier, SymbolColon, SymbolIdentifier)
	PolymerizeKeyKeyList                      = grammar.NewRule(SymbolKeyKeyList, SymbolKeyKeyArray)
	PolymerizeKeyKeyListEmpty                 = grammar.NewRule(SymbolKeyKeyList)
	PolymerizeKeyKeyArrayStart                = grammar.NewRule(SymbolKeyKeyArray, SymbolKeyKey)
	PolymerizeKeyKeyArray                     = grammar.NewRule(SymbolKeyKeyArray, SymbolKeyKeyArray, SymbolComma, SymbolKeyKey)
	PolymerizeMappingObject                   = grammar.NewRule(SymbolExpressionCeil, SymbolExpressionCeil, SymbolRightArrow, SymbolExpression1, SymbolLeftBrace, SymbolKeyKeyList, SymbolRightBrace)
	PolymerizeMappingObjectAuto               = grammar.NewRule(SymbolExpressionCeil, SymbolExpressionCeil, SymbolRightArrow, SymbolExpression1)
	PolymerizeCallWithIndexArray              = grammar.NewRule(SymbolExpressionFloor, SymbolExpression1, SymbolLeftParenthesis, SymbolIndexArray, SymbolRightParenthesis)
	PolymerizeCallWithKeyValueList            = grammar.NewRule(SymbolExpressionFloor, SymbolExpression1, SymbolLeftParenthesis, SymbolKeyValueList, SymbolRightParenthesis)
	PolymerizeAssignment                      = grammar.NewRule(SymbolExpressionIndependent, SymbolExpressionCeil, SymbolEqual, SymbolExpressionCeil)
	PolymerizeComponent                       = grammar.NewRule(SymbolExpressionFloor, SymbolExpressionFloor, SymbolDot, SymbolIdentifier)
	PolymerizeDefine                          = grammar.NewRule(SymbolExpressionIndependent, SymbolVar, SymbolIdentifier)
	PolymerizeDefineAndInit                   = grammar.NewRule(SymbolExpressionIndependent, SymbolVar, SymbolIdentifier, SymbolEqual, SymbolExpressionCeil)
	PolymerizeParentheses                     = grammar.NewRule(SymbolExpressionFloor, SymbolLeftParenthesis, SymbolExpressionCeil, SymbolRightParenthesis)
	PolymerizeIf                              = grammar.NewRule(SymbolExpressionIndependent, SymbolIf, SymbolLeftParenthesis, SymbolExpressionCeil, SymbolRightParenthesis, SymbolLeftBrace, SymbolExpressionList, SymbolRightBrace)
	PolymerizeIfElse                          = grammar.NewRule(SymbolExpressionIndependent, SymbolIf, SymbolLeftParenthesis, SymbolExpressionCeil, SymbolRightParenthesis, SymbolLeftBrace, SymbolExpressionList, SymbolRightBrace, SymbolElse, SymbolLeftBrace, SymbolExpressionList, SymbolRightBrace)
	PolymerizeFor                             = grammar.NewRule(SymbolExpressionIndependent, SymbolFor, SymbolLeftParenthesis, SymbolExpressionIndependentList, SymbolSemicolon, SymbolExpressionCeil, SymbolSemicolon, SymbolExpressionIndependentList, SymbolRightParenthesis, SymbolLeftBrace, SymbolExpressionList, SymbolRightBrace)
	PolymerizeWhile                           = grammar.NewRule(SymbolExpressionIndependent, SymbolFor, SymbolLeftParenthesis, SymbolExpressionCeil, SymbolRightParenthesis, SymbolLeftBrace, SymbolExpressionList, SymbolRightBrace)
	PolymerizeExpression1FromFloor            = grammar.NewRule(SymbolExpression1, SymbolExpressionFloor)
	PolymerizeExpressionCeilFrom1             = grammar.NewRule(SymbolExpressionCeil, SymbolExpression1)
	PolymerizeExpressionIndependentFromCeil   = grammar.NewRule(SymbolExpressionIndependent, SymbolExpressionCeil)
	PolymerizeExpression                      = grammar.NewRule(SymbolExpression, SymbolExpressionIndependent, SymbolSemicolon)
	PolymerizeExpressionIndependentList       = grammar.NewRule(SymbolExpressionIndependentList, SymbolExpressionIndependentArray)
	PolymerizeExpressionIndependentListEmpty  = grammar.NewRule(SymbolExpressionIndependentList)
	PolymerizeExpressionIndependentArrayStart = grammar.NewRule(SymbolExpressionIndependentArray, SymbolExpressionIndependent)
	PolymerizeExpressionIndependentArray      = grammar.NewRule(SymbolExpressionIndependentArray, SymbolExpressionIndependentArray, SymbolComma, SymbolExpressionIndependent)
	PolymerizeExpressionList                  = grammar.NewRule(SymbolExpressionList, SymbolExpressionArray)
	PolymerizeExpressionListEmpty             = grammar.NewRule(SymbolExpressionList)
	PolymerizeExpressionArrayStart            = grammar.NewRule(SymbolExpressionArray, SymbolExpression)
	PolymerizeExpressionArray                 = grammar.NewRule(SymbolExpressionArray, SymbolExpressionArray, SymbolExpression)
	PolymerizeKey                             = grammar.NewRule(SymbolKey, SymbolIdentifier)
	PolymerizeKeyList                         = grammar.NewRule(SymbolKeyList, SymbolKeyArray)
	PolymerizeKeyListEmpty                    = grammar.NewRule(SymbolKeyList)
	PolymerizeKeyArrayStart                   = grammar.NewRule(SymbolKeyArray, SymbolKey)
	PolymerizeKeyArray                        = grammar.NewRule(SymbolKeyArray, SymbolKeyArray, SymbolComma, SymbolKey)
	PolymerizeContinue                        = grammar.NewRule(SymbolExpressionIndependent, SymbolContinue)
	PolymerizeContinueTag                     = grammar.NewRule(SymbolExpressionIndependent, SymbolContinue, SymbolIdentifier)
	PolymerizeBreak                           = grammar.NewRule(SymbolExpressionIndependent, SymbolBreak)
	PolymerizeBreakTag                        = grammar.NewRule(SymbolExpressionIndependent, SymbolBreak, SymbolIdentifier)
	PolymerizeReturn                          = grammar.NewRule(SymbolExpressionIndependent, SymbolReturn)
	PolymerizeThis                            = grammar.NewRule(SymbolExpressionFloor, SymbolThis)
	PolymerizeSelf                            = grammar.NewRule(SymbolExpressionFloor, SymbolSelf)
)

var (
	GrammarRules = []*grammar.Rule{
		PolymerizePageGroup,
		PolymerizePageItemList,
		PolymerizePageItemListEmpty,
		PolymerizePageItemArrayStart,
		PolymerizePageItemArray,
		PolymerizeClassGroup,
		PolymerizeClassItemList,
		PolymerizeClassItemListEmpty,
		PolymerizeClassItemArrayStart,
		PolymerizeClassItemArray,
		PolymerizePageItemFromImportGroup,
		PolymerizePageItemFromPublicGroup,
		PolymerizePageItemFromPrivateGroup,
		PolymerizeClassItemFromProvideGroup,
		PolymerizeClassItemFromRequireGroup,
		PolymerizeImportGroup,
		PolymerizePublicGroup,
		PolymerizePrivateGroup,
		PolymerizeProvideGroup,
		PolymerizeRequireGroup,
		PolymerizeExpressionFromIdentifier,
		PolymerizeExpressionFromVariable,
		PolymerizeIndexFromExpression,
		PolymerizeVariableFromNumber,
		PolymerizeVariableFromString,
		PolymerizeVariableFromBool,
		PolymerizeVariableFromObject,
		PolymerizeVariableFromFunctionGroup,
		PolymerizeVariableFromClassGroup,
		PolymerizeBoolFromTrue,
		PolymerizeBoolFromFalse,
		PolymerizeVariableFromNull,
		PolymerizeObject,
		PolymerizeFunctionGroup,
		PolymerizeDefineFunctionGroup,
		PolymerizeIndexList,
		PolymerizeIndexListEmpty,
		PolymerizeIndexArrayStart,
		PolymerizeIndexArray,
		PolymerizeKeyValue,
		PolymerizeKeyValueList,
		PolymerizeKeyValueListEmpty,
		PolymerizeKeyValueArrayStart,
		PolymerizeKeyValueArray,
		PolymerizeKeyKey,
		PolymerizeKeyKeyList,
		PolymerizeKeyKeyListEmpty,
		PolymerizeKeyKeyArrayStart,
		PolymerizeKeyKeyArray,
		PolymerizeMappingObject,
		PolymerizeMappingObjectAuto,
		PolymerizeCallWithIndexArray,
		PolymerizeCallWithKeyValueList,
		PolymerizeAssignment,
		PolymerizeComponent,
		PolymerizeDefine,
		PolymerizeDefineAndInit,
		PolymerizeParentheses,
		PolymerizeIf,
		PolymerizeIfElse,
		PolymerizeFor,
		PolymerizeWhile,
		PolymerizeExpression1FromFloor,
		PolymerizeExpressionCeilFrom1,
		PolymerizeExpressionIndependentFromCeil,
		PolymerizeExpression,
		PolymerizeExpressionIndependentList,
		PolymerizeExpressionIndependentListEmpty,
		PolymerizeExpressionIndependentArrayStart,
		PolymerizeExpressionIndependentArray,
		PolymerizeExpressionList,
		PolymerizeExpressionListEmpty,
		PolymerizeExpressionArrayStart,
		PolymerizeExpressionArray,
		PolymerizeKey,
		PolymerizeKeyList,
		PolymerizeKeyListEmpty,
		PolymerizeKeyArrayStart,
		PolymerizeKeyArray,
		PolymerizeContinue,
		PolymerizeContinueTag,
		PolymerizeBreak,
		PolymerizeBreakTag,
		PolymerizeReturn,
		PolymerizeThis,
		PolymerizeSelf,
	}

	GrammarEof = SymbolEof

	GrammarGlobal = SymbolPageGroup
)
