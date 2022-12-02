package main

import (
	"bufio"
	"fmt"
	"os"
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

	part1(os.Args[1])
	part2(os.Args[1])
}

func part1(filename string) {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	shape_score := map[byte]int{
		'A': 1, // Rock
		'B': 2, // Paper
		'C': 3, // Scissors
		'X': 1, // Rock
		'Y': 2, // Paper
		'Z': 3, // Scissors
	}
	win_score := map[string]int{
		"AX": 3, // Rocks Draw
		"BY": 3, // Papers Draw
		"CZ": 3, // Scissors Draw
		"AZ": 0, // Rock beats Scissors
		"BX": 0, // Paper beats Rock
		"CY": 0, // Scissors beats Paper
		"AY": 6, // Rock loses to Paper
		"BZ": 6, // Paper loses to Scissors
		"CX": 6, // Scissors loses to Rock
	}
	var score int

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		input := line[0][0]
		output := line[1][0]
		score += shape_score[output]
		score += win_score[string(input)+string(output)]

	}
	fmt.Printf("Part 1 Total Score: %d\n", score)

}

func part2(filename string) {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	mapping_score := map[string]int{
		"AX": 3, // Rock, Lose -> Scissors: 0 + 3
		"BY": 5, // Paper, Draw -> Paper: 3 + 2
		"CZ": 7, // Scissors, Win -> Rock: 6 + 1
		"AZ": 8, // Rock, Win -> Paper: 6 + 2
		"BX": 1, // Paper, Lose -> Rock: 0 + 1
		"CY": 6, // Scissors, Draw -> Scissors: 3 + 3
		"AY": 4, // Rock, Draw -> Rock: 3 + 1
		"BZ": 9, // Paper, Win -> Scissors: 6 + 3
		"CX": 2, // Scissors, Lose -> Paper: 0 + 2
	}
	var score int

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		move := line[0][0]
		outcome := line[1][0]
		score += mapping_score[string(move)+string(outcome)]

	}
	fmt.Printf("Part 2 Total Score: %d\n", score)

}
