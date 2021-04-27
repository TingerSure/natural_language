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
	Class
	Require
	Provide
	Return
	Function
	Get
	Set
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
	SymbolPage             = grammar.NewTerminal(TypePage, KeyPage)
	SymbolImport           = grammar.NewTerminal(TypeImport, KeyImport)
	SymbolExport           = grammar.NewTerminal(TypeExport, KeyExport)
	SymbolClass            = grammar.NewTerminal(TypeClass, KeyClass)
	SymbolRequire          = grammar.NewTerminal(TypeRequire, KeyRequire)
	SymbolProvide          = grammar.NewTerminal(TypeProvide, KeyProvide)
	SymbolReturn           = grammar.NewTerminal(TypeReturn, KeyReturn)
	SymbolFunction         = grammar.NewTerminal(TypeFunction, KeyFunction)
	SymbolGet              = grammar.NewTerminal(TypeGet, KeyGet)
	SymbolSet              = grammar.NewTerminal(TypeSet, KeySet)
)

const (
	TypePageGroup = iota
	TypePageItemArray
	TypePageItem
	TypeImportGroup
	TypeExportGroup
	TypeIndex
)

const (
	KeyPageGroup     = "page_group"
	KeyPageItemArray = "page_item_array"
	KeyPageItem      = "page_item"
	KeyImportGroup   = "import_group"
	KeyExportGroup   = "export_group"
	KeyIndex         = "index"
)

var (
	SymbolPageGroup     = grammar.NewNonterminal(TypePageGroup, KeyPageGroup)
	SymbolPageItemArray = grammar.NewNonterminal(TypePageItemArray, KeyPageItemArray)
	SymbolPageItem      = grammar.NewNonterminal(TypePageItem, KeyPageItem)
	SymbolImportGroup   = grammar.NewNonterminal(TypeImportGroup, KeyImportGroup)
	SymbolExportGroup   = grammar.NewNonterminal(TypeExportGroup, KeyExportGroup)
	SymbolIndex         = grammar.NewNonterminal(TypeIndex, KeyIndex)
)

var (
	PolymerizePageGroup               = grammar.NewRule(SymbolPageGroup, SymbolPageItemArray)
	PolymerizePageItemArrayStart      = grammar.NewRule(SymbolPageItemArray, SymbolPageItem)
	PolymerizePageItemArray           = grammar.NewRule(SymbolPageItemArray, SymbolPageItemArray, SymbolPageItem)
	PolymerizePageItemFromImportGroup = grammar.NewRule(SymbolPageItem, SymbolImportGroup)
	PolymerizeImportGroup             = grammar.NewRule(SymbolImportGroup, SymbolImport, SymbolIdentifier, SymbolString)
	PolymerizeExportGroup             = grammar.NewRule(SymbolExportGroup, SymbolExport, SymbolIdentifier, SymbolIndex)
	PolymerizeIndexFromNumber         = grammar.NewRule(SymbolIndex, SymbolNumber)
	PolymerizeIndexFromString         = grammar.NewRule(SymbolIndex, SymbolString)
	// PolymerizeIndexFromBool         = grammar.NewRule(SymbolIndex, SymbolBool)
)

var (
	GrammarRules = []*grammar.Rule{
		PolymerizePageGroup,
		PolymerizePageItemArrayStart,
		PolymerizePageItemArray,
		PolymerizePageItemFromImportGroup,
		PolymerizeImportGroup,
		PolymerizeExportGroup,
		PolymerizeIndexFromNumber,
		PolymerizeIndexFromString,
	}

	GrammarEnd = SymbolEnd

	GrammarGlobal = SymbolPageGroup
)
