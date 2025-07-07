package main

import "fmt"

type Input struct {
	numOfRuns int
	jsonType  string
}

func inputCheck() Input {
	var input Input

	fmt.Print("Enter number of runs \n")
	fmt.Scanln(&input.numOfRuns)

	fmt.Print("Enter json type \n")
	fmt.Scanln(&input.jsonType)

	return input
}

// TODO buffio
