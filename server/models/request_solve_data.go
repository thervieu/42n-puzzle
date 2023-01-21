package models

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type RequestSolveData struct {
	Option int   `json:"option"`
	State  []int `json:"state"`
}

func SanitizeOption(option int) (err error) {
	if option < 0 || option > 6 {
		return errors.New("Bad Request: option has a wrong value")
	}

	return nil
}

func VerifValuesSlice(arr []int, typeError bool) error {

	begError := ""
	if typeError == true {
		begError = "Bad Request: State: "
	} else {
		begError = "error: file: "
	}
	slice := make([]int, len(arr))
	for _, v := range arr {
		if v < 0 {
			return errors.New(begError + "negative number")
		}
		if v >= len(arr) {
			return errors.New(begError + "value bigger than its length")
		}
		if slice[v] == 1 {
			return errors.New(begError + "value " + strconv.Itoa(v) + " duplicate ")
		}
		slice[v] = 1
	}
	return nil
}

func (r *RequestSolveData) SanitizeState() error {

	int_sqrt := int(math.Sqrt(float64(len(r.State))))
	if int_sqrt*int_sqrt != len(r.State) {
		return errors.New("Bad Request: State length is not a squared number")
	}
	if int_sqrt*int_sqrt != len(r.State) {
		return errors.New("Bad Request: State length is not a squared number")
	}

	return VerifValuesSlice(r.State, true)
}

func (r *RequestSolveData) Sanitize() error {
	if err := SanitizeOption(r.Option); err != nil {
		return err
	}
	if err := r.SanitizeState(); err != nil {
		return err
	}
	return nil
}

func generateNumbersSlice(strPuzzle string) (slice []int) {
	splittedStr := strings.Split(strPuzzle, ",")
	slice = make([]int, len(splittedStr))
	for i, str := range splittedStr {
		_int, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		slice[i] = _int
	}
	return
}

func index(needle int, haystack []int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func (r *RequestSolveData) ExtractPuzzle() (puzzle Puzzle, err error) {

	puzzle.numbers = r.State
	puzzle.size = int(math.Sqrt(float64((len(puzzle.numbers)))))
	puzzle.parent_move = ""
	puzzle.pos_zero = index(0, puzzle.numbers)

	return
}
