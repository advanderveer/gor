package lexer

import "unicode"

// https://go.dev/ref/spec#newline
func isNewline(r rune) bool {
	return r == 0x000A //nolint:gomnd
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
