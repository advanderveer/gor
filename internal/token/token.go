package token

// Token encodes a scanned token.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type Token -linecomment
type Token int

// IsLiteral returns true if the token represents a literal token.
func (tok Token) IsLiteral() bool {
	return literal_beg < tok && tok < literal_end
}

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
func (tok Token) Precedence() int {
	switch tok {
	case LOR:
		return 1
	case LAND:
		return 2
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, QUO, REM, SHL, SHR, AND, AND_NOT:
		return 5
	}

	return LowestPrec
}

const (
	// special tokens.
	ILLEGAL Token = iota // #illegal
	EOF                  // #eof
	IDENT                // #ident
	COMMENT              // #comment

	// keywords.
	BREAK       // break
	CASE        // case
	CHAN        // chan
	CONST       // const
	CONTINUE    // continue
	DEFAULT     // default
	DEFER       // defer
	ELSE        // else
	FALLTHROUGH // fallthrough
	FOR         // for
	FUNC        // func
	GO          // go
	GOTO        // goto
	IF          // if
	IMPORT      // import
	INTERFACE   // interface
	MAP         // map
	PACKAGE     // package
	RANGE       // range
	RETURN      // return
	SELECT      // select
	STRUCT      // struct
	SWITCH      // switch
	TYPE        // type
	VAR         // var

	// literals.
	literal_beg
	INT    // <int>
	FLOAT  // <float>
	STRING // <string>
	literal_end

	SEMICOLON // ;
	COLON     // :
	PERIOD    // .
	ELLIPSIS  // ...
	COMMA     // ,
	DEFINE    // :=
	ARROW     // <-

	INC // ++
	DEC // --

	ASSIGN // =
	EQL    // ==
	NEQ    // !=
	NOT    // !

	AND_NOT // &^
	AND     // &
	XOR     // ^
	OR      // |
	LAND    // &&
	LOR     // ||

	LPAREN // (
	RPAREN // )
	LBRACK // [
	RBRACK // ]
	LBRACE // {
	RBRACE // }

	LSS // <
	LEQ // <=
	SHL // <<
	GTR // >
	GEQ // >=
	SHR // >>

	QUO // /
	MUL // *
	ADD // +
	SUB // -
	REM // %

	QUO_ASSIGN     // /=
	MUL_ASSIGN     // *=
	ADD_ASSIGN     // +=
	SUB_ASSIGN     // -=
	AND_NOT_ASSIGN // &^=
	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	SHR_ASSIGN     // >>=
	SHL_ASSIGN     // <<=
	REM_ASSIGN     // %=
	XOR_ASSIGN     // ^=
)

// Lookup determines if the literal is a keyword token, or
// else an identifier.
func Lookup(lit string) Token {
	switch lit {
	case "break":
		return BREAK
	case "case":
		return CASE
	case "chan":
		return CHAN
	case "const":
		return CONST
	case "continue":
		return CONTINUE
	case "default":
		return DEFAULT
	case "defer":
		return DEFER
	case "else":
		return ELSE
	case "fallthrough":
		return FALLTHROUGH
	case "for":
		return FOR
	case "func":
		return FUNC
	case "go":
		return GO
	case "goto":
		return GOTO
	case "if":
		return IF
	case "import":
		return IMPORT
	case "interface":
		return INTERFACE
	case "map":
		return MAP
	case "package":
		return PACKAGE
	case "range":
		return RANGE
	case "return":
		return RETURN
	case "select":
		return SELECT
	case "struct":
		return STRUCT
	case "switch":
		return SWITCH
	case "type":
		return TYPE
	case "var":
		return VAR
	default:
		return IDENT
	}
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)
