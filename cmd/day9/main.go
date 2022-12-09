package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(e error) {
	if e != nil {
		panic(e)
	}

}

type Point struct {
	x int
	y int
}

type Order struct {
	d     Point
	count int
}

func NewOrder(dir_char byte, count int) *Order {
	var d Point
	if dir_char == 'U' {
		d.y = 1
	} else if dir_char == 'D' {
		d.y = -1
	} else if dir_char == 'L' {
		d.x = -1
	} else if dir_char == 'R' {
		d.x = 1
	} else {
		panic("Unexpected direction: " + string(dir_char))
	}
	order := Order{count: count, d: d}
	return &order
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])
	part1(input)
	part2(input)

}

func readInput(filename string) *[]Order {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	input := []Order{}
	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		count, err := strconv.Atoi(line[1])
		check(err)
		order := NewOrder(line[0][0], count)
		input = append(input, *order)
	}
	return &input
}

func move_knots(orders *[]Order, num_knots int) int {

	knots := make([]Point, num_knots)
	visited := make(map[Point]bool)

	for _, order := range *orders {
		for i := 0; i < order.count; i++ {
			knots[0].x += order.d.x
			knots[0].y += order.d.y

			for j := 1; j < len(knots); j++ {

				if Abs(knots[j-1].x-knots[j].x) < 2 && Abs(knots[j-1].y-knots[j].y) < 2 {
					// Tail does not need to move
					continue
				}
				if knots[j-1].x > knots[j].x {
					knots[j].x += 1
				} else if knots[j-1].x < knots[j].x {
					knots[j].x -= 1
				}
				if knots[j-1].y > knots[j].y {
					knots[j].y += 1
				} else if knots[j-1].y < knots[j].y {
					knots[j].y -= 1
				}

			}
			visited[knots[len(knots)-1]] = true
		}
	}
	return len(visited)

}

func part1(orders *[]Order) {
	fmt.Printf("Part 1: %d\n", move_knots(orders, 2))
}

func part2(orders *[]Order) {
	fmt.Printf("Part 2: %d\n", move_knots(orders, 10))
}
