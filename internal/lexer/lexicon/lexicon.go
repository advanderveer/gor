// Package lexicon implements the Gor language lexicon
package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// LexPackage implements lexing of a file's package declaration.
func LexPackage(lex lexer.Control) lexer.State {
	chr := lex.Peek()

	switch {
	// skip over any whitespace
	case isWhiteSpace(chr):
		lex.Skip(isWhiteSpace)

		return LexPackage
	// tokenize any comments before the package keyword
	case isCommentCharacter(chr):
		return lexCommentAndThen(LexPackage)
	// package keyword
	case isUnicodeLetter(chr):
		if !lex.Keyword("package") {
			return lex.Unexpected(chr, lexerr.ExpectedPackageKeyword)
		}

		lex.Emit(token.PACKAGE)
		lex.Skip(isWhiteSpace)

		return lexIdentifierAndThen(lexImports)
	default:
		return lex.Unexpected(chr,
			lexerr.ExpectedWhiteSpace,
			lexerr.ExpectedComment,
			lexerr.ExpectedUnicodeLetter)
	}
}
