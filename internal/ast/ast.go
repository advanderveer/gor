// Package ast declares the types for representing the Gor syntax tree.
package ast

// An Ident node represents an identifier.
type Ident struct {
	Name string
}

// An File node represents a filename.
type File struct{}
