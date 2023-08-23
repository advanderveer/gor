package lexicon

import (
	"fmt"

	"github.com/advanderveer/gor/internal/lexer"
)

// UnexpectedInputError is returned when the lexer encountered an unexpected
// character or rune.
type UnexpectedInputError struct {
	expected string
	pos      lexer.Pos
	got      rune
}

func (e UnexpectedInputError) Error() string {
	return fmt.Sprintf("%s: invalid input, got: '%s' expected: '%s'", e.pos, string(e.got), e.expected)
}

// unexpectedInput creates a wrapped error for unexpected input.
func unexpectedInput(got rune, pos lexer.Pos, expected string) error {
	return fmt.Errorf("%w", UnexpectedInputError{got: got, pos: pos, expected: expected})
}
