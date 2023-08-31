package parser

import (
	"github.com/advanderveer/gor/internal/ast"
	"github.com/advanderveer/gor/internal/token"
)

func (p *Parser) parseIdent() *ast.Ident {
	pos, name := p.pos, "_"

	if p.tok == token.IDENT {
		name = p.lit
		p.next()
	} else {
		p.expect(token.IDENT) // use expect() error handling
	}

	return &ast.Ident{
		NamePos: pos,
		Name:    name,
	}
}
