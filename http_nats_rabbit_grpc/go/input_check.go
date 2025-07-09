package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func inputCheckBuffio() Input {
	var input Input

	r := bufio.NewScanner(os.Stdin)

	for r.Scan() {
		d := r.Text()
		fmt.Println(d)
	}

	return input
}

// TODO buffio
