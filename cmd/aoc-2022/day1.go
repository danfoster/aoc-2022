package main

import (
	"bufio"
	"os"
	"strconv"
	"github.com/danfoster/aoc-2022/internal/logger"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Elf struct {
	Calories int
}

func read_elves(filename string) []Elf {
    readFile, err := os.Open(filename)
    check(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
	var elves = []Elf{}
	var elf = Elf{}
	
    for fileScanner.Scan() {
		line := fileScanner.Text()
		if (line == "") {
			elves = append(elves, elf)
			elf = Elf{}
		} else {
			line_int, err := strconv.Atoi(line)
			check(err)
			elf.Calories = elf.Calories +  line_int
		}
        
		
		
    }
  
    readFile.Close()
	return elves
}

func main() {

	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	
	elves := read_elves(os.Args[1])
	ceiling := 9999999
	total := 0
	for i:=0; i<3; i++ {
		biggest_index := 0
		biggest_value := 0
		for index, elf := range elves {
			if elf.Calories > biggest_value && elf.Calories < ceiling {
				biggest_index = index
				biggest_value = elf.Calories
			}
		}
		logger.Infof("%s %s", biggest_index, biggest_value)
		ceiling = biggest_value
		total += biggest_value
	}
	logger.Infof("Total: %s", total)
     
}