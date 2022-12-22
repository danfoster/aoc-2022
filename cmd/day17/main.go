package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danfoster/aoc-2022/internal/common"
)

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	day17(os.Args[1])
}

const rockpattern = "-+L|S"

type Rock struct {
	data   [4]byte
	height uint
}

func (rock *Rock) ShiftLeft(chamber *Chamber) {
	// Check we're not up against the left side
	for y := 0; y < 4; y++ {
		if (rock.data[y] & 0b01000000) > 0 {
			return
		}
	}

	// Do the actual shift left
	for y := 0; y < 4; y++ {
		rock.data[y] <<= 1
	}

	// Check for collisions
	collision := false
	for y := uint(0); y < 4; y++ {
		if rock.data[y]&chamber.data[rock.height+y] != 0 {
			collision = true
		}
	}

	if !collision {
		// All good
		return
	}

	// Collision detected, revert shift
	for y := 0; y < 4; y++ {
		rock.data[y] >>= 1
	}
}

func (rock *Rock) ShiftRight(chamber *Chamber) {
	// Check we're not up against the right side
	for y := 0; y < 4; y++ {
		if (rock.data[y] & 0b0000001) > 0 {
			return
		}
	}

	// Do the actual shift right
	for y := 0; y < 4; y++ {
		rock.data[y] >>= 1
	}

	// Check for collisions
	collision := false
	for y := uint(0); y < 4; y++ {
		if rock.data[y]&chamber.data[rock.height+y] != 0 {
			collision = true
		}
	}

	if !collision {
		// All good
		return
	}

	// Collision detected, revert shift
	for y := 0; y < 4; y++ {
		rock.data[y] <<= 1
	}

}

func (rock *Rock) MoveDown(chamber *Chamber) bool {
	if rock.height == 0 {
		return true
	}
	for y := uint(0); y < 4; y++ {
		// fmt.Printf("%d", rock.data[y]&chamber.data[rock.height+y-1])
		// fmt.Printf("%d %08b %08b\n", rock.height, rock.data[y], chamber.data[rock.height+y-1])
		if rock.data[y]&chamber.data[rock.height+y-1] != 0 {
			return true
		}
	}
	rock.height--
	return false
}

func NewRock(rocktype byte, height uint) *Rock {
	rock := Rock{height: height}

	switch rocktype {
	case '-':
		rock.data[3] = 0b00000000
		rock.data[2] = 0b00000000
		rock.data[1] = 0b00000000
		rock.data[0] = 0b00011110
	case '+':
		rock.data[3] = 0b00000000
		rock.data[2] = 0b00001000
		rock.data[1] = 0b00011100
		rock.data[0] = 0b00001000
	case 'L':
		rock.data[3] = 0b00000000
		rock.data[2] = 0b00000100
		rock.data[1] = 0b00000100
		rock.data[0] = 0b00011100
	case '|':
		rock.data[3] = 0b00010000
		rock.data[2] = 0b00010000
		rock.data[1] = 0b00010000
		rock.data[0] = 0b00010000
	case 'S':
		rock.data[3] = 0b00000000
		rock.data[2] = 0b00000000
		rock.data[1] = 0b00011000
		rock.data[0] = 0b00011000
	default:
		panic("Unknown rock type")
	}
	return &rock
}

type Chamber struct {
	data   []byte
	height uint
}

func (chamber *Chamber) GrowToHeight(h int) {
	s := h - len(chamber.data)
	if s > 0 {
		chamber.data = append(chamber.data, make([]byte, s)...)
	}
}

func (chamber *Chamber) StampRock(rock *Rock) {
	yd := uint(0)
	for y := uint(0); y < 4; y++ {
		chamber.data[rock.height+y] |= rock.data[y]
		if rock.data[y] > 0 {
			yd = y + 1
		}
	}
	h := rock.height + yd
	if h > chamber.height {
		chamber.height = h
	}

}

func (chamber *Chamber) Print() {
	height := len(chamber.data)
	fmt.Printf("\n")
	for y := height - 1; y >= 0; y-- {
		fmt.Printf("|")
		for x := 6; x >= 0; x-- {
			p := chamber.data[y] >> x & 1
			switch p {
			case 0:
				fmt.Printf(".")
			case 1:
				fmt.Printf("#")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("+-------+\n")
}

func NewChamber() *Chamber {
	chamber := Chamber{}
	chamber.data = []byte{}
	return &chamber
}

func day17(filename string) {

	jetpattern := readInput(filename)

	chamber := NewChamber()
	jet_idx := 0
	for i := 0; i < 2022; i++ {
		rock_char := rockpattern[i%len(rockpattern)]
		rock := NewRock(rock_char, chamber.height+3)
		chamber.GrowToHeight(int(rock.height) + 4)

		for settled := false; settled == false; {
			switch (*jetpattern)[jet_idx] {
			case '<':
				rock.ShiftLeft(chamber)
			case '>':
				rock.ShiftRight(chamber)
			}
			jet_idx++
			if jet_idx >= len(*jetpattern) {
				jet_idx = 0
			}
			settled = rock.MoveDown(chamber)

		}
		chamber.StampRock(rock)

	}
	chamber.Print()
	fmt.Println(chamber.height)

}

func readInput(filename string) *string {
	readFile, err := os.Open(filename)
	common.Check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	line := fileScanner.Text()

	return &line
}
