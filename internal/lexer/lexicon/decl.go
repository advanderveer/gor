package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// https://go.dev/ref/spec#VarSpec
func lexVarSpecAndThen(next lexer.State, endfn func(rune) bool) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		return lexIdentListAndThen(func(c lexer.Control) lexer.State {
			lex.Next()
			lex.Emit(token.ASSIGN)

			return lexExpressionListAndThen(next, endfn)
		}, isAssign)
	}
}

// lex a block bar decl, inside parentheses.
func lexVarDeclBlockAndThen(next lexer.State) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Peek()

		switch {
		case isWhiteSpace(chr):
			lex.Skip(isWhiteSpace)

			return lexVarDeclBlockAndThen(next)
		case isRightParen(chr):
			lex.Next()
			lex.Emit(token.RPAREN)

			return next
		case isLetter(chr):
			return lexVarSpecAndThen(lexVarDeclBlockAndThen(next), isNewline)
		default:
			panic("not impelmented")
		}
	}
}

// https://go.dev/ref/spec#Variable_declarations
func lexVarDeclAndThen(next lexer.State) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Peek()

		switch {
		// skip any whitespace
		case isWhiteSpace(chr):
			lex.Skip(isWhiteSpace)

			return lexVarDeclAndThen(next)

		// list of variable declaration
		case isLeftParen(chr):
			lex.Next()
			lex.Emit(token.LPAREN)

			return lexVarDeclBlockAndThen(next)

		// single variable declaration
		case isLetter(chr):
			return lexVarSpecAndThen(next, func(r rune) bool {
				return isNewline(r) || isEOF(r)
			})
		default:
			return lex.Unexpected(chr,
				lexerr.ExpectedLeftParenthesis,
				lexerr.ExpectedLetter)
		}
	}
}

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
		lex.Emit(token.VAR)

		return lexVarDeclAndThen(lexDecls)
	// constant declaration
	case lex.Keyword("const"):
		panic("const declaration not implemented")
	default:
		return lex.Unexpected(chr,
			lexerr.ExpectedFuncKeyword,
			lexerr.ExpectedVarKeyword,
			lexerr.ExpectedConstKeyword)
	}
}
