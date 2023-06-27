package solver

import (
	"fmt"
	"testing"
)

func TestMakeGoalArray(t *testing.T) {
	tests := []struct {
		size int
		goal []int
	}{
		{
			size: 3,
			goal: []int{1, 2, 3, 8, 0, 4, 7, 6, 5},
		},
		{
			size: 4,
			goal: []int{1, 2, 3, 4, 12, 13, 14, 5, 11, 0, 15, 6, 10, 9, 8, 7},
		},
		{
			size: 5,
			goal: []int{1, 2, 3, 4, 5, 16, 17, 18, 19, 6, 15, 24, 0, 20, 7, 14, 23, 22, 21, 8, 13, 12, 11, 10, 9, 8},
		},
	}

	for i, test := range tests {
		goal := MakeGoal(test.size)
		if len(goal) != test.size*test.size {
			t.Fatalf("solve: MakeGoal: test " + fmt.Sprintf("%d", i) + ": wrong size")
		}
		for j := range goal {

			if goal[j] != test.goal[j] {
				t.Fatalf("solve: MakeGoal: test " + fmt.Sprintf("%d: ", i) + fmt.Sprintf("%d", j) + ": wrong goal")
			}
		}
	}
}

func TestIsSolvable(t *testing.T) {
	tests := []struct {
		size   int
		puzzle []int
		truth  bool
	}{
		{
			size:   3,
			puzzle: []int{1, 2, 3, 8, 0, 4, 7, 6, 5},
			truth:  true,
		},
		{
			size:   3,
			puzzle: []int{4, 1, 0, 8, 2, 7, 6, 5, 3},
			truth:  false,
		},
		{
			size:   3,
			puzzle: []int{7, 6, 8, 2, 3, 4, 5, 1, 0},
			truth:  false,
		},
		{
			size:   4,
			puzzle: []int{1, 2, 3, 4, 12, 13, 14, 5, 11, 0, 15, 6, 10, 9, 8, 7},
			truth:  true,
		},
		{
			size:   5,
			puzzle: []int{1, 2, 3, 4, 5, 16, 17, 18, 19, 6, 15, 24, 0, 20, 7, 14, 23, 22, 21, 8, 13, 12, 11, 10, 9, 8},
			truth:  true,
		},
	}

	for i, test := range tests {
		rtn_truth := IsSolvable(test.size, test.puzzle, MakeGoal(test.size))
		if test.truth != rtn_truth {
			t.Fatalf("solve: MakeGoal: test " + fmt.Sprintf("%d: ", i) + ": wrong truth: found " + fmt.Sprintf("%t: ", rtn_truth))
		}
	}
}
