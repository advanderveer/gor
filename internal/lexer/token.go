package lexer

// Token type for set of lexical tokens.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type Token
type Token int

const (
	// ILLEGAL indicates a lexical error.
	ILLEGAL Token = iota
	// EOF indicates the end of file.
	EOF
	// WS indicates a whitespace.
	WS
	// PACKAGE keyword.
	PACKAGE
	// IDENT is an identifier.
	IDENT
)
