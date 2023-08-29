// Package scanner implements a lexical scanner for Gor source text.
package scanner

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/advanderveer/gor/internal/token"
)

const (
	// end of file.
	eof = -1
)

// Scanner is responsible for scanning source code text and breaking
// it down into recognized tokens.
type Scanner struct {
	file *token.File
	src  []byte

	// scanning state
	ch         rune // current character
	offset     int  // character offset
	insertSemi bool // insert a semicolon before next newline
	rdOffset   int  // reading offset (position after current character)
	lineOffset int  // current line offset
}

// Init resets and initializes the scanner so it can be reused.
func (s *Scanner) Init(file *token.File, src []byte) {
	s.file = file
	s.src = src
	s.ch = ' '
	s.offset = 0
	s.rdOffset = 0
	s.lineOffset = 0
	s.insertSemi = false
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
func (s *Scanner) scanOpOrAssign(base, assign token.Token) token.Token {
	if s.ch == '=' {
		s.next()
		return assign
	}

	return base
}

// scanOpOrAssign is a 3-way check for the base operator, assign variant or some
// other variant triggered by 'ch2' argument.
func (s *Scanner) scanOpOrAssignOr(base, assign token.Token, ch2 rune, tok2 token.Token) token.Token {
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
func (s *Scanner) scanOpOrAssignOrOr(base, assign token.Token, ch2 rune, tok2, tok3 token.Token) token.Token {
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
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string) {
	s.skipWhitespace()

	switch chr := s.ch; {
	// scan identifier or keyword if the current char is a letter
	case isLetter(chr):
		lit = s.scanIdentifier()
		tok = token.Lookup(lit)

		switch tok { //nolint:exhaustive
		case token.IDENT, token.BREAK, token.CONTINUE, token.FALLTHROUGH, token.RETURN:
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
				return pos, token.SEMICOLON, "\n"
			}

			tok = token.EOF
		case '/':
			if s.ch == '/' {
				tok = token.COMMENT
				lit = s.scanComment()
			} else {
				tok = s.scanOpOrAssign(token.QUO, token.QUO_ASSIGN)
			}
		case '*':
			tok = s.scanOpOrAssign(token.MUL, token.MUL_ASSIGN)
		case '\n':
			// newline character, only reached if skipWhitespace was
			// called with s.insertSemi set to true.
			s.insertSemi = false // newline consumed
			return pos, token.SEMICOLON, "\n"
		case '"':
			s.insertSemi = true
			tok = token.STRING
			lit = s.scanString()
		case '`':
			s.insertSemi = true
			tok = token.STRING
			lit = s.scanRawString()
		case ':':
			tok = s.scanOpOrAssign(token.COLON, token.DEFINE)
		case '.':
			// fractions starting with a '.' are handled by outer switch
			tok = token.PERIOD
			if s.ch == '.' && s.peek() == '.' {
				s.next()
				s.next() // consume last '.'
				tok = token.ELLIPSIS
			}
		case ',':
			tok = token.COMMA
		case ';':
			tok = token.SEMICOLON
			lit = ";"
		case '(':
			tok = token.LPAREN
		case ')':
			s.insertSemi = true
			tok = token.RPAREN
		case '[':
			tok = token.LBRACK
		case ']':
			s.insertSemi = true
			tok = token.RBRACK
		case '{':
			tok = token.LBRACE
		case '}':
			s.insertSemi = true
			tok = token.RBRACE
		case '+':
			tok = s.scanOpOrAssignOr(token.ADD, token.ADD_ASSIGN, '+', token.INC)
			if tok == token.INC {
				s.insertSemi = true
			}
		case '-':
			tok = s.scanOpOrAssignOr(token.SUB, token.SUB_ASSIGN, '-', token.DEC)
			if tok == token.DEC {
				s.insertSemi = true
			}
		case '=':
			tok = s.scanOpOrAssign(token.ASSIGN, token.EQL)
		case '!':
			tok = s.scanOpOrAssign(token.NOT, token.NEQ)
		case '<':
			if s.ch == '-' {
				s.next()
				tok = token.ARROW
			} else {
				tok = s.scanOpOrAssignOrOr(token.LSS, token.LEQ, '<', token.SHL, token.SHL_ASSIGN)
			}
		case '>':
			tok = s.scanOpOrAssignOrOr(token.GTR, token.GEQ, '>', token.SHR, token.SHR_ASSIGN)
		case '&':
			if s.ch == '^' {
				s.next()
				tok = s.scanOpOrAssign(token.AND_NOT, token.AND_NOT_ASSIGN)
			} else {
				tok = s.scanOpOrAssignOr(token.AND, token.AND_ASSIGN, '&', token.LAND)
			}
		case '|':
			tok = s.scanOpOrAssignOr(token.OR, token.OR_ASSIGN, '|', token.LOR)
		case '%':
			tok = s.scanOpOrAssign(token.REM, token.REM_ASSIGN)
		case '^':
			tok = s.scanOpOrAssign(token.XOR, token.XOR_ASSIGN)
		default:
			tok = token.ILLEGAL
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
