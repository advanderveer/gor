package parser

import "github.com/advanderveer/gor/internal/token"

// for advancing past errors to expression end.
var exprEnd = map[token.Token]bool{
	token.COMMA:     true,
	token.COLON:     true,
	token.SEMICOLON: true,
	token.RPAREN:    true,
	token.RBRACK:    true,
	token.RBRACE:    true,
}

// for advancing to the next statement start.
var stmtStart = map[token.Token]bool{
	token.BREAK:       true,
	token.CONST:       true,
	token.CONTINUE:    true,
	token.DEFER:       true,
	token.FALLTHROUGH: true,
	token.FOR:         true,
	token.GO:          true,
	token.GOTO:        true,
	token.IF:          true,
	token.RETURN:      true,
	token.SELECT:      true,
	token.SWITCH:      true,
	token.TYPE:        true,
	token.VAR:         true,
}

// for advancing past invalid declarations.
var declStart = map[token.Token]bool{
	token.IMPORT: true,
	token.CONST:  true,
	token.TYPE:   true,
	token.VAR:    true,
}

// advance consumes tokens until the current token p.tok
// is in the 'to' set, or token.EOF. For error recovery.
func (p *Parser) advance(to map[token.Token]bool) {
	for ; p.tok != token.EOF; p.next() {
		if to[p.tok] {
			return
		}
	}
}
