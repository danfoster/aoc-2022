package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])
	part1(input)
	part2(input)
}

func readInput(filename string) [][]byte {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var input [][]byte
	for fileScanner.Scan() {
		input = append(input, []byte(fileScanner.Text()))
	}
	return input
}

func part1(input [][]byte) {

	total := 0
	for _, line := range input {
		middle := len(line) / 2
		compartment_a := line[0:middle]
		compartment_b := line[middle:]
		item := find_first_common(compartment_a, compartment_b)
		priority := item_to_priority(item)
		total += priority
	}
	fmt.Printf("Part 1: %d\n", total)

}

func part2(input [][]byte) {
	total := 0
	for i := 0; i < len(input); i += 3 {
		item := find_first_common_trio(input[i], input[i+1], input[i+2])
		priority := item_to_priority(item)
		total += priority
	}
	fmt.Printf("Part 2: %d\n", total)
}

func item_to_priority(item byte) int {
	priority := int(item)
	if priority >= 97 {
		priority = priority - 96
	} else {
		priority = priority - 64 + 26
	}
	return priority
}

func find_first_common(a []byte, b []byte) byte {
	for _, char_a := range a {
		for _, char_b := range b {
			if char_a == char_b {
				return char_a
			}
		}
	}
	panic("No common")
}

func find_common(a []byte, b []byte) []byte {
	var common []byte
	for _, char_a := range a {
		for _, char_b := range b {
			if char_a == char_b {
				common = append(common, char_a)
			}
		}
	}
	return common
}

func find_first_common_trio(a []byte, b []byte, c []byte) byte {
	common := find_common(a, b)
	for _, char_a := range common {
		for _, char_c := range c {
			if char_a == char_c {
				return char_a
			}
		}
	}
	panic("No common")
}
