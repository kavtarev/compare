package main

import "fmt"

type Input struct {
	numOfRuns int
	jsonSize  string
}

func inputCheck() Input {
	var input Input

	fmt.Print("Enter number of runs \n")
	fmt.Scanln(&input.numOfRuns)

	fmt.Print("Enter json size \n")
	fmt.Scanln(&input.jsonSize)

	return input
}

// TODO buffio
