package scanner

// scanDigits scans a series of digits.
func (s *Scanner) scanDigits(base int, invalid *int) {
	if base <= 10 {
		max := rune('0' + base)
		for isDecimal(s.ch) {
			if s.ch >= max && *invalid < 0 {
				*invalid = s.offset // record invalid rune offset
			}

			s.next()
		}
	} else {
		for isHex(s.ch) || s.ch == '_' {
			s.next()
		}
	}
}

// scanNumber is called to scan a number given that the current
// character is a decimal, or fractional starting with a period.
func (s *Scanner) scanNumber() (Token, string) {
	start := s.offset
	tok := ILLEGAL
	base := 10        // number base
	prefix := rune(0) // one of 0 (decimal), '0' (0-octal), 'x', 'o', or 'b'
	invalid := -1     // index of invalid digit in literal, or < 0

	// integer part
	if s.ch != '.' {
		tok = INT

		if s.ch == '0' {
			s.next()

			switch lower(s.ch) {
			case 'x':
				s.next()
				base, prefix = 16, 'x'
			case 'o':
				s.next()
				base, prefix = 8, 'o'
			case 'b':
				s.next()
				base, prefix = 2, 'b'
			default:
				base, prefix = 8, '0'
			}
		}

		s.scanDigits(base, &invalid)
	}

	// fractional part
	if s.ch == '.' {
		tok = FLOAT
		if prefix == 'o' || prefix == 'b' {
			s.error(s.offset, "invalid radix point in "+numberKind(prefix))
		}

		s.next()
		s.scanDigits(base, &invalid)
	}

	return tok, string(s.src[start:s.offset])
}

// numberKind returns the kind of number for use in error messages.
func numberKind(prefix rune) string {
	switch prefix {
	case 'x':
		return "hexadecimal literal"
	case 'o', '0':
		return "octal literal"
	case 'b':
		return "binary literal"
	}

	return "decimal literal"
}
