package main

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

func main() {

}
