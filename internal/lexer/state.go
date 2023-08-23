package lexer

// // lexWhiteSpace scans white space.
// func lexWhiteSpace(lex LexControl) {
// 	for {
// 		chr := lex.next()
// 		if !unicode.IsSpace(chr) {
// 			lex.backup()
// 			lex.emit(WS)

// 			return
// 		}
// 	}
// }

// // lexComment will scan comments.
// func lexComment(LexControl) {}

// // lexImports will lex import statements.
// func lexImports(LexControl) StateFunc {
// 	return nil
// }

// // LexPackage scans for the initial file header.
// func LexPackage(lex LexControl) StateFunc {
// 	// chr := lex.peek()

// 	// switch {
// 	// case isWhiteSpace(chr):
// 	// 	lexWhiteSpace(lex)

// 	// 	return LexPackage
// 	// case isCommentSymbol(chr):
// 	// 	lexComment(lex)

// 	// 	return LexPackage
// 	// case isLatinLetter(chr):
// 	// 	if !lex.keyword("package") {
// 	// 		return lex.errorf("expected 'package' keyword")
// 	// 	}

// 	// 	if chr := lex.next(); chr != ' ' {
// 	// 		return lex.errorf("expected ' ' after package keyword, got: '%s'", string(chr))
// 	// 	}

// 	// 	lex.ignore() // ignore whitespace
// 	// 	lex.emit(PACKAGE)

// 	// 	return lexImports
// 	// default:
// 	// 	return lex.errorf("expected 'package', comment or whitespace, got: '%v'", chr)
// 	// }
// 	return nil
// }
