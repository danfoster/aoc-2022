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

type Dir struct {
	name  string
	dirs  []Dir
	files []File
}

func (dir *Dir) size() int {
	size := 0
	for _, child_dir := range dir.dirs {
		size += child_dir.size()
	}
	for _, file := range dir.files {
		size += file.size
	}
	return size
}

func (dir *Dir) sum_dirs_max_size(max_size int) int {
	sum := 0
	size := dir.size()
	if size <= max_size {
		sum += size
	}

	for _, child_dir := range dir.dirs {
		sum += child_dir.sum_dirs_max_size(max_size)
	}
	return sum
}

func (dir *Dir) find_smallest(min int, current_smallest int) int {
	size := dir.size()
	if size >= min {
		if dir.size() < current_smallest {
			current_smallest = size
		}
		for _, child := range dir.dirs {
			current_smallest = child.find_smallest(min, current_smallest)
		}
	}
	return current_smallest
}

func NewDir(fileScanner *bufio.Scanner, depth int) *Dir {
	line := fileScanner.Text()

	parts := strings.Fields(line)
	// fmt.Printf("%s D %s\n", strings.Repeat(" ", depth), parts[2])
	dir := Dir{name: parts[2]}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Fields(line)
		if parts[0] != "$" {
			panic("Expected command, got: " + line)
		}
		if parts[1] == "ls" {
			more := false
			for fileScanner.Scan() {
				line := fileScanner.Text()
				parts = strings.Fields(line)

				if parts[0] == "$" {
					more = true

					break
				} else if parts[0] != "dir" {

					dir.files = append(dir.files, *NewFile(fileScanner, depth+1))
				}

			}
			if !more {
				break
			}
		}
		line = fileScanner.Text()
		parts = strings.Fields(line)
		if parts[0] != "$" {
			panic("Expected command, got: " + line)
		}
		if parts[1] == "cd" {
			if parts[2] == ".." {
				break
			}
			dir.dirs = append(dir.dirs, *NewDir(fileScanner, depth+1))
		}

	}
	return &dir
}

type File struct {
	name string
	size int
}

func NewFile(fileScanner *bufio.Scanner, depth int) *File {
	line := fileScanner.Text()
	parts := strings.Fields(line)
	name := parts[1]
	// fmt.Printf("%s F %s\n", strings.Repeat(" ", depth), name)
	size, err := strconv.Atoi(parts[0])
	check(err)
	file := File{name: name, size: size}
	return &file
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])
	part1(input)
	part2(input)

}

func readInput(filename string) *Dir {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	dir := NewDir(fileScanner, 0)
	return dir
}

func part1(root *Dir) {
	fmt.Printf("Part 1: %d\n", (*root).sum_dirs_max_size(100000))
}

func part2(root *Dir) {
	required_space := 30000000 - (70000000 - (*root).size())
	fmt.Printf("Part 2: %d\n", (*root).find_smallest(required_space, (*root).size()))
}
