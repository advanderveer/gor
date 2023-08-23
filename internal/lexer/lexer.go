package lexer

import "fmt"

// Control exposes the combined scanning and lexing controls to state functions
// to implement any kind of Lexicon.
type Control interface {
	Emit(Token)
	Pos() Pos
	Peek() rune
	Ignore()
	Backup()
	Next() rune
	Keyword(string) bool
	Accept(func(rune) bool)
	Errorf(string, ...any) StateFunc
}

// StateFunc defines what token is valid after another.
type StateFunc func(Control) StateFunc

// Lexer implements Gor lexer.
type Lexer struct {
	*Scanner
	output []Item
	state  StateFunc
}

// New inits the lexer.
func New(input string, start StateFunc) *Lexer {
	return &Lexer{
		Scanner: NewScanner(input),
		state:   start,
	}
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

// Lex runs the lexer over the input source code.
func (l *Lexer) Lex() []Item {
	for l.state != nil {
		l.state = l.state(l)
	}

	return l.output
}
