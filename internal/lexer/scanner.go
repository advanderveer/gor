// Package lexer implements a tokenizer for Gor source code.
package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	// NewLine character used for human-readable scan. It is exported to
	// allow other packages to use the same character.
	NewLine = 0x000A

	// EOF represents a special EOF rune to indicate file has ended.
	EOF rune = -1
)

// Pos describes an understandable position in a file.
type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d.%d", p.Offset, p.Line, p.Column)
}

// Scanner scans source code.
type Scanner struct {
	input string // input into the scanner
	curr  Pos    // current position in the input
	start Pos    // start position of the current token being scanned
	width int    // width of the last encountered rune
}

// NewScanner inits the scanner.
func NewScanner(input string) *Scanner {
	s := &Scanner{input: input}

	return s
}

// Value returns the currently pending token value.
func (s *Scanner) Value() string {
	return s.input[s.start.Offset:s.curr.Offset]
}

// Pos returns the lex'ers position info.
func (s *Scanner) Pos() Pos {
	return s.curr
}

// Skip runes for the pending token until 'f' returns false.
func (s *Scanner) Skip(fnc func(rune) bool) {
	defer s.Ignore()

	for {
		r := s.Next()
		if r == EOF {
			return
		}

		if !fnc(r) {
			s.Backup()

			return
		}
	}
}

// Accept runes for the pending token until 'f' returns false.
func (s *Scanner) Accept(fnc func(rune) bool) {
	for {
		r := s.Next()
		if r == EOF {
			return
		}

		if !fnc(r) {
			s.Backup()

			return
		}
	}
}

// Keyword will look ahead for 'word' and skips over it
// to not include it in the next emitted token's value.
func (s *Scanner) Keyword(word string) bool {
	if strings.HasPrefix(s.input[s.curr.Offset:], word) {
		for i := 0; i < utf8.RuneCountInString(word); i++ {
			s.Next()
		}

		return true
	}

	return false
}

// Ignore prevents any currently pending runes from being emitted
// for the next token.
func (s *Scanner) Ignore() {
	s.start = s.curr
}

// Peek checks the next rune without accepting it for the next
// token emit.
func (s *Scanner) Peek() rune {
	curr := s.Next()
	s.Backup()

	return curr
}

// Backup rewinds the currently read rune and unschedules it from
// being emitted as value for the next token.
func (s *Scanner) Backup() {
	// update line information
	prev, _ := utf8.DecodeRuneInString(s.input[s.curr.Offset-s.width:])
	if prev == NewLine {
		s.curr.Column = 0
		s.curr.Line--
	} else {
		s.curr.Column--
	}

	s.curr.Offset -= s.width
}

// Next progresses the lexer and accepts the Next character for the
// token to be emitted.
func (s *Scanner) Next() rune {
	if s.curr.Offset >= len(s.input) {
		s.width = 0

		return EOF
	}

	var curr rune
	curr, s.width = utf8.DecodeRuneInString(s.input[s.curr.Offset:])
	s.curr.Offset += s.width

	// update line information
	if curr == NewLine {
		s.curr.Column = 0
		s.curr.Line++
	} else {
		s.curr.Column++
	}

	return curr
}
