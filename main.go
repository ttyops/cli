package main

import (
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	parseConfig()
	runCommand()
}
