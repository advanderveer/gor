package ast

import (
	gotoken "go/token"

	"github.com/advanderveer/gor/internal/token"
)

// A BasicLit node describes a basic literal.
type BasicLit struct {
	ValuePos gotoken.Pos // literal position
	Kind     token.Token // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
	Value    string      // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
}
