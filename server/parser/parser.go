package parser

import (
	"os"
	"fmt"
	"strings"
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
		if (arg == "-h" || arg == "--help") {
			help() // prints and quits
		}
		if (strings.Contains(arg, "=") == false ||
			len(strings.Split(arg, "=") != 2)) {

			wrongArg(arg)
		}
		if (strings.hasPrefix(arg, "--search=") &&
			(strings.Split(arg, "=")[1] == "open" ||
			strings.Split(arg, "=")[1] == "greedy")) {
			
			search = strings.Split(arg, "=")[1]

		} else if (strings.hasPrefix(arg, "--heuristic=")&&
			(strings.Split(arg, "=")[1] == "euclidean" ||
			strings.Split(arg, "=")[1] == "manhattan" ||
			strings.Split(arg, "=")[1] == "hamming")) {
			
			heuristic = strings.Split(arg, "=")[1]
		} else if (strings.hasPrefix(arg, "--filename=" && 
			strings.hasSuffix(arg, ".txt"))) {
			fileName = strings.Split(arg, "=")[1]
		} else {
			wrongArg(arg);
		}

	}
	return
}