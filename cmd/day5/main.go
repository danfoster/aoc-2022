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
	stacks, orders := readInput(os.Args[1])

	for _, v := range *stacks {
		fmt.Printf("%c\n", v)
	}
	fmt.Printf("%v\n", *orders)
}

func readInput(filename string) (*[][]byte, *[][]int) {
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

	orders := [][]int{}

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
