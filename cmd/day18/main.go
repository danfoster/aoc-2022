package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/danfoster/aoc-2022/internal/common"
	"github.com/danfoster/aoc-2022/internal/geom"
)

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	day17(os.Args[1])
}

func day17(filename string) {

	cubes, min, max := readInput(filename)


	dirs := [6]geom.Point3D{
		{X: 0, Y: 1, Z: 0},
		{X: 0, Y: -1, Z: 0},
		{X: 1, Y: 0, Z: 0},
		{X: -1, Y: 0, Z: 0},
		{X: 0, Y: 0, Z: 1},
		{X: 0, Y: 0, Z: -1},
	}


	// Part 1
	count := 0
	for point, _ := range *cubes {
		for _, dir := range dirs {
			p := point.Add(dir)
			_, ok := (*cubes)[p]
			if !ok {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)


	// Part 2
	count = 0
	queue := list.New()
	queued_points := map[geom.Point3D]bool{}
	start := min
	start.X--
	queue.PushBack(start)
	queued_points[start] = true
	for queue.Len() > 0 {
		item := queue.Front()
		point := item.Value.(geom.Point3D)
		queue.Remove(item)
		for _, dir := range dirs {

			p := point.Add(dir)
			if p.X < min.X-1 || p.Y < min.Y-1 || p.Z < min.Z-1 || p.X > max.X+1 || p.Y > max.Y+1 || p.Z > max.Z+1 {
				continue
			}

			_, ok := (*cubes)[p]
			if ok {
				count++
			} else if !queued_points[p] {
				queue.PushBack(p)
				queued_points[p] = true
			}
		}
	}
	fmt.Printf("%d\n", count)

}

func readInput(filename string) (*map[geom.Point3D]bool, geom.Point3D, geom.Point3D) {
	readFile, err := os.Open(filename)
	common.Check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	points := map[geom.Point3D]bool{}

	min := geom.Point3D{X: int((^uint(0)) >> 1), Y: int((^uint(0)) >> 1), Z: int((^uint(0)) >> 1)}
	max := geom.Point3D{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		common.Check(err)
		y, err := strconv.Atoi(parts[1])
		common.Check(err)
		z, err := strconv.Atoi(parts[2])
		common.Check(err)
		if x < min.X {
			min.X = x
		}
		if y < min.Y {
			min.Y = y
		}
		if z < min.Z {
			min.Z = z
		}
		if x > max.X {
			max.X = x
		}
		if y > max.Y {
			max.Y = y
		}
		if z > max.Z {
			max.Z = z
		}
		point := geom.Point3D{X: x, Y: y, Z: z}
		points[point] = true
	}
	return &points, min, max
}
