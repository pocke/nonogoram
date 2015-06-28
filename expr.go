package main

type Expr interface {
	And(Expr) Expr
	tseitin() ([][]Var, string)
}

type BinExpr struct {
	Left  Expr
	Op    byte
	Right Expr
}
type NotExpr struct{ Right Expr }
type Literal struct{ Literal string }
type EmptyExpr struct{}

func (left *BinExpr) And(right Expr) Expr { return and(left, right) }
func (left *NotExpr) And(right Expr) Expr { return and(left, right) }
func (left *Literal) And(right Expr) Expr { return and(left, right) }
func (_ *EmptyExpr) And(e Expr) Expr      { return e }

var and = func(left, right Expr) Expr {
	return &BinExpr{
		Left:  left,
		Op:    '&',
		Right: right,
	}
}

// Interface check
var _ Expr = &BinExpr{}
var _ Expr = &NotExpr{}
var _ Expr = &Literal{}
var _ Expr = &EmptyExpr{}
