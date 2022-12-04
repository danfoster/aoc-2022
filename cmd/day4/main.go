package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

type ElfPair struct {
	first  *Elf
	second *Elf
}

func (ep *ElfPair) overlaps() bool {
	if ep.first.lower_section <= ep.second.upper_section &&
		ep.first.lower_section >= ep.second.lower_section {
		return true
	} else if ep.first.upper_section <= ep.second.upper_section &&
		ep.first.upper_section >= ep.second.lower_section {
		return true
	} else if ep.second.lower_section <= ep.first.upper_section &&
		ep.second.lower_section >= ep.first.lower_section {
		return true
	} else if ep.second.upper_section <= ep.first.upper_section &&
		ep.second.upper_section >= ep.first.lower_section {
		return true
	} else {
		return false
	}
}

func (ep *ElfPair) fully_contains() bool {
	if ep.first.lower_section >= ep.second.lower_section &&
		ep.first.upper_section <= ep.second.upper_section {
		return true
	} else if ep.second.lower_section >= ep.first.lower_section &&
		ep.second.upper_section <= ep.first.upper_section {
		return true
	} else {
		return false
	}
}

func NewElfPair(input string) *ElfPair {
	parts := strings.Split(input, ",")
	elf_pair := ElfPair{}
	elf_pair.first = NewElf(parts[0])
	elf_pair.second = NewElf(parts[1])
	return &elf_pair
}

type Elf struct {
	lower_section int
	upper_section int
}

func NewElf(input string) *Elf {
	parts := strings.Split(input, "-")
	elf := Elf{}
	var err error
	elf.lower_section, err = strconv.Atoi(parts[0])
	check(err)
	elf.upper_section, err = strconv.Atoi(parts[1])
	check(err)
	return &elf
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])
	part1(input)
	part2(input)
}

func readInput(filename string) []ElfPair {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var elf_pairs = []ElfPair{}
	for fileScanner.Scan() {
		elf_pair := NewElfPair(fileScanner.Text())
		elf_pairs = append(elf_pairs, *elf_pair)
	}
	return elf_pairs
}

func part1(input []ElfPair) {
	count := 0
	for _, elf_pair := range input {
		if elf_pair.fully_contains() {
			count += 1
		}
	}
	fmt.Printf("Part 1: %d\n", count)

}

func part2(input []ElfPair) {
	count := 0
	for _, elf_pair := range input {
		if elf_pair.overlaps() {
			count += 1
		}
	}
	fmt.Printf("Part 2: %d\n", count)

}
