package main

import (
	"os"
	"fmt"
	"github.com/thervieu/42n-puzzle/server"
)

func main() {
	args := os.Args[1:]

	if (len(args) == 0) {
		fmt.Println("no args");
		server.start();
	} else {
		fmt.Println("at least one arg");

	}
}
