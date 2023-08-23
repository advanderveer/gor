package lexicon

import (
	"go/token"

	"github.com/advanderveer/gor/internal/lexer"
)

// https://go.dev/ref/spec#Identifiers
func acceptIdentifier(lex lexer.Control) error {
	chr := lex.Next()
	if !isLetter(chr) {
		return unexpectedInput(chr, lex.Pos(), "letter")
	}

	lex.Accept(func(r rune) bool {
		return isLetter(r) || isUnicodeDigit(r)
	})

	lex.Emit(token.IDENT)

	return nil
}
