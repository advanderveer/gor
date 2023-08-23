// Package lexer implements a tokenizer for Gor source code.
package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// special eof rune to indicate file has ended.
const eof rune = -1

// Item encapsulates a scanned token.
type Item struct {
	Tok Token
	Val string
	Pos Pos
}

func (it Item) String() string {
	return fmt.Sprintf("%s:%s(%s)", it.Pos, it.Tok, it.Val)
}

// Pos describes an understandable position in a file.
type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d.%d", p.Offset, p.Line, p.Column)
}

// Lexer tokenizes Gor source code.
type Lexer struct {
	input  string // input into the lexer
	output []Item // scanned tokens
	state  StateFunc

	curr  Pos // current position in the input
	start Pos // start position of the current token being scanned
	width int // width of the last encountered rune
}

// New inits the lexer.
func New(input string, start StateFunc) *Lexer {
	l := &Lexer{state: start, input: input}

	return l
}

// Pos returns the lex'ers position info.
func (l *Lexer) Pos() Pos {
	return l.curr
}

// Accept runes for the pending token until 'f' returns false.
func (l *Lexer) Accept(f func(rune) bool) {
	for {
		if !f(l.Next()) {
			break
		}
	}
}

// Keyword will look ahead for 'word' and skips over it
// to not include it in the next emitted token's value.
func (l *Lexer) Keyword(word string) bool {
	if strings.HasPrefix(l.input[l.curr.Offset:], word) {
		for i := 0; i < utf8.RuneCountInString(word); i++ {
			l.Next()
		}
		l.Ignore() // keyword itself is never important

		return true
	}

	return false
}

// Emit the currently pending runes as the value of a [Token].
func (l *Lexer) Emit(tok Token) {
	l.output = append(l.output, Item{
		Tok: tok,
		Pos: l.start,
		Val: l.input[l.start.Offset:l.curr.Offset],
	})
	l.start = l.curr
}

// Errorf makes the lexer enter the error state while emitting
// an error token for the developer.
func (l *Lexer) Errorf(format string, args ...any) StateFunc {
	l.output = append(l.output, Item{
		Tok: ILLEGAL,
		Val: fmt.Sprintf(format, args...),
		Pos: l.curr,
	})

	return nil
}

// Ignore prevents any currently pending runes from being emitted
// for the next token.
func (l *Lexer) Ignore() {
	l.start = l.curr
}

// Peek checks the next rune without accepting it for the next
// token emit.
func (l *Lexer) Peek() rune {
	curr := l.Next()
	l.Backup()

	return curr
}

// Backup rewinds the currently read rune and unschedules it from
// being emitted as value for the next token.
func (l *Lexer) Backup() {
	// update line information
	prev, _ := utf8.DecodeRuneInString(l.input[l.curr.Offset-l.width:])
	if isNewline(prev) {
		l.curr.Column = 0
		l.curr.Line--
	} else {
		l.curr.Column--
	}

	l.curr.Offset -= l.width
}

// Next progresses the lexer and accepts the Next character for the
// token to be emitted.
func (l *Lexer) Next() rune {
	if l.curr.Offset >= len(l.input) {
		l.width = 0

		return eof
	}

	var curr rune
	curr, l.width = utf8.DecodeRuneInString(l.input[l.curr.Offset:])
	l.curr.Offset += l.width

	// update line information
	if isNewline(curr) {
		l.curr.Column = 0
		l.curr.Line++
	} else {
		l.curr.Column++
	}

	return curr
}

// Run the lexer over and return the output tokens.
func (l *Lexer) Run() []Item {
	for l.state != nil {
		l.state = l.state(l)
	}

	return l.output
}
