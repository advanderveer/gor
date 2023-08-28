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
	// skip any whitespace
	case isWhiteSpace(chr):
		lex.Skip(isWhiteSpace)

		return lexImportSpec
	// unnamed import
	case isDoubleQuote(chr):
		return lexStringLiteralAndThen(lexImportSpec)
	// import package under a name
	case isLetter(chr):
		return lexIdentifierAndThen(func(c lexer.Control) lexer.State {
			lex.Skip(isWhiteSpace)

			return lexStringLiteralAndThen(lexImportSpec)
		})
	// close import statement
	case isRightParen(chr):
		lex.Next()
		lex.Emit(token.RPAREN)

		return lexDecls // transition to declaration lexing
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
	// file may end
	case isEOF(chr):
		return nil
	// skip over any whitespace
	case isWhiteSpace(chr):
		lex.Skip(isWhiteSpace)

		return lexImports
	// import statement start
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
	// tokenize comments before import statement
	case isCommentCharacter(chr):
		return lexCommentAndThen(lexImports)
	default:
		return lex.Unexpected(chr,
			lexerr.ExpectedWhiteSpace,
			lexerr.ExpectedComment,
			lexerr.ExpectedUnicodeLetter)
	}
}
