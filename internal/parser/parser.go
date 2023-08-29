// Package parser implements a parser for Gor source files.
package parser

import (
	"go/token"

	"github.com/advanderveer/gor/internal/scanner"
)

// Parser state.
type Parser struct {
	file    *token.File
	scanner scanner.Scanner
}

// Init resets the parser state so it can be re-used.
func (p *Parser) Init(fset *token.FileSet, filename string, src []byte) {
	p.file = fset.AddFile(filename, -1, len(src))

	// @TODO include error handling

	p.scanner.Init(p.file, src, nil)
}
