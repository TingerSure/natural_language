package rule

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

const (
	TypeLeftParenthesis  = iota // (
	TypeRightParenthesis        // )
	TypeLeftBracket             // [
	TypeRightBracket            // ]
	TypeLeftBrace               // {
	TypeRightBrace              // }
	TypeLeftArrow               // <-
	TypeRightArrow              // ->
	TypeSpace                   // \r \n \t [:space:]
	TypeColon                   // :
	TypeSemicolon               // ;
	TypeDot                     // .
	TypeComma                   // ,
	TypeNumber                  // [:number:]
	TypeString                  // [:string:]
	TypeEnd                     // [:end:]
	TypeIdentifier              // [:identifier:]
	TypeComment                 // [:comment:]
	TypePage                    // page
	TypeImport                  // import
	TypeExport                  // export
	TypeVar                     // var
	TypeClass                   // class
	TypeRequire                 // require
	TypeProvide                 // provide
	TypeReturn                  // return
	TypeFunction                // function
	TypeGet                     // get
	TypeSet                     // set
	TypeTrue                    // true
	TypeFalse                   // false
)

const (
	KeyLeftParenthesis  = "left_parenthesis"
	KeyRightParenthesis = "right_parenthesis"
	KeyLeftBracket      = "left_bracket"
	KeyRightBracket     = "right_bracket"
	KeyLeftBrace        = "left_brace"
	KeyRightBrace       = "right_brace"
	KeyLeftArrow        = "left_arrow"
	KeyRightArrow       = "right_arrow"
	KeySpace            = "space"
	KeyColon            = "colon"
	KeySemicolon        = "semicolon"
	KeyDot              = "dot"
	KeyComma            = "comma"
	KeyNumber           = "number"
	KeyString           = "string"
	KeyEnd              = "end"
	KeyIdentifier       = "identifier"
	KeyComment          = "comment"
	KeyPage             = "page"
	KeyImport           = "import"
	KeyExport           = "export"
	KeyVar              = "var"
	KeyClass            = "class"
	KeyRequire          = "require"
	KeyProvide          = "provide"
	KeyReturn           = "return"
	KeyFunction         = "function"
	KeyGet              = "get"
	KeySet              = "set"
	KeyTrue             = "true"
	KeyFalse            = "false"
)

var (
	LexerEnd = lexer.NewToken(TypeEnd, KeyEnd, "EOF")

	LexerTrim = []int{TypeSpace, TypeComment}

	LexerRules = []*lexer.Rule{
		lexer.NewRule("\\(", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftParenthesis, KeyLeftParenthesis, string(value))
		}),
		lexer.NewRule("\\)", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightParenthesis, KeyRightParenthesis, string(value))
		}),
		lexer.NewRule("\\[", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftBracket, KeyLeftBracket, string(value))
		}),
		lexer.NewRule("\\]", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightBracket, KeyRightBracket, string(value))
		}),
		lexer.NewRule("\\{", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftBrace, KeyLeftBrace, string(value))
		}),
		lexer.NewRule("\\}", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightBrace, KeyRightBrace, string(value))
		}),
		lexer.NewRule("<\\-", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftArrow, KeyLeftArrow, string(value))
		}),
		lexer.NewRule("\\->", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightArrow, KeyRightArrow, string(value))
		}),
		lexer.NewRule(":", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeColon, KeyColon, string(value))
		}),
		lexer.NewRule(";", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSemicolon, KeySemicolon, string(value))
		}),
		lexer.NewRule("\\.", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeDot, KeyDot, string(value))
		}),
		lexer.NewRule("\\,", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeComma, KeyComma, string(value))
		}),
		lexer.NewRule("\\n", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, string(value))
		}),
		lexer.NewRule("\\r", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, string(value))
		}),
		lexer.NewRule("\\t", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, string(value))
		}),
		lexer.NewRule(" ", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, string(value))
		}),
		lexer.NewRule("([\\+|-]?\\d+)(\\.\\d+)?(E[\\+|-]?\\d+)?", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeNumber, KeyNumber, string(value))
		}),
		lexer.NewRule("\"[^\"\\r\\n]*\"", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeString, KeyString, string(value))
		}),
		lexer.NewRule("//.*", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeComment, KeyComment, string(value))
		}),
		lexer.NewRule("/\\*[\\s\\S]*?\\*/", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeComment, KeyComment, string(value))
		}),
		lexer.NewRule("[a-zA-Z_][a-zA-Z0-9_]*", func(value []byte) *lexer.Token {
			valueIdentifier := string(value)
			switch valueIdentifier {
			case KeyPage:
				return lexer.NewToken(TypePage, KeyPage, KeyPage)
			case KeyImport:
				return lexer.NewToken(TypeImport, KeyImport, KeyImport)
			case KeyExport:
				return lexer.NewToken(TypeExport, KeyExport, KeyExport)
			case KeyVar:
				return lexer.NewToken(TypeVar, KeyVar, KeyVar)
			case KeyClass:
				return lexer.NewToken(TypeClass, KeyClass, KeyClass)
			case KeyRequire:
				return lexer.NewToken(TypeRequire, KeyRequire, KeyRequire)
			case KeyProvide:
				return lexer.NewToken(TypeProvide, KeyProvide, KeyProvide)
			case KeyReturn:
				return lexer.NewToken(TypeReturn, KeyReturn, KeyReturn)
			case KeyFunction:
				return lexer.NewToken(TypeFunction, KeyFunction, KeyFunction)
			case KeyGet:
				return lexer.NewToken(TypeGet, KeyGet, KeyGet)
			case KeySet:
				return lexer.NewToken(TypeSet, KeySet, KeySet)
			case KeyTrue:
				return lexer.NewToken(TypeTrue, KeyTrue, KeyTrue)
			case KeyFalse:
				return lexer.NewToken(TypeFalse, KeyFalse, KeyFalse)
			default:
				return lexer.NewToken(TypeIdentifier, KeyIdentifier, valueIdentifier)
			}
		}),
	}
)
