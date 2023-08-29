package scanner

// scanComment returns the text of the comment.
func (s *Scanner) scanComment() string {
	start := s.offset - 1 // position of initial '/', already consumed

	// "//" style comment
	if s.ch == '/' {
		s.next()
		for s.ch != '\n' && s.ch >= 0 {
			s.next()
		}

		goto exit
	}

	s.error(start, "comment not terminated")

exit:
	return string(s.src[start:s.offset])
}
