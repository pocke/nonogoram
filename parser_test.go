package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	r := strings.NewReader(`4 5
1 2
3
4
5
6
1 2
1
2
3`)
	b, err := Parse(r)
	if err != nil {
		t.Error(err)
	}

	expected := &Board{
		Y: 4,
		X: 5,
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

	if !reflect.DeepEqual(b, expected) {
		t.Errorf("Expected: %+v, but got %+v", expected, b)
	}
}
