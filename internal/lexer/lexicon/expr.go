package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// https://go.dev/ref/spec#Expression
func lexExpressionAndThen(next lexer.State) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Peek()

		switch {
		case isDoubleQuote(chr):
			return lexStringLiteralAndThen(next)
		default:
			if chr == ')' {
				panic("huh")
			}

			return lex.Unexpected(chr,
				lexerr.ExpectedDoubleQuote)
		}
	}
}

// https://go.dev/ref/spec#ExpressionList
func lexExpressionListAndThen(next lexer.State, endfn func(rune) bool) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Peek()

		switch {
		case isWhiteSpace(chr):
			lex.Skip(isWhiteSpace)

			return lexExpressionListAndThen(next, endfn)
		case isComma(chr):
			lex.Next()
			lex.Emit(token.COMMA)

			return lexExpressionListAndThen(next, endfn)
		case endfn(chr):
			return next
		default:
			return lexExpressionAndThen(
				lexExpressionListAndThen(next, endfn))
		}
	}
}
