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

	commands := make([]string, len(input))
	units := make([]int, len(input))
	for i, item := range input {
		action := strings.Split(item, " ")
		commands[i] = action[0]
		units[i], err = strconv.Atoi(action[1])
		if err != nil {
			panic(err)
		}
	}

	horizontal := 0
	depth := 0
	for i := 0; i < len(input); i++ {
		if commands[i] == "forward" {
			horizontal += units[i]
		} else if commands[i] == "down" {
			depth += units[i]
		} else if commands[i] == "up" {
			depth -= units[i]
		}
	}

	fmt.Println(horizontal * depth)
}

func part2() {
	// read file
	dat, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(dat), "\n")

	commands := make([]string, len(input))
	units := make([]int, len(input))
	for i, item := range input {
		action := strings.Split(item, " ")
		commands[i] = action[0]
		units[i], err = strconv.Atoi(action[1])
		if err != nil {
			panic(err)
		}
	}

	horizontal := 0
	depth := 0
	aim := 0
	for i := 0; i < len(input); i++ {
		if commands[i] == "forward" {
			horizontal += units[i]
			depth += aim * units[i]
		} else if commands[i] == "down" {
			aim += units[i]
		} else if commands[i] == "up" {
			aim -= units[i]
		}
	}

	fmt.Println(horizontal * depth)
}

func main() {
	part1()
	part2()
}
