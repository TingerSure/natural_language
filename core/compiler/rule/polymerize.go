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
	Space
	Colon
	Semicolon
	Dot
	Comma
	Number
	Identifier
	End
	String
	Import
	Export
	Class
	Require
	Provide
	Return
	Function
	Get
	Set
	===========nonterminal=========
	Page
	Sentence
	===========rule=========
	Page -> Sentence

	Sentence -> Sentence Sentence
	Sentence -> Import Identifier String Semicolon
*/

const (
	TypePage = iota
	TypeSentence
)

const (
	KeyPage     = "page"
	KeySentence = "sentence"
)

var (
	SymbolLeftParenthesis  = grammar.NewTerminal(TypeLeftParenthesis, KeyLeftParenthesis)
	SymbolRightParenthesis = grammar.NewTerminal(TypeRightParenthesis, KeyRightParenthesis)
	SymbolLeftBracket      = grammar.NewTerminal(TypeLeftBracket, KeyLeftBracket)
	SymbolRightBracket     = grammar.NewTerminal(TypeRightBracket, KeyRightBracket)
	SymbolLeftBrace        = grammar.NewTerminal(TypeLeftBrace, KeyLeftBrace)
	SymbolRightBrace       = grammar.NewTerminal(TypeRightBrace, KeyRightBrace)
	SymbolLeftArrow        = grammar.NewTerminal(TypeLeftArrow, KeyLeftArrow)
	SymbolSpace            = grammar.NewTerminal(TypeSpace, KeySpace)
	SymbolColon            = grammar.NewTerminal(TypeColon, KeyColon)
	SymbolSemicolon        = grammar.NewTerminal(TypeSemicolon, KeySemicolon)
	SymbolDot              = grammar.NewTerminal(TypeDot, KeyDot)
	SymbolComma            = grammar.NewTerminal(TypeComma, KeyComma)
	SymbolNumber           = grammar.NewTerminal(TypeNumber, KeyNumber)
	SymbolIdentifier       = grammar.NewTerminal(TypeIdentifier, KeyIdentifier)
	SymbolEnd              = grammar.NewTerminal(TypeEnd, KeyEnd)
	SymbolString           = grammar.NewTerminal(TypeString, KeyString)
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

var (
	SymbolPage     = grammar.NewNonterminal(TypePage, KeyPage)
	SymbolSentence = grammar.NewNonterminal(TypeSentence, KeySentence)
)

var (
	GrammarRules = []*grammar.Rule{
		grammar.NewRule(SymbolPage, SymbolSentence),
		grammar.NewRule(SymbolSentence, SymbolSentence, SymbolSentence),
		grammar.NewRule(SymbolSentence, SymbolImport, SymbolIdentifier, SymbolString, SymbolSemicolon),
	}
)
