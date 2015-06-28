package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
    1
    2 3 4 5 6
1 2
  1
  2
  3

From the above board, to the following struct.

Board{
	X: 5,
	Y: 4,
	RuleX: [][]int{
		{1, 2},
		{1},
		{2},
		{3},
	},
	RuleY: [][]int{
		{1, 2},
		{3},
		{4},
		{5},
		{6},
	},
}
*/
type Board struct {
	X     int
	Y     int
	RuleX [][]int
	RuleY [][]int
}

/*
Input format
Y X
RuleY(1, 1), RuleY(1, 2)...
RuleY(2, 1), RuleY(2, 2)...
...
RuleY(X, 1), RuleY(X, 2)...
RuleX(1, 1), RuleX(1, 2)...
RuleX(2, 1), RuleX(2, 2)...
...
RuleX(Y, 1), RuleX(Y, 2)...

Example(The above)
4 5
1 2
3
4
5
6
1 2
1
2
3
*/
func Parse(r io.Reader) (*Board, error) {
	b := &Board{}
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	l, err := scanSlice(sc)
	if err != nil {
		return nil, err
	}
	if len(l) != 2 {
		return nil, fmt.Errorf("Error") // TODO: message
	}
	b.Y = l[0]
	b.X = l[1]
	b.RuleY = make([][]int, 0, b.X)
	b.RuleX = make([][]int, 0, b.Y)

	for i := 0; i < b.X; i++ {
		l, err := scanSlice(sc)
		if err != nil {
			return nil, err
		}
		b.RuleY = append(b.RuleY, l)
	}
	for i := 0; i < b.Y; i++ {
		l, err := scanSlice(sc)
		if err != nil {
			return nil, err
		}
		b.RuleX = append(b.RuleX, l)
	}
	return b, nil
}

func scanSlice(sc *bufio.Scanner) ([]int, error) {
	sc.Scan()
	l := strings.Split(sc.Text(), " ")
	res := make([]int, 0, len(l))

	for _, v := range l {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res = append(res, n)
	}

	return res, nil
}
