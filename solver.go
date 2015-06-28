package main

//TODO
func (b *Board) Solve() {
	expr := b.toExpr()
	Tseitin(expr)
	// tseitined := expr.Tseitin()
	// solve
}

func (b *Board) toExpr() Expr {
	var root Expr = &EmptyExpr{}

	for idx, _ := range b.RuleY {
		e := b.ruleToExpr(idx, 'y')
		root = root.And(e)
	}

	for idx, _ := range b.RuleX {
		e := b.ruleToExpr(idx, 'x')
		root = root.And(e)
	}

	return root
}

//TODO
func (b *Board) ruleToExpr(idx int, xy byte) Expr {
	return nil
}
