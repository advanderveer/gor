package ast

import (
	gotoken "go/token"

	"github.com/advanderveer/gor/internal/token"
)

// Expr represents an expression.
type Expr interface {
	isExpr()
}

// A BadExpr node is a placeholder for an expression containing
// syntax errors for which a correct expression node cannot be
// created.
type BadExpr struct {
	From, To gotoken.Pos // position range of bad expression
}

// Ident is an identifier.
type Ident struct {
	NamePos gotoken.Pos
	Name    string
}

func (*Ident) isExpr() {}

// BinaryExpr represents an expression operating on two operands.
type BinaryExpr struct {
	X     Expr        // left operand
	OpPos gotoken.Pos // position of Op
	Op    token.Token // operator
	Y     Expr        // right operand
}

// UnaryExpr represents an expression with one operand.
type UnaryExpr struct {
	X     Expr        // left operand
	OpPos gotoken.Pos // position of Op
	Op    token.Token // operator
}

// StarExpr represents an expression with one operand.
type StarExpr struct {
	X    Expr        // left operand
	Star gotoken.Pos // operator
}

func (e *BadExpr) isExpr()    {}
func (e *BinaryExpr) isExpr() {}
func (e *UnaryExpr) isExpr()  {}
func (e *StarExpr) isExpr()   {}
