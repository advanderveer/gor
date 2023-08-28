package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// lexImportSpec lexes a single import.
func lexImportSpec(lex lexer.Control) lexer.State {
	chr := lex.Peek()

	switch {
	case isWhiteSpace(chr):
		lex.Skip(isWhiteSpace)

		return lexImportSpec
	case isDoubleQuote(chr):
		return lexStringLiteralAndThen(lexImportSpec)
	case isLetter(chr):
		return lexIdentifierAndThen(func(c lexer.Control) lexer.State {
			lex.Skip(isWhiteSpace)

			return lexStringLiteralAndThen(lexImportSpec)
		})
	case isRightParen(chr):
		lex.Next()
		lex.Emit(token.RPAREN)

		return lexDecls
	default:
		return lex.Unexpected(chr,
			lexerr.ExpectedWhiteSpace,
			lexerr.ExpectedLetter,
			lexerr.ExpectedStringLiteral,
		)
	}
}

// lexImports implements lexing of import statements.
func lexImports(lex lexer.Control) lexer.State {
	chr := lex.Peek()

	switch {
	case isEOF(chr):
		return nil // done
	case isWhiteSpace(chr):
		lex.Skip(isWhiteSpace)

		return lexImports
	case isUnicodeLetter(chr):
		if !lex.Keyword("import") {
			return lex.Unexpected(chr, lexerr.ExpectedImportKeyword)
		}

		lex.Emit(token.IMPORT)
		lex.Skip(isWhiteSpace)

		chr := lex.Next()
		if !isLeftParen(chr) {
			return lex.Unexpected(chr, lexerr.ExpectedLeftParenthesis)
		}

		lex.Emit(token.LPAREN)

		return lexImportSpec
	case isCommentCharacter(chr):
		return lexCommentAndThen(lexImports)
	default:
		return lex.Unexpected(chr,
			lexerr.ExpectedWhiteSpace,
			lexerr.ExpectedComment,
			lexerr.ExpectedUnicodeLetter)
	}
}
