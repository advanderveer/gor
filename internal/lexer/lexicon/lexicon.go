// Package lexicon implements the Gor language lexicon
package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// LexImports implements lexing of import statements.
func LexImports(lexer.Control) lexer.State {
	return nil
}

// LexPackage implements lexing of a file's package declaration.
func LexPackage(lex lexer.Control) lexer.State {
	chr := lex.Peek()

	switch {
	case isWhiteSpace(chr):
		lex.Skip(isWhiteSpace)

		return LexPackage
	case isUnicodeLetter(chr):
		if !lex.Keyword("package") {
			return lex.Unexpected(chr, lexerr.ExpectedPackageKeyword)
		}

		lex.Emit(token.PACKAGE)
		lex.Skip(isWhiteSpace)

		// @TODO: add implicit semi-colon?

		return lexIdentifier(LexImports)
	default:
		return lex.Unexpected(chr,
			lexerr.ExpectedWhiteSpace,
			lexerr.ExpectedUnicodeLetter)
	}
}
