package models

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
			goal: []int{4, 0, 1, 2, 5, 8, 7, 6, 3},
		},
		{
			size: 4,
			goal: []int{9, 0, 1, 2, 3, 7, 11, 15, 14, 13, 12, 8, 4, 5, 6, 10},
		},
		{
			size: 5,
			goal: []int{12, 0, 1, 2, 3, 4, 9, 14, 19, 24, 23, 22, 21, 20, 15, 10, 5, 6, 7, 8, 13, 18, 17, 16, 11},
		},
	}

	for i, test := range tests {
		goal := MakeGoalArray(test.size)
		if len(goal) != test.size*test.size {
			t.Fatalf("models: MakeGoal: test " + fmt.Sprintf("%d", i) + ": wrong size")
		}
		for j := range goal {

			if goal[j] != test.goal[j] {
			fmt.Println()
			fmt.Println(test.goal[j])
			fmt.Println()
			fmt.Println(goal[j])
			fmt.Println()
			t.Fatalf("models: MakeGoal: test " + fmt.Sprintf("%d: ", i) + fmt.Sprintf("%d", j) + ": wrong goal")
			}
		}
	}
}

func TestMakeGoal(t *testing.T) {
	tests := []struct {
		size int
		goal []Point
	}{
		{
			size: 3,
			goal: []Point{{1, 1}, {0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}, {1, 2}, {0, 2}, {0, 1}},
		},
		{
			size: 4,
			goal: []Point{{1, 2}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {3, 1}, {3, 2}, {3, 3}, {2, 3}, {1, 3}, {0, 3}, {0, 2}, {0, 1}, {1, 1}, {2, 1}, {2, 2}},
		},
		{
			size: 5,
			goal: []Point{{2, 2}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 3}, {0, 2}, {0, 1}, {1, 1}, {2, 1}, {3, 1}, {3, 2}, {3, 3}, {2, 3}, {1, 3}, {1, 2}},
		},
	}

	for i, test := range tests {
		goal := MakeGoal(test.size)
		if len(goal) != test.size*test.size {
			t.Fatalf("models: MakeGoal: test " + fmt.Sprintf("%d", i) + ": wrong size")
		}
		for j := range goal {

			if goal[j].x != test.goal[j].x || goal[j].y != test.goal[j].y {
			fmt.Println()
			fmt.Println(test.goal[j])
			fmt.Println()
			fmt.Println(goal[j])
			fmt.Println()
			t.Fatalf("models: MakeGoal: test " + fmt.Sprintf("%d: ", i) + fmt.Sprintf("%d", j) + ": wrong goal")
			}
		}
	}
}
