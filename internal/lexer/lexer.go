package lexer

import (
	"fmt"
	"go/token"
)

// Control exposes the combined scanning and lexing controls to state functions
// to implement any kind of Lexicon.
type Control interface {
	Emit(token.Token)
	Pos() Pos
	Peek() rune
	Value() string
	Ignore()
	Backup()
	Next() rune
	Keyword(string) bool
	Accept(func(rune) bool)
	Errorf(string, ...any) State
}

// State defines a state that the lexer can be in.
type State func(Control) State

// Lexer implements Gor lexer.
type Lexer struct {
	*Scanner
	output []Item
	state  State
}

// New inits the lexer.
func New(input string, start State) *Lexer {
	return &Lexer{
		Scanner: NewScanner(input),
		state:   start,
	}
}

// Emit the currently pending runes as the value of a [Token].
func (l *Lexer) Emit(tok token.Token) {
	l.output = append(l.output, Item{
		Tok: tok,
		Pos: l.start,
		Val: l.Value(),
	})

	l.start = l.curr
}

// Errorf makes the lexer enter the error state while emitting
// an error token for the developer.
func (l *Lexer) Errorf(format string, args ...any) State {
	l.output = append(l.output, Item{
		Tok: token.ILLEGAL,
		Val: fmt.Sprintf(format, args...),
		Pos: l.curr,
	})

	return nil
}

// Lex runs the lexer over the input source code.
func (l *Lexer) Lex() []Item {
	for l.state != nil {
		l.state = l.state(l)
	}

	return l.output
}
