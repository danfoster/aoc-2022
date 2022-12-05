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

type Order struct {
	amount int
	from   int
	to     int
}

func NewOrder(line []string) *Order {
	order := Order{}
	var err error
	order.amount, err = strconv.Atoi(line[1])
	check(err)
	order.from, err = strconv.Atoi(line[3])
	check(err)
	order.to, err = strconv.Atoi(line[5])
	check(err)
	return &order
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	stacks, orders := readInput(os.Args[1])
	stacks2, _ := readInput(os.Args[1])
	part1(stacks, orders)
	part2(stacks2, orders)

}

func readInput(filename string) (*[][]byte, *[]Order) {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	line := fileScanner.Text()
	numstacks := (len(line) + 1) / 4
	stacks := make([][]byte, numstacks)
	readLine(line, &stacks)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line[1] == '1' {
			break
		}
		readLine(line, &stacks)

	}
	fileScanner.Scan() // Blank line
	orders := []Order{}
	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		order := NewOrder(line)
		orders = append(orders, *order)
	}

	return &stacks, &orders
}

func readLine(line string, stacks *[][]byte) {
	numstacks := len(*stacks)
	for i := 0; i < numstacks; i++ {
		c := line[(i*4)+1]
		if c != ' ' {
			(*stacks)[i] = append([]byte{c}, (*stacks)[i]...)
		}
	}
}

func part1(stacks *[][]byte, orders *[]Order) {

	var pop byte
	for _, order := range *orders {
		for i := 0; i < order.amount; i++ {
			pop = (*stacks)[order.from-1][len((*stacks)[order.from-1])-1]
			(*stacks)[order.from-1] = (*stacks)[order.from-1][:len((*stacks)[order.from-1])-1]
			(*stacks)[order.to-1] = append((*stacks)[order.to-1], pop)
		}
	}
	fmt.Printf("Part 1: ")
	for _, stack := range *stacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}
	fmt.Printf("\n")
}

func part2(stacks *[][]byte, orders *[]Order) {

	var pop []byte
	for _, order := range *orders {
		pop = (*stacks)[order.from-1][len((*stacks)[order.from-1])-order.amount:]
		(*stacks)[order.from-1] = (*stacks)[order.from-1][:len((*stacks)[order.from-1])-order.amount]
		(*stacks)[order.to-1] = append((*stacks)[order.to-1], pop...)

	}
	fmt.Printf("Part 2: ")
	for _, stack := range *stacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}
	fmt.Printf("\n")
}
