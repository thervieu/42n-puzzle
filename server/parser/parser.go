package parser

import (
	"os"
	"fmt"
	"bufio"
	"errors"
	"strings"
	"strconv"
)

func help() {
	fmt.Println("")
	os.Exit(0)
}

func wrongArg(arg string) {
	// check if split by "=" has length 1 for specific message
	fmt.Println("n-puzzle: parsing: error: wrong argument '" + arg + "'")
	fmt.Println("run program with --help for correct arguments")
	os.Exit(1)
}

func parseArgs(args []string) (search string, heuristic string, fileName string) {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			help() // prints and quits
		}
		if (strings.Contains(arg, "=") == false ||
			len(strings.Split(arg, "=")) != 2) {

			wrongArg(arg)
		}
		if (strings.HasPrefix(arg, "--search=") &&
			len(strings.Split(arg, "=")) == 2 &&
			(strings.Split(arg, "=")[1] == "open" ||
				strings.Split(arg, "=")[1] == "greedy")) {

			search = strings.Split(arg, "=")[1]

		} else if (strings.HasPrefix(arg, "--heuristic=") &&
			len(strings.Split(arg, "=")) == 2 &&
			(strings.Split(arg, "=")[1] == "euclidean" ||
			strings.Split(arg, "=")[1] == "manhattan" ||
			strings.Split(arg, "=")[1] == "hamming")) {

			heuristic = strings.Split(arg, "=")[1]
		} else if (strings.HasPrefix(arg, "--filename=") && 
			len(strings.Split(arg, "=")) == 2 &&
			len(strings.Split(strings.Split(arg, "=")[1], ".")) == 2 &&
			strings.HasSuffix(arg, ".txt")) {
			fileName = strings.Split(arg, "=")[1]
		} else {
			wrongArg(arg)
		}

	}
	return
}
func isStringNumerical(word string) (bool) {

	for _, c := range word {
		if (string(c) < "0" || string(c) > "9") {
			return false
		}
	}
	return true
}

func contains(arr []int, number int) (bool) {
	for _, nb := range arr {
		if (nb == number) {
			return true
		}
	}
	return false
}

func parseSize(strSplitted []string) (int, error) {

	if (len(strSplitted) > 1) { // too many arguments
		return 0, errors.New("size has too many arguments")
	}

	number, err := strconv.Atoi(strSplitted[0])
	if err != nil {
		return 0,  errors.New(fmt.Sprintf("err in strconv.Atoi : %T \n %v\n", number, number))
	}
	if (number < 3) {
		return 0,  errors.New("only puzzles of size striclty superior to 2 are considered")
	}

	return number, nil
}

func concatStrArr(strSplitted []string) (str string) {
	str = ""
	for _, word := range strSplitted {
		str += word
	}
	return str
}

func parsePuzzleArray(size int, p []int, strSplitted []string) ([]int, error) {

	if (size != len(strSplitted)) {
		return []int{}, errors.New("line: " + concatStrArr(strSplitted) + " size ("  + strconv.Itoa(size) + ") differs from number of integers (" + strconv.Itoa(len(strSplitted)) + ")")
	}
	for _, word := range strSplitted {
		if (isStringNumerical(word) == false) {
			return []int{}, errors.New("word: " + word + " is not numeral")
		}
	
		number, err := strconv.Atoi(word)
		if err != nil {
			return []int{}, errors.New(fmt.Sprintf("err in strconv.Atoi : %T \n %v\n", number, number))
		}
	
		if (number >= size * size) {
			return []int{}, errors.New("number " + strconv.Itoa(number)  + " is too big for a " + strconv.Itoa(size) + "-puzzle" )
		}
		if (contains(p, number)) { // number already exists
			return []int{}, errors.New("number " + strconv.Itoa(number) + " already in array")
		}
	
		p = append(p, number)
	}
	return p, nil
}

func parseFile(fileName string) (uint8, []int, error) {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	size := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	
	puzzleArray := []int{}
	for fileScanner.Scan() {
		str := fileScanner.Text()

		if (strings.HasPrefix(str, "#")) { // if comment line
			continue
		}
		str = strings.Split(str, "#")[0] // line without comment
		strSplitted := strings.Fields(str)
		

		if (size == 0) { // set size on first
			if (len(strSplitted) == 0) {
				continue
			}

			number, err := parseSize(strSplitted)
			if err != nil {
				readFile.Close()
				return 0, []int{}, err
			}
			size = number
			continue
		}
		
		arr, err := parsePuzzleArray(size, puzzleArray, strSplitted)
		if (err != nil) {
			readFile.Close()
			return 0, []int{}, err
		}
		puzzleArray = append(puzzleArray, arr...)
	}

	if (size * size != len(puzzleArray)) {
		readFile.Close()
		return 0, []int{}, errors.New("erro: file " + fileName + ": has " + strconv.Itoa(len(puzzleArray)) + " numbers, need " + strconv.Itoa(size * size))
	}

	readFile.Close()
	return uint8(size), puzzleArray, nil
}
