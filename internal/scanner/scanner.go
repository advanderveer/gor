// Package scanner implements a lexical scanner for Gor source text.
package scanner

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// File with source code.
type File struct{}

// AddLine adds the line offset for a new line.
func (File) AddLine(int) {
}

// Pos encodes the position in a source file.
type Pos int

// Token encodes a scanned token.
type Token int

const (
	// special tokens.
	ILLEGAL Token = iota
	EOF
	IDENT
	COMMENT

	// keywords.
	BREAK       // break
	CONTINUE    // continue
	FALLTHROUGH // fallthrough
	RETURN      // return

	// literals.
	INT    // <int>
	FLOAT  // <float>
	STRING // <string>

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

// AsToken determines if the literal is a keyword token, or
// else an identifier.
func AsToken(lit string) Token {
	switch lit {
	case "break":
		return BREAK
	case "continue":
		return CONTINUE
	case "fallthrough":
		return FALLTHROUGH
	case "return":
		return RETURN
	default:
		return IDENT
	}
}

const (
	// end of file.
	eof = -1
)

// Scanner is responsible for scanning source code text and breaking
// it down into recognized tokens.
type Scanner struct {
	file *File
	src  []byte

	// scanning state
	ch         rune // current character
	offset     int  // character offset
	insertSemi bool // insert a semicolon before next newline
	rdOffset   int  // reading offset (position after current character)
	lineOffset int  // current line offset
}

// New inits the source file tokenizer.
func New(file *File, src []byte) *Scanner {
	return &Scanner{
		file:   file,
		src:    src,
		ch:     ' ',
		offset: 0,
	}
}

// error during scanning.
func (s *Scanner) error(offs int, msg string) {
	panic(fmt.Sprintf("error@%d: %s", offs, msg))
}

// peek returns the byte following the most recently read character without
// advancing the scanner. If the scanner is at EOF, peek returns 0.
func (s *Scanner) peek() byte {
	if s.rdOffset < len(s.src) {
		return s.src[s.rdOffset]
	}

	return 0
}

// Read the next Unicode char into s.ch.
func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		if s.ch == '\n' {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}

		r, w := rune(s.src[s.rdOffset]), 1

		switch {
		case r == 0:
			s.error(s.offset, "illegal character NUL")
		case r >= utf8.RuneSelf:
			// not ASCII
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				s.error(s.offset, "illegal UTF-8 encoding")
			}
		}

		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		if s.ch == '\n' {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}
		s.ch = eof
	}
}

// skipWhitespace moves the scanner over any whitespace.
func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' && !s.insertSemi || s.ch == '\r' {
		s.next()
	}
}

// scanIdentifier can be called to scan the identifier if it already
// known that the current character is a letter.
func (s *Scanner) scanIdentifier() string {
	start := s.offset

	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}

	return string(s.src[start:s.offset])
}

// scanOpOrAssign returns oof the two tokens based on wether the current
// scanned char is '='.
func (s *Scanner) scanOpOrAssign(base, assign Token) Token {
	if s.ch == '=' {
		s.next()
		return assign
	}

	return base
}

// scanOpOrAssign is a 3-way check for the base operator, assign variant or some
// other variant triggered by 'ch2' argument.
func (s *Scanner) scanOpOrAssignOr(base, assign Token, ch2 rune, tok2 Token) Token {
	if s.ch == '=' {
		s.next()
		return assign
	}

	if s.ch == ch2 {
		s.next()
		return tok2
	}

	return base
}

// scanOpOrAssignOrOr is a 4-way variant of the base token. Assign or more variants that
// trigger depending on the 'ch2' argument.
func (s *Scanner) scanOpOrAssignOrOr(base, assign Token, ch2 rune, tok2, tok3 Token) Token {
	if s.ch == '=' {
		s.next()
		return assign
	}

	if s.ch == ch2 {
		s.next()
		if s.ch == '=' {
			s.next()
			return tok3
		}
		return tok2
	}

	return base
}

