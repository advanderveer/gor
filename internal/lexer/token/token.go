// Package token describes the different tokens
package token

// Token type for set of lexical tokens.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type Token
type Token int

const (
	// ILLEGAL indicates a lexical error.
	ILLEGAL Token = iota
	// IDENT encodes an identifier.
	IDENT
)
