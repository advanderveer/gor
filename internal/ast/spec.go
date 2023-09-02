package ast

// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.
type Spec interface {
	isSpec()
}

// ImportSpec represents a node for an import declaration.
type ImportSpec struct {
	Name *Ident    // local package name (including "."); or nil
	Path *BasicLit // import path
}

func (s *ImportSpec) isSpec() {}

// ValueSpec represents a node for a value declaration.
type ValueSpec struct {
	Names  []*Ident // value names (len(Names) > 0)
	Type   Expr     // value type; or nil
	Values []Expr   // initial values; or nil
}

func (s *ValueSpec) isSpec() {}
