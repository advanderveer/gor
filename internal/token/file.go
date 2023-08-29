// Package token defines constants representing the lexical tokens of the Gor language.
package token

// File with source code.
type File struct{}

// AddLine adds the line offset for a new line.
func (File) AddLine(int) {}

// Pos encodes the position in a source file.
type Pos int
