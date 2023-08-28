// Package lexerr describes errors during lexical scanning.
package lexerr

import (
	"fmt"
	"strings"
)

// ExpectCode codifies what a lexer expects so we can consistently
// provide errors to the developer.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type ExpectCode -linecomment
type ExpectCode int

const (
	// ExpectedPackageKeyword codes the expectation of a package keyword.
	ExpectedPackageKeyword ExpectCode = iota // 'package' keyword
	// ExpectedImportKeyword codes the expectation of a import keyword.
	ExpectedImportKeyword // 'import' keyword
	// ExpectedFuncKeyword codes the expectation for a const keyword.
	ExpectedFuncKeyword // 'func' keyword
	// ExpectedVarKeyword codes the expectation for a const keyword.
	ExpectedVarKeyword // 'var' keyword
	// ExpectedConstKeyword codes the expectation for a const keyword.
	ExpectedConstKeyword // 'const' keyword
	// ExpectedLetter codes the expectation of a letter.
	ExpectedLetter // letter
	// ExpectedWhiteSpace codes the expectation of a white space.
	ExpectedWhiteSpace // white space
	// ExpectedComment codes the expectation of a comment.
	ExpectedComment // comment
	// FirstCommentCharacter codes the expectation for the comment first forward slash.
	FirstCommentCharacter // first comment character
	// SecondCommentCharacter codes the expectation for the comment Second forward slash.
	SecondCommentCharacter // second comment character
	// ExpectedLeftParenthesis codes the expectation for the left parenthesis.
	ExpectedLeftParenthesis // left parenthesis
	// ExpectedRightParenthesis codes the expectation for the right parenthesis.
	ExpectedRightParenthesis // right parenthesis
	// ExpectedStringLiteral codes the expectation for a string literal.
	ExpectedStringLiteral // string literal
	// ExpectedDoubleQuote codes the expectation for a double quote.
	ExpectedDoubleQuote // double quote
)

// UnexpectedError encodes the error for an unexpected character during tokenization.
type UnexpectedError struct {
	got rune
	exp []ExpectCode
}

func (e UnexpectedError) Error() string {
	switch {
	case len(e.exp) < 1:
		return fmt.Sprintf("unexpected input '%s'", string(e.got))
	case len(e.exp) < 2: //nolint: gomnd
		return fmt.Sprintf("unexpected input '%s', expected: %s", string(e.got), e.exp[0])
	default:
		var msg strings.Builder

		fmt.Fprintf(&msg, `unexpected input '%s', expected:`, string(e.got))

		for i := 0; i < len(e.exp)-1; i++ {
			if i != 0 {
				msg.WriteRune(',')
			}

			fmt.Fprintf(&msg, ` %s`, e.exp[i])
		}

		fmt.Fprintf(&msg, ` or %s`, e.exp[len(e.exp)-1])

		return msg.String()
	}
}

// Unexpected creates an unexpected rune error for the lexer.
func Unexpected(got rune, exp ExpectCode, more ...ExpectCode) error {
	return fmt.Errorf("%w", &UnexpectedError{got: got, exp: append([]ExpectCode{exp}, more...)})
}
