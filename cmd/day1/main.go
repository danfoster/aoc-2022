package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Elf struct {
	Calories int
}

func read_elves(filename string) []Elf {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var elves = []Elf{}
	var elf = Elf{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			elves = append(elves, elf)
			elf = Elf{}
		} else {
			line_int, err := strconv.Atoi(line)
			check(err)
			elf.Calories = elf.Calories + line_int
		}

	}

	readFile.Close()
	return elves
}

func main() {

	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}

	elves := read_elves(os.Args[1])

	part1(elves)
	part2(elves)

}

func part1(elves []Elf) {
	biggest_value := 0
	for _, elf := range elves {
		if elf.Calories > biggest_value {
			biggest_value = elf.Calories
		}
	}
	fmt.Printf("Part 1: %d\n", biggest_value)
}

func part2(elves []Elf) {
	ceiling := 9999999
	total := 0
	for i := 0; i < 3; i++ {
		biggest_value := 0
		for _, elf := range elves {
			if elf.Calories > biggest_value && elf.Calories < ceiling {
				biggest_value = elf.Calories
			}
		}
		ceiling = biggest_value
		total += biggest_value
	}
	fmt.Printf("Part 2: %d\n", total)
}
