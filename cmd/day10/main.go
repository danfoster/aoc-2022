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

var lettersMap = map[uint32]byte{
	311928102: 'A',
	244620583: 'B',
	210797862: 'C',
	244622631: 'D',
	504405039: 'E',
	34642991:  'F',
	479626534: 'G',
	311737641: 'H',
	474091662: 'I',
	211034380: 'J',
	307399849: 'K',
	504398881: 'L',
	311737833: 'M',
	311735657: 'N',
	211068198: 'O',
	34841895:  'P',
	341091622: 'Q',
	307471655: 'R',
	243467310: 'S',
	138547359: 'T',
	211068201: 'U',
	145049137: 'V',
	318219561: 'W',
	581046609: 'X',
	138553905: 'Y',
	504434959: 'Z',
	0:         ' ',
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
	display [6][40]int8
	enabled bool
	cpu     *CPU
}

func NewCRT(enabled bool, cpu *CPU) *CRT {
	return &CRT{enabled: enabled, cpu: cpu}
}

func (crt *CRT) clockCycleToPos() Point {
	width := len(crt.display[0])
	y := int(crt.cpu.clock / width)
	x := crt.cpu.clock % width
	return Point{x: x, y: y}
}

func (crt *CRT) execute() {
	if crt.enabled {
		p := crt.clockCycleToPos()
		if p.x >= crt.cpu.x-1 && p.x <= crt.cpu.x+1 {
			crt.display[p.y][p.x] = 1
		}
	}
}

func (crt *CRT) print_display() {
	for y := 0; y < len(crt.display); y++ {
		for x := 0; x < len(crt.display[0]); x++ {
			switch crt.display[y][x] {
			case 0:
				fmt.Printf(".")
			case 1:
				fmt.Printf("#")
			}

		}
		fmt.Printf("\n")
	}
}

func (crt *CRT) decode_display() string {
	var word [8]byte
	for x := 0; x < 8; x++ {
		word[x] = crt.decode_char(x)
	}
	return string(word[:])
}

func (crt *CRT) decode_char(char_pos int) byte {
	var letter uint32
	for y := 0; y < 6; y++ {
		for x := 0; x < 5; x++ {
			if crt.display[y][x+(char_pos*5)] == 1 {
				letter |= 1 << ((y * 5) + x)
			}
		}
	}
	return lettersMap[letter]
}

type CPU struct {
	stackptr int // Current Position of program
	exectime int // Time left to execute current command
	clock    int
	x        int         // Register x
	prog     *[]ProgLine // The Program
	crt      *CRT
}

func NewCPU(prog *[]ProgLine, crt_enabled bool) *CPU {
	e := instructionTimeMap[(*prog)[0].instruction]
	c := CPU{x: 1, prog: prog, exectime: e}
	c.crt = &CRT{enabled: crt_enabled, cpu: &c}
	return &c
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
	cpu.crt.execute()
	cpu.exectime--
	cpu.clock++
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
	prog := readInput(os.Args[1])
	cpu := NewCPU(prog, true)
	total := 0

	for i := 20; i < 241; i += 40 {
		if i == 20 {
			cpu.executeMany(20)
		} else {
			cpu.executeMany(40)
		}
		signal := i * cpu.x
		total += signal
	}
	cpu.executeMany(20)
	fmt.Printf("Part 1: %v\n", total)
	fmt.Printf("Part 2: %v\n", cpu.crt.decode_display())

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
