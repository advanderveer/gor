package scanner

import (
	"fmt"
	"unicode"
)

// scanRawString can be called to scan a uninterpolated (backtick) string.
func (s *Scanner) scanRawString() string {
	start := s.offset - 1 // opening "`" already scanned
	for {
		ch := s.ch
		if ch < 0 {
			s.error(start, "raw string literal not terminated")
			break
		}

		s.next()
		if ch == '`' {
			break
		}
	}

	return string(s.src[start:s.offset])
}

// scanString will scan a string literal and return it.
func (s *Scanner) scanString() string {
	offs := s.offset - 1 // opening '"' already consumed

	for {
		ch := s.ch
		if ch == '\n' || ch < 0 {
			s.error(offs, "string literal not terminated")
			break
		}

		s.next()
		if ch == '"' {
			break
		}

		if ch == '\\' {
			s.scanEscape('"')
		}
	}

	return string(s.src[offs:s.offset])
}

// scanEscape parses an escape sequence where rune is the accepted
// escaped quote. In case of a syntax error, it stops at the offending
// character (without consuming it) and returns false. Otherwise
// it returns true.
func (s *Scanner) scanEscape(quote rune) bool {
	offs := s.offset

	var num int
	var base, max uint32

	switch s.ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		s.next()
		return true
	case '0', '1', '2', '3', '4', '5', '6', '7':
		num, base, max = 3, 8, 255
	case 'x':
		s.next()
		num, base, max = 2, 16, 255
	case 'u':
		s.next()
		num, base, max = 4, 16, unicode.MaxRune
	case 'U':
		s.next()
		num, base, max = 8, 16, unicode.MaxRune
	default:
		msg := "unknown escape sequence"
		if s.ch < 0 {
			msg = "escape sequence not terminated"
		}

		s.error(offs, msg)
		return false
	}

	var x uint32
	for num > 0 {
		d := uint32(digitVal(s.ch))
		if d >= base {
			msg := fmt.Sprintf("illegal character %#U in escape sequence", s.ch)
			if s.ch < 0 {
				msg = "escape sequence not terminated"
			}

			s.error(s.offset, msg)
			return false
		}

		x = x*base + d
		s.next()
		num--
	}

	if x > max || 0xD800 <= x && x < 0xE000 {
		s.error(offs, "escape sequence is invalid Unicode code point")
		return false
	}

	return true
}

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= lower(ch) && lower(ch) <= 'f':
		return int(lower(ch) - 'a' + 10)
	}

	return 16 // larger than any legal digit val
}
