package main

import (
	"euler-go/problems"
	"flag"
	"fmt"
)

// command line options
type options struct {
	problemNum int // Id of problem to run
}

func setupOptions() options {
	// Retrieves options from command line flags
	problemNum := flag.Int("problem-number", 1, "Problem number to run")
	flag.Parse()
	return options{*problemNum}
}

func runProblem(n int) {
	// Run problem of given number
	fmt.Printf("Running problem %d\n", n)
	problems.Problem1()
}

func main() {
	opts := setupOptions()
	runProblem(opts.problemNum)
}
