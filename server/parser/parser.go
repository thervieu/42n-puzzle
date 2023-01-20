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

func contains(arr []int, number int) (bool) {
	for _, nb := range arr {
		if (nb == number) {
			return true
		}
	}
	return false
}


func parseFile(fileName string) ([]int, error) {
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
		strSplitted := strings.Fields(str) //

		if (len(strSplitted) <= 2) {
			return []int{}, errors.New("line: " + str + " has two numbers or less")
		}

		if (size == 0) { // set size on first
			size = len(strSplitted)
		} else if (size != len(strSplitted)) {
			return []int{}, errors.New("line: " + str + " does not have the right size")
		}

		for _, word := range strSplitted {
			for _, c := range word {
				if (string(c) < "0" || string(c) > "9") {
					return []int{}, errors.New("word: " + word + " is not numeral")
				}
			}
			number, err := strconv.Atoi(word)
			if err != nil {
				fmt.Printf("err in strconv.Atoi : %T \n %v\n", number, number)
			}

			if (number >= size * size) {
				return []int{}, errors.New("number " + strconv.Itoa(number)  + " is too big for a " + strconv.Itoa(size) + "-puzzle" )

			}
			if (contains(puzzleArray, number)) {
				return []int{}, errors.New("number " + strconv.Itoa(number) + " already in array")
			}
			puzzleArray = append(puzzleArray, number)
		}
		
	}

	readFile.Close()

	return puzzleArray, nil
}
