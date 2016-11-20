package main

import (
	"github.com/ConnorAtherton/tasks/cli"
	"os"
)

func main() {
	exitCode := cli.Run(os.Args[1:])
	os.Exit(exitCode)
}
