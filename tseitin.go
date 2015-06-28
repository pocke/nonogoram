package main

import "strconv"

type Var struct {
	Sign    bool
	Literal string
}

func Tseitin(e Expr) [][]Var {
	res, v := e.tseitin()
	return append(res, []Var{{Sign: true, Literal: v}})
}

// dummy
func (_ *EmptyExpr) tseitin() ([][]Var, string) { panic("") }

func (e *Literal) tseitin() ([][]Var, string) {
	return [][]Var{}, e.Literal
}

func (e *NotExpr) tseitin() ([][]Var, string) {
	x := rengbang()
	res, y := e.Right.tseitin()
	return append(res,
		[]Var{{Sign: false, Literal: x}, {Sign: false, Literal: y}},
		[]Var{{Sign: true, Literal: x}, {Sign: true, Literal: y}},
	), x
}

func (e *BinExpr) tseitin() ([][]Var, string) {
	x := rengbang()
	resR, y := e.Right.tseitin()
	resL, z := e.Left.tseitin()
	res := append(resR, resL...)

	switch e.Op {
	case '|':
		return append(res,
			[]Var{{Sign: false, Literal: y}, {Sign: true, Literal: x}},
			[]Var{{Sign: false, Literal: z}, {Sign: true, Literal: x}},
			[]Var{{Sign: false, Literal: x}, {Sign: true, Literal: y}, {Sign: true, Literal: z}},
		), x
	case '&':
		return append(res,
			[]Var{{Sign: false, Literal: x}, {Sign: true, Literal: y}},
			[]Var{{Sign: false, Literal: x}, {Sign: true, Literal: z}},
			[]Var{{Sign: false, Literal: y}, {Sign: false, Literal: z}, {Sign: true, Literal: x}},
		), x
	default:
		panic("")
	}
}

// "1", "2", "3", "4", ...
var rengbang = func() func() string {
	i := 1
	ch := make(chan string, 0)
	go func() {
		for {
			ch <- strconv.Itoa(i)
			i++
		}
	}()

	return func() string { return <-ch }
}()
