package manager

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

const (
	LeftParenthesis  = iota // (
	RightParenthesis        // )
	LeftBracket             // [
	RightBracket            // ]
	Leftbrace               // {
	Rightbrace              // }
	LeftArrow               // <-
	Space                   // \r \n \t [:space:]
	Colon                   // :
	Semicolon               // ;
	Dot                     // .
	Comma                   // ,
	Number                  // [:number:]
	String                  // [:string:]
	Identifier              // [:identifier:]
	End                     // [:end:]
)

var (
	LexerRules = []*lexer.Rule{
		lexer.NewRule("\\(", func(value []byte) *lexer.Token {
			return lexer.NewToken(LeftParenthesis, string(value))
		}),
		lexer.NewRule("\\)", func(value []byte) *lexer.Token {
			return lexer.NewToken(RightParenthesis, string(value))
		}),
		lexer.NewRule("\\[", func(value []byte) *lexer.Token {
			return lexer.NewToken(LeftBracket, string(value))
		}),
		lexer.NewRule("\\]", func(value []byte) *lexer.Token {
			return lexer.NewToken(RightBracket, string(value))
		}),
		lexer.NewRule("\\{", func(value []byte) *lexer.Token {
			return lexer.NewToken(Leftbrace, string(value))
		}),
		lexer.NewRule("\\}", func(value []byte) *lexer.Token {
			return lexer.NewToken(Rightbrace, string(value))
		}),
		lexer.NewRule("<\\-", func(value []byte) *lexer.Token {
			return lexer.NewToken(LeftArrow, string(value))
		}),
		lexer.NewRule(":", func(value []byte) *lexer.Token {
			return lexer.NewToken(Colon, string(value))
		}),
		lexer.NewRule(";", func(value []byte) *lexer.Token {
			return lexer.NewToken(Semicolon, string(value))
		}),
		lexer.NewRule("\\.", func(value []byte) *lexer.Token {
			return lexer.NewToken(Dot, string(value))
		}),
		lexer.NewRule("\\,", func(value []byte) *lexer.Token {
			return lexer.NewToken(Comma, string(value))
		}),
		lexer.NewRule("\\n", func(value []byte) *lexer.Token {
			return lexer.NewToken(Space, string(value))
		}),
		lexer.NewRule("\\r", func(value []byte) *lexer.Token {
			return lexer.NewToken(Space, string(value))
		}),
		lexer.NewRule("\\t", func(value []byte) *lexer.Token {
			return lexer.NewToken(Space, string(value))
		}),
		lexer.NewRule(" ", func(value []byte) *lexer.Token {
			return lexer.NewToken(Space, string(value))
		}),
		lexer.NewRule("([\\+|-]?\\d+)(\\.\\d+)?(E[\\+|-]?\\d+)?", func(value []byte) *lexer.Token {
			return lexer.NewToken(Number, string(value))
		}),
		lexer.NewRule("\"\\S*\"", func(value []byte) *lexer.Token {
			return lexer.NewToken(String, string(value))
		}),
		lexer.NewRule("[a-zA-Z_][a-zA-Z0-9_]*", func(value []byte) *lexer.Token {
			return lexer.NewToken(Identifier, string(value))
		}),
	}
)
