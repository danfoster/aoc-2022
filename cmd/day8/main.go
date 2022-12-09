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

func readInput(filename string) *[][]int {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	input := [][]int{}
	for fileScanner.Scan() {
		x := []int{}
		for _, c := range fileScanner.Text() {
			v := int(c - 0)
			x = append(x, v)
		}
		input = append(input, x)
	}
	return &input
}

func part1(forest *[][]int) {
	width := len((*forest)[0])
	height := len(*forest)
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if check_clear_all_dirs(forest, x, y) {
				count += 1
			}
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}

func part2(forest *[][]int) {
	width := len((*forest)[0])
	height := len(*forest)
	score := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			s := calc_score(forest, x, y)
			if s > score {
				score = s
			}
		}
	}
	fmt.Printf("Part 2: %d\n", score)
}

func check_clear_all_dirs(forest *[][]int, x int, y int) bool {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	width := len((*forest)[0])
	height := len(*forest)

	if x == 0 || y == 0 || x == width-1 || y == height-1 {
		return true
	}
	for _, v := range dirs {
		dy := v[0]
		dx := v[1]
		if check_clear_single_dir(forest, x, y, dx, dy) {
			return true
		}
	}
	return false
}

func check_clear_single_dir(forest *[][]int, x int, y int, dx int, dy int) bool {
	tree_height := (*forest)[y][x]
	width := len((*forest)[0])
	height := len(*forest)
	x += dx
	y += dy
	for x >= 0 && y >= 0 && x < width && y < height {
		if (*forest)[y][x] >= tree_height {
			return false
		}
		x += dx
		y += dy
	}
	return true
}

func calc_score(forest *[][]int, x int, y int) int {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	width := len((*forest)[0])
	height := len(*forest)

	if x == 0 || y == 0 || x == width-1 || y == height-1 {
		return 0
	}
	score := 1
	for _, v := range dirs {
		dy := v[0]
		dx := v[1]
		score *= calc_score_single_dir(forest, x, y, dx, dy)
	}
	return score
}

func calc_score_single_dir(forest *[][]int, x int, y int, dx int, dy int) int {
	tree_height := (*forest)[y][x]
	width := len((*forest)[0])
	height := len(*forest)
	x += dx
	y += dy
	score := 0
	for x >= 0 && y >= 0 && x < width && y < height {
		score += 1
		if (*forest)[y][x] >= tree_height {
			return score
		}
		x += dx
		y += dy

	}
	return score
}
