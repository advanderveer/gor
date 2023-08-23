package lexer

// https://go.dev/ref/spec#Identifiers
// func acceptIdentifier(lex LexControl) error {
// 	chr := lex.Next()
// 	if !isLetter(chr) {
// 		return unexpectedInput(chr, lex.Pos(), "letter")
// 	}

// 	lex.Accept(func(r rune) bool {
// 		return isLetter(r) || isUnicodeDigit(r)
// 	})

// 	lex.Emit(IDENT)

// 	return nil
// }
