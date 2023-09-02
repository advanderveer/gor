package parser

import (
	gotoken "go/token"

	"github.com/advanderveer/gor/internal/token"
)

// expectSemi consumes a semicolon and returns the applicable line comment.
func (p *Parser) expectSemi() {
	if p.tok == token.RPAREN || p.tok == token.RBRACE {
		return // semicolon is optional before a closing ')' or '}'
	}

	switch p.tok {
	case token.COMMA:
		// permit a ',' instead of a ';' but complain
		p.errorExpected(p.pos, "';'")

		fallthrough
	case token.SEMICOLON:
		p.next()

		return
	default:
		p.errorExpected(p.pos, "';'")
		p.advance(stmtStart)
	}
}

// errorExpected errors for a token that was encountered unexpectedly.
func (p *Parser) errorExpected(pos gotoken.Pos, msg string) {
	msg = "expected " + msg

	if pos == p.pos {
		switch {
		// implicit semicolon was inserted because of a newline so the error
		// message mentions that instead.
		case p.tok == token.SEMICOLON && p.lit == "\n":
			msg += ", found newline"
		// for literals the actual value is printed.
		case p.tok.IsLiteral():
			msg += ", found " + p.lit
		// else mention the stringed of the token.
		default:
			msg += ", found '" + p.tok.String() + "'"
		}
	}

	p.error(pos, msg)
}

// expect errors if the provided token is not encountered or errors.
func (p *Parser) expect(tok token.Token) gotoken.Pos {
	pos := p.pos
	if p.tok != tok {
		p.errorExpected(pos, "'"+tok.String()+"'")
	}

	p.next() // progress

	return pos
}
