package parser

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		args      []string
		search    string
		heuristic string
		filename  string
		err       error
	}{
		{
			args:      []string{" "},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("error: wrong argument ' '\nrun program with --help for correct arguments"),
		},
		{
			args:      []string{"--sea", "fdzze", "fez"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("error: wrong argument '--sea'\nrun program with --help for correct arguments"),
		},
		{
			args:      []string{"--search", "fdzze", "fez"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("error: wrong argument '--search'\nrun program with --help for correct arguments"),
		},
		{
			args:      []string{"search=uniform", "fdzze", "fez"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("error: wrong argument 'search=uniform'\nrun program with --help for correct arguments"),
		},
		{
			args:      []string{"--search=uniform", "fdzze", "fez"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("error: wrong argument 'fdzze'\nrun program with --help for correct arguments"),
		},
		{
			args:      []string{"--search=uniformm", "fez"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("error: wrong argument for search 'uniformm'"),
		},
		{
			args:      []string{"--search=uniform", "fdzze", "fez"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("error: wrong argument 'fdzze'\nrun program with --help for correct arguments"),
		},
		{
			args:      []string{"--search=uniform", "--heuristic=manhattan"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("arguments: some arguments were not given. See help."),
		},
		{
			args:      []string{"--search=uniform", "--filename=aaa.txt"},
			search:    "",
			heuristic: "",
			filename:  "",
			err:       errors.New("arguments: some arguments were not given. See help."),
		},
		{
			args:      []string{"--search=uniform", "--heuristic=manhattan", "--filename=aaa.txt"},
			search:    "uniform",
			heuristic: "manhattan",
			filename:  "aaa.txt",
			err:       nil,
		},
		{
			args:      []string{"--search=greedy", "--heuristic=linear_conflict", "--filename=files/aaa.txt"},
			search:    "greedy",
			heuristic: "linear_conflict",
			filename:  "files/aaa.txt",
			err:       nil,
		},
	}

	for i, test := range tests {

		search, heur, file, err := ParseArgs(test.args)
		if search != test.search {
			t.Fatalf("parser: ParseArgs: test " + fmt.Sprintf("%d", i) + ": wrong search")
		}
		if heur != test.heuristic {
			t.Fatalf("parser: ParseArgs: test " + fmt.Sprintf("%d", i) + ": wrong heur")
		}
		if file != test.filename {
			t.Fatalf("parser: ParseArgs: test " + fmt.Sprintf("%d", i) + ": wrong file")
		}
		if (err == nil && err != test.err) || (err != nil && test.err != nil && err.Error() != test.err.Error()) {
			t.Fatalf("parser: ParseArgs: test " + fmt.Sprintf("%d", i) + ": wrong err")
		}
	}
}

// return 0, []int{}, errors.New("error: file " + fileName + ": has " + strconv.Itoa(len(puzzleArray)) + " numbers, need " + strconv.Itoa(size*size))

func TestParseFile(t *testing.T) {
	tests := []struct {
		filename string
		size     int
		array    []int
		err      error
	}{
		{
			filename: "notExist.txt",
			size:     0,
			array:    []int{},
			err:      errors.New("open notExist.txt: no such file or directory"),
		},
		{
			filename: "../files/valid_3.txt",
			size:     3,
			array:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			err:      nil,
		},
		{
			filename: "../files/valid_4.txt",
			size:     4,
			array:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			err:      nil,
		},
		{
			filename: "../files/valid_5.txt",
			size:     5,
			array:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			err:      nil,
		},
		{
			filename: "../files/valid_6.txt",
			size:     6,
			array:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35},
			err:      nil,
		},
		{
			filename: "../files/valid_7.txt",
			size:     7,
			array:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48},
			err:      nil,
		},
		{
			filename: "../files/valid_3_comments.txt",
			size:     3,
			array:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			err:      nil,
		},
		{
			filename: "../files/alpha_3.txt",
			size:     0,
			array:    []int{},
			err:      errors.New("error: file: 9a is not numeral"),
		},
		{
			filename: "../files/size_low.txt",
			size:     0,
			array:    []int{},
			err:      errors.New("only puzzles of size striclty superior to 2 are considered"),
		},
		{
			filename: "../files/size_diff_3.txt",
			size:     0,
			array:    []int{},
			err:      errors.New("error: file: line: '3 4 5 2' number of integers (4) differs from  size (3)"),
		},
		{
			filename: "../files/same_values_3.txt",
			size:     0,
			array:    []int{},
			err:      errors.New("error: file: value 1 duplicate "),
		},
		{
			filename: "../files/too_little_3.txt",
			size:     0,
			array:    []int{},
			err:      errors.New("error: file: has 6 numbers, need 9"),
		},
		{
			filename: "../files/too_many_3.txt",
			size:     0,
			array:    []int{},
			err:      errors.New("error: file: has 12 numbers, need 9"),
		},
		{
			filename: "../files/subject_1.txt",
			size:     3,
			array:    []int{3, 2, 6, 1, 4, 0, 8, 7, 5},
			err:      nil,
		},
		{
			filename: "../files/subject_2.txt",
			size:    4,
			array:    []int{0, 10, 5, 7, 11, 14, 4, 8, 1, 2, 6, 13, 12, 3, 15, 9},
			err:      nil,
		},
		{
			filename: "../files/subject_3.txt",
			size:    4,
			array:    []int{0, 10, 5, 7, 11, 14, 4, 8, 1, 2, 6, 13, 12, 3, 15, 9},
			err:      nil,
		},
	}

	for i, test := range tests {

		size, array, err := ParseFile(test.filename)
		fmt.Println(i, size, array, err)
		if size != test.size {
			t.Fatalf("parser: ParseArgs: test " + fmt.Sprintf("%d", i) + ": wrong size")
		}
		for i := range array {
			if array[i] != test.array[i] {
				t.Fatalf("parser: ParseArgs: test " + fmt.Sprintf("%d", i) + ": wrong array")
			}
		}
		if (err == nil && err != test.err) || (err != nil && test.err != nil && err.Error() != test.err.Error()) {
			t.Fatalf("parser: ParseArgs: test " + fmt.Sprintf("%d", i) + ": wrong err")
		}
	}
}
