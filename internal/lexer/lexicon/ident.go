package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// https://go.dev/ref/spec#Identifiers
func lexIdentifier(next lexer.State) func(lexer.Control) lexer.State {
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
