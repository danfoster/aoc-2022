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

type Monkey struct {
	items        []uint64
	operation    func(uint64) uint64
	test_div     uint64
	action_true  uint64
	action_false uint64
	inspections  uint64
	monkeys      *[]Monkey
}

func NewMonkey(fileScanner *bufio.Scanner, monkeys *[]Monkey) *Monkey {
	monkey := Monkey{monkeys: monkeys}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			break
		} else if strings.HasPrefix(line, "  Starting items: ") {
			items := strings.Split(line[18:], ", ")
			for _, i := range items {
				v, err := strconv.ParseUint(i, 10, 64)
				check(err)
				monkey.items = append(monkey.items, v)
			}
		} else if strings.HasPrefix(line, "  Operation: new = ") {
			monkey.operation = parse_operation(line[19:])
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			v, err := strconv.ParseUint(line[21:], 10, 64)
			check(err)
			monkey.test_div = v
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			v, err := strconv.ParseUint(line[29:], 10, 64)
			check(err)
			monkey.action_true = v
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			v, err := strconv.ParseUint(line[30:], 10, 64)
			check(err)
			monkey.action_false = v
		} else {
			panic("Cannot parse line: " + line)
		}

	}
	return &monkey
}

func (monkey *Monkey) display(index int) {
	fmt.Printf("Monkey %d: %v\n", index, monkey.items)
}

func parse_operation(input string) func(uint64) uint64 {
	parts := strings.Fields(input)
	if input == "old * old" {
		return func(i uint64) uint64 { return i * i }
	}
	v, err := strconv.ParseUint(parts[2], 10, 64)
	check(err)
	if parts[0] == "old" && parts[1] == "+" {
		return func(i uint64) uint64 { return i + v }
	} else if parts[0] == "old" && parts[1] == "*" {
		return func(i uint64) uint64 { return i * v }
	} else if parts[0] == "old" && parts[1] == "/" {
		return func(i uint64) uint64 { return i / v }
	} else if parts[0] == "old" && parts[1] == "-" {
		return func(i uint64) uint64 { return i - v }
	}

	panic("Unknown operation: " + input)
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])
	part1(input)
	input = readInput(os.Args[1])
	part2(input)
}

func readInput(filename string) *[]Monkey {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	monkeys := []Monkey{}

	for fileScanner.Scan() {
		monkey := NewMonkey(fileScanner, &monkeys)
		monkeys = append(monkeys, *monkey)
	}
	return &monkeys
}

func part1(monkeys *[]Monkey) {
	for i := 0; i < 20; i++ {

		for j := 0; j < len(*monkeys); j++ {
			monkey := &(*monkeys)[j]
			for _, item := range monkey.items {
				item = monkey.operation(item)
				item /= 3
				if item%monkey.test_div == 0 {
					(*monkey.monkeys)[monkey.action_true].items = append((*monkey.monkeys)[monkey.action_true].items, item)
				} else {
					(*monkey.monkeys)[monkey.action_false].items = append((*monkey.monkeys)[monkey.action_false].items, item)
				}
				monkey.inspections++
			}
			monkey.items = make([]uint64, 0)
		}
	}
	var max1, max2 uint64
	for j := 0; j < len(*monkeys); j++ {
		monkey := &(*monkeys)[j]
		if monkey.inspections > max1 {
			max2 = max1
			max1 = monkey.inspections
		} else if monkey.inspections > max2 {
			max2 = monkey.inspections
		}
	}
	fmt.Printf("Part 1: %d\n", max1*max2)
}

func part2(monkeys *[]Monkey) {
	var lcm uint64
	lcm = 1
	for _, monkey := range *monkeys {
		lcm *= uint64(monkey.test_div)
	}

	for i := 0; i < 10000; i++ {

		for j := 0; j < len(*monkeys); j++ {
			monkey := &(*monkeys)[j]
			for _, item := range monkey.items {
				item = monkey.operation(item)
				item = item % lcm
				if item%monkey.test_div == 0 {
					(*monkey.monkeys)[monkey.action_true].items = append((*monkey.monkeys)[monkey.action_true].items, item)
				} else {
					(*monkey.monkeys)[monkey.action_false].items = append((*monkey.monkeys)[monkey.action_false].items, item)
				}
				monkey.inspections++
			}
			monkey.items = nil
		}
	}
	var max1, max2 uint64
	for j := 0; j < len(*monkeys); j++ {
		monkey := &(*monkeys)[j]
		if monkey.inspections > max1 {
			max2 = max1
			max1 = monkey.inspections
		} else if monkey.inspections > max2 {
			max2 = monkey.inspections
		}
	}
	fmt.Printf("Part 2: %d\n", max1*max2)
}
