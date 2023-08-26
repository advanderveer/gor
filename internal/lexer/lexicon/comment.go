package lexicon

import "github.com/advanderveer/gor/internal/lexer"

// https://go.dev/ref/spec#Comments
func lexCommentAndThen(next lexer.State) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		// @TODO implement

		return next
	}
}
