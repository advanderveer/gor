package lexer

import (
	"errors"
	"fmt"

	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// Control exposes the combined scanning and lexing controls to state functions
// to implement any kind of Lexicon.
type Control interface {
	// Pos returns the current position of the scanner.
	Pos() Pos
	// Peek checks the next rune without accepting it for the next
	// token emit.
	Peek() rune
	// Value returns the currently pending token value.
	Value() string
	// Ignore prevents any currently pending runes from being emitted
	// for the next token.
	Ignore()
	// Backup rewinds the currently read rune and unschedules it from
	// being emitted as value for the next token.
	Backup()
	// Next progresses the lexer and accepts the Next character for the
	// token to be emitted.
	Next() rune
	// Keyword will look ahead for 'word' and skips over it
	// to not include it in the next emitted token's value.
	Keyword(string) bool
	// Accept runes for the pending token until 'f' returns false.
	Accept(func(rune) bool)
	// Skip runes for the pending token until 'f' returns false.
	Skip(func(rune) bool)
	// Emit the currently pending runes as the value of a [Token].
	Emit(tok token.Token)
	// Fail makes the lexer enter the error state while emitting
	// an error token for the developer.
	Fail(error) State
	// Unexpected fails the lexer while setting the fail state with a unexpected character error.
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

// Lex runs the lexer over the input source code. If any item contains an error
// they are joined into a single error and returned here.
func (l *Lexer) Lex() ([]Item, error) {
	for l.state != nil {
		l.state = l.state(l)
	}

	var err error
	for _, it := range l.output {
		err = errors.Join(err, it.Err)
	}

	return l.output, err
}
