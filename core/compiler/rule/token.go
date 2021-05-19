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
	TypeEqual                   // =
	TypeColon                   // :
	TypeSemicolon               // ;
	TypeDot                     // .
	TypeComma                   // ,
	TypeNumber                  // [:number:]
	TypeString                  // [:string:]
	TypeEof                     // [:eof:]
	TypeIdentifier              // [:identifier:]
	TypeComment                 // [:comment:]
	TypePage                    // page
	TypeImport                  // import
	TypePublic                  // public
	TypePrivate                 // private
	TypeClass                   // class
	TypeRequire                 // require
	TypeProvide                 // provide
	TypeReturn                  // return
	TypeFunction                // function
	TypeGet                     // get
	TypeSet                     // set
	TypeTrue                    // true
	TypeFalse                   // false
	TypeNull                    // null
	TypeVar                     // var
	TypeIf                      // if
	TypeElse                    // else
	TypeFor                     // for
	TypeContinue                // continue
	TypeBreak                   // break
	TypeEnd                     // end
	TypeThis                    // this
	TypeSelf                    // self
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
	KeyEqual            = "equal"
	KeyColon            = "colon"
	KeySemicolon        = "semicolon"
	KeyDot              = "dot"
	KeyComma            = "comma"
	KeyNumber           = "number"
	KeyString           = "string"
	KeyEof              = "eof"
	KeyIdentifier       = "identifier"
	KeyComment          = "comment"
	KeyPage             = "page"
	KeyImport           = "import"
	KeyPublic           = "public"
	KeyPrivate          = "private"
	KeyClass            = "class"
	KeyRequire          = "require"
	KeyProvide          = "provide"
	KeyReturn           = "return"
	KeyFunction         = "function"
	KeyGet              = "get"
	KeySet              = "set"
	KeyTrue             = "true"
	KeyFalse            = "false"
	KeyNull             = "null"
	KeyVar              = "var"
	KeyIf               = "if"
	KeyElse             = "else"
	KeyFor              = "for"
	KeyContinue         = "continue"
	KeyBreak            = "break"
	KeyEnd              = "end"
	KeyThis             = "this"
	KeySelf             = "self"
)

var (
	LexerEof = lexer.NewToken(TypeEof, KeyEof, "EOF")

	LexerTrim = []int{TypeSpace, TypeComment}

	LexerRules = []*lexer.Rule{
		lexer.NewRule("\\(", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftParenthesis, KeyLeftParenthesis, "(")
		}),
		lexer.NewRule("\\)", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightParenthesis, KeyRightParenthesis, ")")
		}),
		lexer.NewRule("\\[", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftBracket, KeyLeftBracket, "[")
		}),
		lexer.NewRule("\\]", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightBracket, KeyRightBracket, "]")
		}),
		lexer.NewRule("\\{", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftBrace, KeyLeftBrace, "{")
		}),
		lexer.NewRule("\\}", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightBrace, KeyRightBrace, "}")
		}),
		lexer.NewRule("<\\-", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeLeftArrow, KeyLeftArrow, "<-")
		}),
		lexer.NewRule("\\->", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeRightArrow, KeyRightArrow, "->")
		}),
		lexer.NewRule("=", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeEqual, KeyEqual, "=")
		}),
		lexer.NewRule(":", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeColon, KeyColon, ":")
		}),
		lexer.NewRule(";", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSemicolon, KeySemicolon, ";")
		}),
		lexer.NewRule("\\.", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeDot, KeyDot, ".")
		}),
		lexer.NewRule("\\,", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeComma, KeyComma, ",")
		}),
		lexer.NewRule("\\n", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, "\n")
		}),
		lexer.NewRule("\\r", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, "\r")
		}),
		lexer.NewRule("\\t", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, "\t")
		}),
		lexer.NewRule(" ", func(value []byte) *lexer.Token {
			return lexer.NewToken(TypeSpace, KeySpace, " ")
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
			case KeyPublic:
				return lexer.NewToken(TypePublic, KeyPublic, KeyPublic)
			case KeyPrivate:
				return lexer.NewToken(TypePrivate, KeyPrivate, KeyPrivate)
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
			case KeyNull:
				return lexer.NewToken(TypeNull, KeyNull, KeyNull)
			case KeyVar:
				return lexer.NewToken(TypeVar, KeyVar, KeyVar)
			case KeyIf:
				return lexer.NewToken(TypeIf, KeyIf, KeyIf)
			case KeyElse:
				return lexer.NewToken(TypeElse, KeyElse, KeyElse)
			case KeyFor:
				return lexer.NewToken(TypeFor, KeyFor, KeyFor)
			case KeyContinue:
				return lexer.NewToken(TypeContinue, KeyContinue, KeyContinue)
			case KeyBreak:
				return lexer.NewToken(TypeBreak, KeyBreak, KeyBreak)
			case KeyEnd:
				return lexer.NewToken(TypeEnd, KeyEnd, KeyEnd)
			case KeyThis:
				return lexer.NewToken(TypeThis, KeyThis, KeyThis)
			case KeySelf:
				return lexer.NewToken(TypeSelf, KeySelf, KeySelf)
			default:
				return lexer.NewToken(TypeIdentifier, KeyIdentifier, valueIdentifier)
			}
		}),
	}
)
