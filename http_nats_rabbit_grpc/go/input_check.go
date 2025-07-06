package main

import "fmt"

type Input struct {
	numOfRuns int
	jsonSize  string
	jsonType  string
}

func inputCheck() Input {
	var input Input

	fmt.Print("Enter number of runs \n")
	fmt.Scanln(&input.numOfRuns)

	fmt.Print("Enter json size \n")
	fmt.Scanln(&input.jsonSize)

	fmt.Print("Enter json type \n")
	fmt.Scanln(&input.jsonType)

	return input
}

// TODO buffio
