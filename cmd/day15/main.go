package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/danfoster/aoc-2022/internal/geom"
	"github.com/danfoster/aoc-2022/internal/math"
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
	input := readInput(os.Args[1])
	result := findBeaconRangesInRow(input, 2000000)
	fmt.Printf("Part 1: %d\n", result.Sum())

	search_size := 4000000
	for i := 0; i < search_size; i++ {
		result := findBeaconRangesInRow(input, i)
		if result.NumSections() != 1 {
			fmt.Printf("Part 2: %d\n", ((result.Ranges[0].End+1)*4000000)+i)
			break
		}
	}

}

func findBeaconRangesInRow(drones *([]Drone), row int) *geom.Ranges {

	ranges := geom.Ranges{}
	for _, drone := range *drones {
		overlaps := drone.overlapsRow(row)
		if overlaps {
			// fmt.Println(drone.display())
			drone.overlappingXPositions(row, &ranges)

		}

	}

	ranges.Sort()
	ranges.Compact()
	return &ranges
}

func findBeaconInRowConstrained(drones *([]Drone), row int, min int, max int) (*(geom.Point), bool) {
	set := make(map[int]bool)

	for _, drone := range *drones {
		overlaps := drone.overlapsRow(row)
		fully_overlaps := math.AbsInt(drone.drone.Y-row) >= drone.distance
		if overlaps && !fully_overlaps {
			// fmt.Println(drone.display())
			drone.overlappingXPositionsConstrained(row, &set, min, max)
		}
		// fmt.Println(fully_overlaps)
	}

	if len(set) < max {
		beacon := geom.Point{}
		beacon.Y = row
		for i := 0; i < max; i++ {
			if !set[i] {
				beacon.X = i
				break
			}
		}
		return &beacon, true
	}

	return nil, false
}

type Drone struct {
	drone    geom.Point
	beacon   geom.Point
	distance int
}

func (drone *Drone) display() string {
	return fmt.Sprintf("%s -> %s = %d", drone.drone.Display(), drone.beacon.Display(), drone.distance)
}

func (drone *Drone) overlapsRow(row int) bool {
	return math.AbsInt(drone.drone.Y-row) < drone.distance
}

func (drone *Drone) overlappingXPositions(row int, ranges *(geom.Ranges)) {
	diff := drone.distance - math.AbsInt(drone.drone.Y-row)
	min := drone.drone.X - diff
	max := drone.drone.X + diff
	ranges.Add(geom.Range{Start: min, End: max})
}

func (drone *Drone) overlappingXPositionsConstrained(row int, set *(map[int]bool), overall_min int, overall_max int) {
	diff := drone.distance - math.AbsInt(drone.drone.Y-row)
	min := drone.drone.X - diff
	max := drone.drone.X + diff
	if min < overall_min {
		min = overall_min
	}
	if max >= overall_max {
		max = overall_max - 1
	}
	for i := min; i <= max; i++ {
		(*set)[i] = true
	}
}

func NewDrone(drone geom.Point, beacon geom.Point) *Drone {
	d := Drone{drone: drone, beacon: beacon}
	d.distance = drone.Distance(beacon)
	return &d
}

func readInput(filename string) *([]Drone) {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	drones := []Drone{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		re := regexp.MustCompile("Sensor at x=([0-9-]+), y=([0-9-]+): closest beacon is at x=([0-9-]+), y=([0-9-]+)")
		matches := re.FindStringSubmatch(line)
		x1, err := strconv.Atoi(matches[1])
		check(err)
		y1, err := strconv.Atoi(matches[2])
		check(err)
		x2, err := strconv.Atoi(matches[3])
		check(err)
		y2, err := strconv.Atoi(matches[4])
		check(err)
		p1 := geom.Point{X: x1, Y: y1}
		p2 := geom.Point{X: x2, Y: y2}
		drone := NewDrone(p1, p2)
		drones = append(drones, *drone)

	}

	return &drones
}
