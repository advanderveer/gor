package ast

import (
	gotoken "go/token"

	"github.com/advanderveer/gor/internal/token"
)

// Decl declaration node int the AST.
type Decl interface {
	isDecl()
}

// BadDecl represents a declaration that couldn't be parsed.
type BadDecl struct {
	From, To gotoken.Pos
}

func (d *BadDecl) isDecl() {}

// A GenDecl node (generic declaration node) represents an import,
// constant, type or variable declaration.
type GenDecl struct {
	TokPos gotoken.Pos // position of Tok
	Tok    token.Token // IMPORT, CONST, TYPE, or VAR
	Lparen gotoken.Pos // position of '(', if any
	Specs  []Spec
	Rparen gotoken.Pos // position of ')', if any
}

func (d *GenDecl) isDecl() {}
