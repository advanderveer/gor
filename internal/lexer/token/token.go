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

	// PACKAGE keyword.
	PACKAGE
	// IMPORT keyword.
	IMPORT
	// VAR keyword.
	VAR

	// COMMENT holds source code comments.
	COMMENT
	// LPAREN is a left parenthesis.
	LPAREN
	// RPAREN is a right parenthesis.
	RPAREN
	// COMMA character.
	COMMA
	// ASSIGN character.
	ASSIGN

	// STRING describes a string literal.
	STRING
)
