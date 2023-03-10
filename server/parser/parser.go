package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thervieu/42n-puzzle/models"
)

func help() {
	fmt.Printf("You must set search, heuristic options and the filename to read the puzzle from\n\n")
	fmt.Println("Available Options:")
	fmt.Println("  search: uniform or greedy")
	fmt.Println("  heuristic: hamming, manhattan or linear_conflict")
	fmt.Printf("  filename: must end with .txt and be a valid file.\n  You can generate files by running the python script\n")
	os.Exit(0)
}

func ParseArgs(args []string) (search string, heuristic string, fileName string, err error) {

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			help() // prints and quits
		}

		if strings.Contains(arg, "=") == false ||
			len(strings.Split(arg, "=")) != 2 {
			return "", "", "", errors.New("error: wrong argument '" + arg + "'\nrun program with --help for correct arguments")
		}

		if strings.HasPrefix(arg, "--search=") &&
			len(strings.Split(arg, "=")) == 2 {

			if strings.Split(arg, "=")[1] != "uniform" &&
				strings.Split(arg, "=")[1] != "greedy" {

				return "", "", "", errors.New("error: wrong argument for search '" + strings.Split(arg, "=")[1] + "'")
			}
			search = strings.Split(arg, "=")[1]

		} else if strings.HasPrefix(arg, "--heuristic=") &&
			len(strings.Split(arg, "=")) == 2 {

			if strings.Split(arg, "=")[1] != "hamming" &&
				strings.Split(arg, "=")[1] != "manhattan" &&
				strings.Split(arg, "=")[1] != "linear_conflict" {

				return "", "", "", errors.New("error: wrong argument for heuristic '" + strings.Split(arg, "=")[1] + "'")
			}

			heuristic = strings.Split(arg, "=")[1]

		} else if strings.HasPrefix(arg, "--filename=") &&
			len(strings.Split(arg, "=")) == 2 &&
			len(strings.Split(strings.Split(arg, "=")[1], ".")) == 2 {

			if strings.HasSuffix(arg, ".txt") == false {
				return "", "", "", errors.New("error: wrong argument for file '" + strings.Split(arg, "=")[1] + "'. Should end with .txt")
			}

			fileName = strings.Split(arg, "=")[1]

		} else {

			return "", "", "", errors.New("error: wrong argument '" + arg + "'\nrun program with --help for correct arguments")
		}
	}

	if search == "" || heuristic == "" || fileName == "" {
		return "", "", "", errors.New("arguments: some arguments were not given. See help.")
	}

	err = nil
	return
}
func isStringNumerical(word string) bool {

	for _, c := range word {
		if string(c) < "0" || string(c) > "9" {
			return false
		}
	}
	return true
}

func contains(arr []int, number int) bool {
	for _, nb := range arr {
		if nb == number {
			return true
		}
	}
	return false
}

func parseSize(strSplitted []string) (int, error) {

	if len(strSplitted) > 1 { // too many arguments
		return 0, errors.New("size has too many arguments")
	}

	if isStringNumerical(strSplitted[0]) == false {
		return 0, errors.New("error: file: " + strSplitted[0] + " is not numeral")
	}

	number, err := strconv.Atoi(strSplitted[0])
	if err != nil {
		return 0, errors.New(fmt.Sprintf("err in strconv.Atoi : %T \n %v\n", number, number))
	}
	if number < 3 {
		return 0, errors.New("only puzzles of size striclty superior to 2 are considered")
	}

	return number, nil
}

func concatStrArr(strSplitted []string) (str string) {
	str = ""
	for _, word := range strSplitted {
		str += word + " "
	}
	return str[0 : len(str)-1]
}

func parsePuzzleArray(size int, strSplitted []string) ([]int, error) {
	arr := []int{}

	if size != len(strSplitted) {
		return []int{}, errors.New("error: file: line: '" + concatStrArr(strSplitted) + "' number of integers (" + strconv.Itoa(len(strSplitted)) + ") differs from  size (" + strconv.Itoa(size) + ")")
	}
	for _, word := range strSplitted {
		if isStringNumerical(word) == false {
			return []int{}, errors.New("error: file: " + word + " is not numeral")
		}

		number, err := strconv.Atoi(word)
		if err != nil {
			return []int{}, errors.New(fmt.Sprintf("err in strconv.Atoi : %T \n %v\n", number, number))
		}

		if number >= size*size {
			return []int{}, errors.New("error: file: number " + strconv.Itoa(number) + " is too big for a " + strconv.Itoa(size) + "-puzzle")
		}

		arr = append(arr, number)
	}

	return arr, nil
}

func ParseFile(fileName string) (int, []int, error) {
	readFile, err := os.Open(fileName)

	if err != nil {
		return 0, []int{}, err
	}

	size := 0

	fmt.Println(fileName)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	puzzleArray := []int{}
	for fileScanner.Scan() {
		str := fileScanner.Text()
		fmt.Println(str)
		if str == "" {
			continue
		}

		if strings.HasPrefix(str, "#") { // if comment line
			continue
		}
		str = strings.Split(str, "#")[0] // line without comment
		strSplitted := strings.Fields(str)
		fmt.Println(strSplitted)
		if size == 0 { // set size on first
			if len(strSplitted) == 0 {
				continue
			}

			number, err := parseSize(strSplitted)
			if err != nil {
				readFile.Close()
				return 0, []int{}, err
			}
			fmt.Println(number)
			size = number
			continue
		}

		arr, err := parsePuzzleArray(size, strSplitted)
		if err != nil {
			readFile.Close()
			return 0, []int{}, err
		}
		puzzleArray = append(puzzleArray, arr...)
	}

	if size*size != len(puzzleArray) {
		readFile.Close()
		return 0, []int{}, errors.New("error: file: has " + strconv.Itoa(len(puzzleArray)) + " numbers, need " + strconv.Itoa(size*size))
	}

	readFile.Close()

	err = models.VerifValuesSlice(puzzleArray, false)

	if err != nil {
		return 0, []int{}, err

	}
	return size, puzzleArray, nil
}
