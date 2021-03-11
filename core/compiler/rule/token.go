package rule

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

const (
	TokenLeftParenthesis  = iota // (
	TokenRightParenthesis        // )
	TokenLeftBracket             // [
	TokenRightBracket            // ]
	TokenLeftbrace               // {
	TokenRightbrace              // }
	TokenLeftArrow               // <-
	TokenSpace                   // \r \n \t [:space:]
	TokenColon                   // :
	TokenSemicolon               // ;
	TokenDot                     // .
	TokenComma                   // ,
	TokenNumber                  // [:number:]
	TokenString                  // [:string:]
	TokenIdentifier              // [:identifier:]
	TokenEnd                     // [:end:]
)

var (
	LexerRules = []*lexer.Rule{
		lexer.NewRule("\\(", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenLeftParenthesis, string(value))
		}),
		lexer.NewRule("\\)", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenRightParenthesis, string(value))
		}),
		lexer.NewRule("\\[", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenLeftBracket, string(value))
		}),
		lexer.NewRule("\\]", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenRightBracket, string(value))
		}),
		lexer.NewRule("\\{", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenLeftbrace, string(value))
		}),
		lexer.NewRule("\\}", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenRightbrace, string(value))
		}),
		lexer.NewRule("<\\-", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenLeftArrow, string(value))
		}),
		lexer.NewRule(":", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenColon, string(value))
		}),
		lexer.NewRule(";", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenSemicolon, string(value))
		}),
		lexer.NewRule("\\.", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenDot, string(value))
		}),
		lexer.NewRule("\\,", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenComma, string(value))
		}),
		lexer.NewRule("\\n", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenSpace, string(value))
		}),
		lexer.NewRule("\\r", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenSpace, string(value))
		}),
		lexer.NewRule("\\t", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenSpace, string(value))
		}),
		lexer.NewRule(" ", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenSpace, string(value))
		}),
		lexer.NewRule("([\\+|-]?\\d+)(\\.\\d+)?(E[\\+|-]?\\d+)?", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenNumber, string(value))
		}),
		lexer.NewRule("\"\\S*\"", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenString, string(value))
		}),
		lexer.NewRule("[a-zA-Z_][a-zA-Z0-9_]*", func(value []byte) *lexer.Token {
			return lexer.NewToken(TokenIdentifier, string(value))
		}),
	}
)