// Scan the next token while returning the position and literal value.
func (s *Scanner) Scan() (pos Pos, tok Token, lit string) {
	s.skipWhitespace()

	switch chr := s.ch; {
	// scan identifier or keyword if the current char is a letter
	case isLetter(chr):
		lit = s.scanIdentifier()
		tok = AsToken(lit)

		switch tok { //nolint:exhaustive
		case IDENT, BREAK, CONTINUE, FALLTHROUGH, RETURN:
			s.insertSemi = true
		}

	// scan number if decimal or fractional starting with a period.
	case isDecimal(chr) || chr == '.' && isDecimal(rune(s.peek())):
		s.insertSemi = true
		tok, lit = s.scanNumber()

	// all the other tokens
	default:
		s.next() // already move the scanner forward
		switch chr {
		case eof:
			if s.insertSemi {
				s.insertSemi = false // EOF consumed
				return pos, SEMICOLON, "\n"
			}

			tok = EOF
		case '/':
			if s.ch == '/' {
				tok = COMMENT
				lit = s.scanComment()
			} else {
				tok = s.scanOpOrAssign(QUO, QUO_ASSIGN)
			}
		case '*':
			tok = s.scanOpOrAssign(MUL, MUL_ASSIGN)
		case '\n':
			// newline character, only reached if skipWhitespace was
			// called with s.insertSemi set to true.
			s.insertSemi = false // newline consumed
			return pos, SEMICOLON, "\n"
		case '"':
			s.insertSemi = true
			tok = STRING
			lit = s.scanString()
		case '`':
			s.insertSemi = true
			tok = STRING
			lit = s.scanRawString()
		case ':':
			tok = s.scanOpOrAssign(COLON, DEFINE)
		case '.':
			// fractions starting with a '.' are handled by outer switch
			tok = PERIOD
			if s.ch == '.' && s.peek() == '.' {
				s.next()
				s.next() // consume last '.'
				tok = ELLIPSIS
			}
		case ',':
			tok = COMMA
		case ';':
			tok = SEMICOLON
			lit = ";"
		case '(':
			tok = LPAREN
		case ')':
			s.insertSemi = true
			tok = RPAREN
		case '[':
			tok = LBRACK
		case ']':
			s.insertSemi = true
			tok = RBRACK
		case '{':
			tok = LBRACE
		case '}':
			s.insertSemi = true
			tok = RBRACE
		case '+':
			tok = s.scanOpOrAssignOr(ADD, ADD_ASSIGN, '+', INC)
			if tok == INC {
				s.insertSemi = true
			}
		case '-':
			tok = s.scanOpOrAssignOr(SUB, SUB_ASSIGN, '-', DEC)
			if tok == DEC {
				s.insertSemi = true
			}
		case '=':
			tok = s.scanOpOrAssign(ASSIGN, EQL)
		case '!':
			tok = s.scanOpOrAssign(NOT, NEQ)
		case '<':
			if s.ch == '-' {
				s.next()
				tok = ARROW
			} else {
				tok = s.scanOpOrAssignOrOr(LSS, LEQ, '<', SHL, SHL_ASSIGN)
			}
		case '>':
			tok = s.scanOpOrAssignOrOr(GTR, GEQ, '>', SHR, SHR_ASSIGN)
		case '&':
			if s.ch == '^' {
				s.next()
				tok = s.scanOpOrAssign(AND_NOT, AND_NOT_ASSIGN)
			} else {
				tok = s.scanOpOrAssignOr(AND, AND_ASSIGN, '&', LAND)
			}
		case '|':
			tok = s.scanOpOrAssignOr(OR, OR_ASSIGN, '|', LOR)
		case '%':
			tok = s.scanOpOrAssign(REM, REM_ASSIGN)
		case '^':
			tok = s.scanOpOrAssign(XOR, XOR_ASSIGN)
		default:
			tok = ILLEGAL
			lit = string(chr)
		}
	}

	return pos, tok, lit
}

// isLetter returns true if the rune is considered a letter.
func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

// isDigit returns true if the rune is considered a digit.
func isDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

// lower converts the rune to lower case.
func lower(ch rune) rune {
	return ('a' - 'A') | ch
}

// isDecimal return true if the rune represents a decimal.
func isDecimal(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isHex returns whether the character is part of the hexadecimal encoding.
func isHex(ch rune) bool {
	return '0' <= ch && ch <= '9' || 'a' <= lower(ch) && lower(ch) <= 'f'
}
