// Package ast declares the types used to represent syntax trees for Gor packages.
package ast

import (
	gotoken "go/token"
)

// File node in the AST.
type File struct {
	Package gotoken.Pos
	Name    *Ident
	Decls   []Decl
}
