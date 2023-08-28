package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
)

// lexDecls implements lexing of package declarations.
func lexDecls(lex lexer.Control) lexer.State {
	chr := lex.Peek()

	switch {
	// eof is fine
	case isEOF(chr):
		return nil
	// skip over whitespace
	case isWhiteSpace(chr):
		lex.Skip(isWhiteSpace)

		return lexDecls
	// function or method declaration
	case lex.Keyword("func"):
		return nil
	// global package variable declaration
	case lex.Keyword("var"):
		return nil
	// constant declaration
	case lex.Keyword("const"):
		return nil
	default:
		return lex.Unexpected(chr,
			lexerr.ExpectedFuncKeyword,
			lexerr.ExpectedVarKeyword,
			lexerr.ExpectedConstKeyword)
	}
}
