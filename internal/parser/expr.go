package parser

import (
	"github.com/advanderveer/gor/internal/ast"
	"github.com/advanderveer/gor/internal/token"
)

// The result may be a type or even a raw type ([...]int).
func (p *Parser) parseExpr() ast.Expr {
	return p.parseBinaryExpr(nil, token.LowestPrec+1)
}

// parseBinaryExpr parses a (possibly) binary expression.
// If x is non-nil, it is used as the left operand.
func (p *Parser) parseBinaryExpr(x ast.Expr, prec1 int) ast.Expr {
	if x == nil {
		x = p.parseUnaryExpr()
	}

	for {
		tok, oprec := p.tok, p.tok.Precedence()
		if oprec < prec1 {
			return x
		}

		pos := p.expect(tok)
		y := p.parseBinaryExpr(nil, oprec+1)
		x = &ast.BinaryExpr{
			X:     x,
			OpPos: pos,
			Op:    tok,
			Y:     y,
		}
	}
}

// parseUnaryExpr parses an expression with one operant.
func (p *Parser) parseUnaryExpr() ast.Expr {
	switch p.tok {
	case token.ADD, token.SUB, token.NOT, token.XOR, token.AND:
		pos, op := p.pos, p.tok
		p.next()
		x := p.parseUnaryExpr()

		return &ast.UnaryExpr{OpPos: pos, Op: op, X: x}
	case token.MUL:
		// pointer type or unary "*" expression
		pos := p.pos
		p.next()
		x := p.parseUnaryExpr()

		return &ast.StarExpr{Star: pos, X: x}
	}

	return p.parsePrimaryExpr(nil)
}

func (p *Parser) parsePrimaryExpr(x ast.Expr) ast.Expr {
	if x == nil {
		x = p.parseOperand()
	}

	for {
		switch p.tok {
		case token.PERIOD:
			panic("PERIOD: not implemented")
		case token.LBRACK:
			panic("LBRACK: not implemented")
		case token.LPAREN:
			panic("LPAREN: not implemented")
		case token.LBRACE:
			panic("LBRACE: not implemented")
		default:
			return x
		}
	}
}

// parseOperand may return an expression or a raw type (incl. array
// types of the form [...]T). Callers must verify the result.
func (p *Parser) parseOperand() ast.Expr {
	switch p.tok {
	case token.IDENT:
		x := p.parseIdent()

		return x
	case token.INT, token.FLOAT, token.STRING:
		x := &ast.BasicLit{ValuePos: p.pos, Kind: p.tok, Value: p.lit}
		p.next()

		return x
	case token.LPAREN:
		panic("LPAREN: not implemented")
	case token.FUNC:
		panic("FUNC: not implemented")
	}

	// we have an error
	pos := p.pos
	p.errorExpected(pos, "operand")
	p.advance(stmtStart)

	return &ast.BadExpr{From: pos, To: p.pos}
}
