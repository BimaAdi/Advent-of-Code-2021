package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	// read file
	dat, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(dat), "\n")

	// convert to integer
	inputInt := make([]int, len(input))
	for i, item := range input {
		data, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		inputInt[i] = data
	}

	var before int
	var after int
	numIncreased := 0
	for i := 1; i < len(inputInt); i++ {
		before = i - 1
		after = i
		if inputInt[before] <= inputInt[after] {
			numIncreased += 1
		}
	}

	fmt.Println(numIncreased)
}

func part2() {
	// read file
	dat, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(dat), "\n")

	// convert to integer
	inputInt := make([]int, len(input))
	for i, item := range input {
		data, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		inputInt[i] = data
	}

	numIncreased := 0
	for i := 0; i < len(inputInt)-3; i++ {
		window1a := i
		window1b := i + 1
		window1c := i + 2

		window2a := i + 1
		window2b := i + 2
		window2c := i + 3

		if inputInt[window1a]+inputInt[window1b]+inputInt[window1c] < inputInt[window2a]+inputInt[window2b]+inputInt[window2c] {
			numIncreased += 1
		}
	}

	fmt.Println(numIncreased)
}

func main() {
	part1()
	part2()
}
