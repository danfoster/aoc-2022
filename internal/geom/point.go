package geom

import (
	"fmt"

	"github.com/danfoster/aoc-2022/internal/math"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Display() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

func (p *Point) Distance(p2 Point) int {
	d := Point{}
	d.X = math.AbsInt(p.X - p2.X)
	d.Y = math.AbsInt(p.Y - p2.Y)
	return d.X + d.Y
}
