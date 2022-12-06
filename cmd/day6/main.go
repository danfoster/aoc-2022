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

func readInput(filename string) *[]byte {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	input := []byte(fileScanner.Text())
	return &input

}

func unique(input *[]byte) bool {
	for i1, v1 := range *input {
		for i2, v2 := range *input {
			if i1 != i2 && v1 == v2 {
				return false
			}
		}
	}
	return true
}

func part1(input *[]byte) {
	marker_pos := find_marker(input, 4)
	fmt.Printf("Part 1: %d\n", marker_pos)
}

func part2(input *[]byte) {
	marker_pos := find_marker(input, 14)
	fmt.Printf("Part 1: %d\n", marker_pos)
}

func find_marker(input *[]byte, size int) int {

	for i := 0; i < len(*input)-size; i++ {
		chunk := (*input)[i : i+size]
		if unique(&chunk) {
			return i + size
		}
	}
	panic("No marker found")
}
