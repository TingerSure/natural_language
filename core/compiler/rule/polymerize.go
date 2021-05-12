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
	Export
	Var
	Class
	Require
	Provide
	Return
	Function
	Get
	Set
	True
	False
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
	SymbolExport           = grammar.NewTerminal(TypeExport, KeyExport)
	SymbolVar              = grammar.NewTerminal(TypeVar, KeyVar)
	SymbolClass            = grammar.NewTerminal(TypeClass, KeyClass)
	SymbolRequire          = grammar.NewTerminal(TypeRequire, KeyRequire)
	SymbolProvide          = grammar.NewTerminal(TypeProvide, KeyProvide)
	SymbolReturn           = grammar.NewTerminal(TypeReturn, KeyReturn)
	SymbolFunction         = grammar.NewTerminal(TypeFunction, KeyFunction)
	SymbolGet              = grammar.NewTerminal(TypeGet, KeyGet)
	SymbolSet              = grammar.NewTerminal(TypeSet, KeySet)
	SymbolTrue             = grammar.NewTerminal(TypeTrue, KeyTrue)
	SymbolFalse            = grammar.NewTerminal(TypeFalse, KeyFalse)
)

const (
	TypePageGroup = iota
	TypePageItemList
	TypePageItemArray
	TypePageItem
	TypeImportGroup
	TypeExportGroup
	TypeVarGroup
	TypeIndex
	TypeIndexList
	TypeIndexArray
	TypeKeyValue
	TypeKeyValueList
	TypeKeyValueArray
	TypeBool
	TypeObject
	TypeVariable
)

const (
	KeyPageGroup     = "page_group"
	KeyPageItemList  = "page_item_list"
	KeyPageItemArray = "page_item_array"
	KeyPageItem      = "page_item"
	KeyImportGroup   = "import_group"
	KeyExportGroup   = "export_group"
	KeyVarGroup      = "var_group"
	KeyIndex         = "index"
	KeyIndexList     = "index_list"
	KeyIndexArray    = "index_array"
	KeyKeyValue      = "key_value"
	KeyKeyValueList  = "key_value_list"
	KeyKeyValueArray = "key_value_array"
	KeyBool          = "bool"
	KeyObject        = "object"
	KeyVariable      = "variable"
)

var (
	SymbolPageGroup     = grammar.NewNonterminal(TypePageGroup, KeyPageGroup)
	SymbolPageItemList  = grammar.NewNonterminal(TypePageItemList, KeyPageItemList)
	SymbolPageItemArray = grammar.NewNonterminal(TypePageItemArray, KeyPageItemArray)
	SymbolPageItem      = grammar.NewNonterminal(TypePageItem, KeyPageItem)
	SymbolImportGroup   = grammar.NewNonterminal(TypeImportGroup, KeyImportGroup)
	SymbolExportGroup   = grammar.NewNonterminal(TypeExportGroup, KeyExportGroup)
	SymbolVarGroup      = grammar.NewNonterminal(TypeVarGroup, KeyVarGroup)
	SymbolIndex         = grammar.NewNonterminal(TypeIndex, KeyIndex)
	SymbolIndexList     = grammar.NewNonterminal(TypeIndexList, KeyIndexList)
	SymbolIndexArray    = grammar.NewNonterminal(TypeIndexArray, KeyIndexArray)
	SymbolKeyValue      = grammar.NewNonterminal(TypeKeyValue, KeyKeyValue)
	SymbolKeyValueList  = grammar.NewNonterminal(TypeKeyValueList, KeyKeyValueList)
	SymbolKeyValueArray = grammar.NewNonterminal(TypeKeyValueArray, KeyKeyValueArray)
	SymbolBool          = grammar.NewNonterminal(TypeBool, KeyBool)
	SymbolObject        = grammar.NewNonterminal(TypeObject, KeyObject)
	SymbolVariable      = grammar.NewNonterminal(TypeVariable, KeyVariable)
)

