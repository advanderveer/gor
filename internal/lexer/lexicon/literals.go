package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// https://go.dev/ref/spec#string_lit
func lexStringLiteralAndThen(next lexer.State) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Peek()

		switch {
		case isDoubleQuote(chr):
			lex.Next()
			lex.Ignore()

			lex.Accept(func(r rune) bool {
				return !isDoubleQuote(r)
			})

			lex.Emit(token.STRING)
			lex.Next()
			lex.Ignore()

			return next
		default:
			return lex.Unexpected(chr,
				lexerr.ExpectedDoubleQuote)
		}
	}
}
