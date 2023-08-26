package lexer

import (
	"fmt"

	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
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
	Skip(func(rune) bool)
	Fail(error) State
	Unexpected(r rune, exp lexerr.ExpectCode, more ...lexerr.ExpectCode) State
}

// State defines a state that the lexer can be in.
type State func(Control) State

// Item encapsulates a scanned token.
type Item struct {
	Tok token.Token
	Val string
	Err error
	Pos Pos
}

func (it Item) String() string {
	return fmt.Sprintf("%s:%s(%s)", it.Pos, it.Tok, it.Val)
}

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

// Unexpected fails the lexer while setting the fail state with a unexpected character error.
func (l *Lexer) Unexpected(r rune, exp lexerr.ExpectCode, more ...lexerr.ExpectCode) State {
	return l.Fail(lexerr.Unexpected(r, exp, more...))
}

// Fail makes the lexer enter the error state while emitting
// an error token for the developer.
func (l *Lexer) Fail(err error) State {
	l.output = append(l.output, Item{
		Tok: token.ILLEGAL,
		Val: err.Error(),
		Err: err,
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
