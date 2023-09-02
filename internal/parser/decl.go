package parser

import (
	gotoken "go/token"

	"github.com/advanderveer/gor/internal/ast"
	"github.com/advanderveer/gor/internal/token"
)

// signature for sharing the parsing logic between different declarations.
type parseSpecFunction func(keyword token.Token, idx int) ast.Spec

// parseDecl parses a declaration.
func (p *Parser) parseDecl(sync map[token.Token]bool) ast.Decl {
	var f parseSpecFunction

	switch p.tok {
	case token.IMPORT:
		f = p.parseImportSpec

	case token.CONST, token.VAR:
		panic("var/const: not implemented")

	case token.TYPE:
		panic("type: not implemented")

	case token.FUNC:
		panic("func: not implemented")

	default:
		pos := p.pos
		p.errorExpected(pos, "declaration")
		p.advance(sync)

		return &ast.BadDecl{From: pos, To: p.pos}
	}

	return p.parseGenDecl(p.tok, f)
}

// parseGenDecl represents the shared logic for parsing imports, const, var and type declarations.
func (p *Parser) parseGenDecl(keyword token.Token, f parseSpecFunction) *ast.GenDecl {
	pos := p.expect(keyword)

	var (
		lparen, rparen gotoken.Pos
		list           []ast.Spec
	)

	if p.tok == token.LPAREN {
		lparen = p.pos
		p.next()

		for idx := 0; p.tok != token.RPAREN && p.tok != token.EOF; idx++ {
			list = append(list, f(keyword, idx))
		}

		rparen = p.expect(token.RPAREN)
		p.expectSemi()
	} else {
		list = append(list, f(keyword, 0))
	}

	return &ast.GenDecl{
		TokPos: pos,
		Tok:    keyword,
		Lparen: lparen,
		Specs:  list,
		Rparen: rparen,
	}
}

// parseImportSpec parses imports as part of a general declaration.
func (p *Parser) parseImportSpec(_ token.Token, _ int) ast.Spec {
	var ident *ast.Ident

	switch p.tok {
	case token.IDENT:
		ident = p.parseIdent()
	case token.PERIOD:
		ident = &ast.Ident{NamePos: p.pos, Name: "."}
		p.next()
	}

	pos := p.pos
	path := ""

	switch {
	case p.tok == token.STRING:
		path = p.lit
		p.next()
	case p.tok.IsLiteral():
		p.error(pos, "import path must be a string")
		p.next()
	default:
		p.error(pos, "missing import path")
		p.advance(exprEnd)
	}

	p.expectSemi()

	return &ast.ImportSpec{
		Name: ident,
		Path: &ast.BasicLit{ValuePos: pos, Kind: token.STRING, Value: path},
	}
}
