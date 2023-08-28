package lexicon

import (
	"unicode"

	"github.com/advanderveer/gor/internal/lexer"
)

// https://go.dev/ref/spec#Tokens: White space, formed from spaces (U+0020), horizontal tabs (U+0009),
// carriage returns (U+000D), and newlines (U+000A).
func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || isNewline(r)
}

// isDoubleQuote returns true if the rune is a double quote.
func isDoubleQuote(r rune) bool {
	return r == '"'
}

// isLeftParen returns the rune for a left parenthesis.
func isLeftParen(r rune) bool {
	return r == '('
}

// isRightParen returns the rune for a right parenthesis.
func isRightParen(r rune) bool {
	return r == ')'
}

// isCommentCharacter return true if its the comment character.
func isCommentCharacter(r rune) bool {
	return r == '/'
}

// isEOF return true if the rune is the special EOF indicator.
func isEOF(r rune) bool {
	return r == lexer.EOF
}

// https://go.dev/ref/spec#newline
func isNewline(r rune) bool {
	return r == lexer.NewLine
}

// https://go.dev/ref/spec#unicode_char
func isUnicodeChar(r rune) bool {
	return !isNewline(r)
}

// https://go.dev/ref/spec#unicode_letter
func isUnicodeLetter(r rune) bool {
	return unicode.IsLetter(r)
}

// https://go.dev/ref/spec#unicode_digit
func isUnicodeDigit(r rune) bool {
	return unicode.Is(unicode.Nd, r)
}

// https://go.dev/ref/spec#decimal_digit
func isDecimalDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// https://go.dev/ref/spec#letter
func isLetter(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}
