package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"
)

// https://go.dev/ref/spec#Comments
func lexCommentAndThen(next lexer.State) func(lexer.Control) lexer.State {
	return func(lex lexer.Control) lexer.State {
		chr := lex.Next()
		if !isCommentCharacter(chr) {
			return lex.Unexpected(chr, lexerr.FirstCommentCharacter)
		}

		chr = lex.Next()
		if !isCommentCharacter(chr) {
			return lex.Unexpected(chr, lexerr.SecondCommentCharacter)
		}

		lex.Ignore()
		lex.Accept(func(r rune) bool { return !isNewline(r) })
		lex.Emit(token.COMMENT)

		return next
	}
}
