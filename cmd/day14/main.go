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

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input, max_x, max_y := readInput(os.Args[1])

	// Part1
	cave := genCave(input, max_x, max_y, false)
	var i int
	for i = 0; dropSand(cave, Point{x: 500, y: 0}); i++ {
	}
	// displayCave(cave, 400)
	fmt.Printf("Part 1: %d\n", i)

	// Part 2
	cave = genCave(input, max_x+max_y, max_y, true)
	for i = 0; dropSand(cave, Point{x: 500, y: 0}); i++ {
	}
	// displayCave(cave, 400)
	fmt.Printf("Part 2: %d\n", i+1)
}

type Point struct {
	x int
	y int
}

func genCave(lines *([][]Point), width int, height int, add_floor bool) *([][]int) {
	if add_floor {
		height += 2
	}
	cave := make([][]int, height)
	for y := 0; y < height; y++ {
		cave[y] = make([]int, width)
	}

	for _, line := range *lines {
		for i := 0; i < len(line)-1; i++ {
			y := line[i].y
			x := line[i].x
			cave[y][x] = 1
			for y != line[i+1].y || x != line[i+1].x {

				if y < line[i+1].y {
					y++
				} else if y > line[i+1].y {
					y--
				} else if x < line[i+1].x {
					x++
				} else if x > line[i+1].x {
					x--
				}
				cave[y][x] = 1
			}
		}
	}

	if add_floor {
		for x := 0; x < width; x++ {
			cave[height-1][x] = 1
		}
	}

	return &cave
}

func displayCave(cave *([][]int), x_offset int) {
	for y := 0; y < len(*cave); y++ {
		for x := x_offset; x < len((*cave)[0]); x++ {
			v := (*cave)[y][x]
			switch v {
			case 1:
				// Wall
				fmt.Printf("▓")
			case 2:
				// Sand
				fmt.Printf("*")
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func dropSand(cave *([][]int), pos Point) bool {
	source_x := pos.x
	source_y := pos.y
	height := len(*cave)
	// displayCave(cave, 450)
	for pos.y < height-1 {
		if (*cave)[pos.y+1][pos.x] == 0 {
			pos.y++
		} else if (*cave)[pos.y+1][pos.x-1] == 0 {
			pos.y++
			pos.x--
		} else if (*cave)[pos.y+1][pos.x+1] == 0 {
			pos.y++
			pos.x++
		} else {
			(*cave)[pos.y][pos.x] = 2
			if pos.x == source_x && pos.y == source_y {
				return false
			}
			return true
		}
	}
	return false
}
func readInput(filename string) (*([][]Point), int, int) {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	max_x := 0
	max_y := 0
	paths := [][]Point{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " -> ")
		points := []Point{}
		for _, v := range parts {
			pos := strings.Split(v, ",")
			x, err := strconv.Atoi(pos[0])
			check(err)
			y, err := strconv.Atoi(pos[1])
			check(err)
			if x > max_x {
				max_x = x
			}
			if y > max_y {
				max_y = y
			}
			p := Point{x: x, y: y}
			points = append(points, p)
		}
		paths = append(paths, points)

	}

	return &paths, max_x + 1, max_y + 1
}