var (
	PolymerizePageGroup               = grammar.NewRule(SymbolPageGroup, SymbolPageItemList)
	PolymerizePageItemList            = grammar.NewRule(SymbolPageItemList, SymbolPageItemArray)
	PolymerizePageItemListEmpty       = grammar.NewRule(SymbolPageItemList)
	PolymerizePageItemArrayStart      = grammar.NewRule(SymbolPageItemArray, SymbolPageItem)
	PolymerizePageItemArray           = grammar.NewRule(SymbolPageItemArray, SymbolPageItemArray, SymbolPageItem)
	PolymerizePageItemFromImportGroup = grammar.NewRule(SymbolPageItem, SymbolImportGroup)
	PolymerizePageItemFromExportGroup = grammar.NewRule(SymbolPageItem, SymbolExportGroup)
	PolymerizePageItemFromVarGroup    = grammar.NewRule(SymbolPageItem, SymbolVarGroup)
	PolymerizeImportGroup             = grammar.NewRule(SymbolImportGroup, SymbolImport, SymbolIdentifier, SymbolString)
	PolymerizeExportGroup             = grammar.NewRule(SymbolExportGroup, SymbolExport, SymbolIdentifier, SymbolIndex)
	PolymerizeVarGroup                = grammar.NewRule(SymbolVarGroup, SymbolVar, SymbolIdentifier, SymbolIndex)
	PolymerizeIndexFromIdentifier     = grammar.NewRule(SymbolIndex, SymbolIdentifier)
	PolymerizeIndexFromVariable       = grammar.NewRule(SymbolIndex, SymbolVariable)
	PolymerizeVariableFromNumber      = grammar.NewRule(SymbolVariable, SymbolNumber)
	PolymerizeVariableFromString      = grammar.NewRule(SymbolVariable, SymbolString)
	PolymerizeVariableFromBool        = grammar.NewRule(SymbolVariable, SymbolBool)
	PolymerizeVariableFromObject      = grammar.NewRule(SymbolVariable, SymbolObject)
	PolymerizeBoolFromTrue            = grammar.NewRule(SymbolBool, SymbolTrue)
	PolymerizeBoolFromFalse           = grammar.NewRule(SymbolBool, SymbolFalse)
	PolymerizeIndexList               = grammar.NewRule(SymbolIndexList, SymbolIndexArray)
	PolymerizeIndexListEmpty          = grammar.NewRule(SymbolIndexList)
	PolymerizeIndexArrayStart         = grammar.NewRule(SymbolIndexArray, SymbolIndex)
	PolymerizeIndexArray              = grammar.NewRule(SymbolIndexArray, SymbolIndexArray, SymbolComma, SymbolIndex)
	PolymerizeKeyValue                = grammar.NewRule(SymbolKeyValue, SymbolIdentifier, SymbolColon, SymbolIndex)
	PolymerizeKeyValueList            = grammar.NewRule(SymbolKeyValueList, SymbolKeyValueArray)
	PolymerizeKeyValueListEmpty       = grammar.NewRule(SymbolKeyValueList)
	PolymerizeKeyValueArrayStart      = grammar.NewRule(SymbolKeyValueArray, SymbolKeyValue)
	PolymerizeKeyValueArray           = grammar.NewRule(SymbolKeyValueArray, SymbolKeyValueArray, SymbolComma, SymbolKeyValue)
	PolymerizeCallWithIndexArray      = grammar.NewRule(SymbolIndex, SymbolIndex, SymbolLeftParenthesis, SymbolIndexArray, SymbolRightParenthesis)
	PolymerizeCallWithKeyValueList    = grammar.NewRule(SymbolIndex, SymbolIndex, SymbolLeftParenthesis, SymbolKeyValueList, SymbolRightParenthesis)
	PolymerizeComponent               = grammar.NewRule(SymbolIndex, SymbolIndex, SymbolDot, SymbolIdentifier)
	PolymerizeObjectWithKeyValueList  = grammar.NewRule(SymbolObject, SymbolLeftBrace, SymbolKeyValueList, SymbolRightBrace)
)

var (
	GrammarRules = []*grammar.Rule{
		PolymerizePageGroup,
		PolymerizePageItemList,
		PolymerizePageItemListEmpty,
		PolymerizePageItemArrayStart,
		PolymerizePageItemArray,
		PolymerizePageItemFromImportGroup,
		PolymerizePageItemFromExportGroup,
		PolymerizePageItemFromVarGroup,
		PolymerizeImportGroup,
		PolymerizeExportGroup,
		PolymerizeVarGroup,
		PolymerizeIndexFromIdentifier,
		PolymerizeIndexFromVariable,
		PolymerizeVariableFromNumber,
		PolymerizeVariableFromString,
		PolymerizeVariableFromBool,
		PolymerizeVariableFromObject,
		PolymerizeBoolFromTrue,
		PolymerizeBoolFromFalse,
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
		PolymerizeComponent,
		PolymerizeObjectWithKeyValueList,
	}

	GrammarEnd = SymbolEnd

	GrammarGlobal = SymbolPageGroup
)
