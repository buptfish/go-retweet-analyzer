package main

import (
	"os"
)

func run(args []string) int {
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
