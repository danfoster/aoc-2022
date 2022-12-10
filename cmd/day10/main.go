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

type Instruction int

const (
	NoOp Instruction = iota
	AddX             = iota
)

var instructionMap = map[string]Instruction{
	"addx": AddX,
}

var instructionTimeMap = map[Instruction]int{
	NoOp: 1,
	AddX: 2,
}

type Point struct {
	x int
	y int
}

type ProgLine struct {
	instruction Instruction
	args        []int
}

func NewProgLine(line string) *ProgLine {
	parts := strings.Fields(line)
	instruction := instructionMap[parts[0]]
	args := []int{}
	for i := 1; i < len(parts); i++ {
		arg, err := strconv.Atoi(parts[i])
		check(err)
		args = append(args, arg)
	}
	i := ProgLine{instruction: instruction, args: args}
	return &i
}

type CRT struct {
	display := [][]int
	const width := 40
	const width := 6
}

func (crt *CRT) clockCycleToPos(clock int) Point {
	p = Point{}
	
	return p
}

type CPU struct {
	stackptr int         // Current Position of program
	exectime int         // Time left to execute current command
	x        int         // Register x
	prog     *[]ProgLine // The Program
}

func NewCPU(prog *[]ProgLine) *CPU {
	e := instructionTimeMap[(*prog)[0].instruction]
	return &CPU{x: 1, prog: prog, exectime: e}
}

func (cpu *CPU) execute() {
	if cpu.exectime == 0 {
		line := (*cpu.prog)[cpu.stackptr]
		switch line.instruction {
		case AddX:
			cpu.x += line.args[0]
		}
		cpu.stackptr++
		line = (*cpu.prog)[cpu.stackptr]
		cpu.exectime = instructionTimeMap[line.instruction]
	}
	cpu.exectime--
}

func (cpu *CPU) executeMany(count int) {
	for i := 0; i < count; i++ {
		cpu.execute()
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])
	part1(input)

}

func readInput(filename string) *[]ProgLine {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	input := []ProgLine{}
	for fileScanner.Scan() {
		progline := NewProgLine(fileScanner.Text())
		input = append(input, *progline)
	}
	return &input
}

func part1(prog *[]ProgLine) {
	cpu := NewCPU(prog)
	total := 0

	for i := 20; i < 221; i += 40 {
		if i == 20 {
			cpu.executeMany(20)
		} else {
			cpu.executeMany(40)
		}
		signal := i * cpu.x
		total += signal
	}
	fmt.Printf("Part 1: %v\n", total)
}
