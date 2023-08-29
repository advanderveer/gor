package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// https://go.dev/ref/spec#Identifiers
func lexIdentifierAndThen(next lexer.State) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Next()
		if !isLetter(chr) {
			return lex.Unexpected(chr, lexerr.ExpectedLetter)
		}

		lex.Accept(func(r rune) bool {
			return isLetter(r) || isUnicodeDigit(r)
		})

		lex.Emit(token.IDENT)

		return next
	}
}

// https://go.dev/ref/spec#IdentifierList
func lexIdentListAndThen(next lexer.State, endfn func(rune) bool) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Peek()

		switch {
		case isWhiteSpace(chr):
			lex.Skip(isWhiteSpace)

			return lexIdentListAndThen(next, endfn)
		case isLetter(chr):
			return lexIdentifierAndThen(lexIdentListAndThen(next, endfn))
		case isComma(chr):
			lex.Next()
			lex.Emit(token.COMMA)

			return lexIdentListAndThen(next, endfn)
		case endfn(chr):
			return next
		default:
			return lex.Unexpected(chr,
				lexerr.ExpectedWhiteSpace,
				lexerr.ExpectedLetter,
				lexerr.ExpectedComma,
			)
		}
	}
}
