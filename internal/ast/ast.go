// Package ast declares the types used to represent syntax trees for Gor packages.
package ast

import (
	gotoken "go/token"
)

// Ident is an identifier.
type Ident struct {
	NamePos gotoken.Pos
	Name    string
}

// File node in the AST.
type File struct {
	Package gotoken.Pos
	Name    *Ident
	Decls   []Decl
}

// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.
type Spec interface {
	isSpec()
}
