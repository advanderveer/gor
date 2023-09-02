// Package parser implements a parser for Gor source files.
package parser

import (
	goscan "go/scanner"
	gotoken "go/token"

	"github.com/advanderveer/gor/internal/ast"
	"github.com/advanderveer/gor/internal/scanner"
	"github.com/advanderveer/gor/internal/token"
)

// Parser state.
type Parser struct {
	file    *gotoken.File
	scanner *scanner.Scanner
	errors  goscan.ErrorList

	// scanned token
	pos gotoken.Pos
	tok token.Token
	lit string
}

// Init resets the parser state so it can be re-used.
func (p *Parser) Init(fset *gotoken.FileSet, filename string, src []byte) {
	p.file = fset.AddFile(filename, -1, len(src))
	p.errors = goscan.ErrorList{}
	p.scanner = &scanner.Scanner{}
	p.scanner.Init(p.file, src, p.errors.Add)
	p.next()
}

// ParseExpr parses a single expression.
func ParseExpr(x string) (ast.Expr, error) {
	p := &Parser{}
	p.Init(gotoken.NewFileSet(), "", []byte(x))

	expr := p.parseExpr()
	p.expectSemi()
	p.expect(token.EOF)

	return expr, p.errors.Err()
}

// ParseFile parses a file.
func ParseFile(fset *gotoken.FileSet, filename string, src []byte) (*ast.File, error) {
	p := &Parser{}
	p.Init(fset, filename, src)

	return p.parseFile(), p.errors.Err()
}

// next progresses the parser.
func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}

// error repors a parsing error.
func (p *Parser) error(pos gotoken.Pos, msg string) {
	epos := p.file.Position(pos)
	p.errors.Add(epos, msg)
}

// parseFile parses a file.
func (p *Parser) parseFile() *ast.File {
	pos := p.expect(token.PACKAGE)
	ident := p.parseIdent()
	p.expectSemi()

	var decls []ast.Decl
	for p.tok != token.EOF {
		decls = append(decls, p.parseDecl(declStart))
	}

	return &ast.File{
		Name:    ident,
		Package: pos,
		Decls:   decls,
	}
}
