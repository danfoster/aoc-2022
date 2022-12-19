package geom

import (
	"sort"

	"github.com/danfoster/aoc-2022/internal/math"
)

type Range struct {
	Start int
	End   int
}

type Ranges struct {
	Ranges []Range
}

func (r *Ranges) Sort() {
	sort.SliceStable(r.Ranges, func(i, j int) bool {
		return r.Ranges[i].Start < r.Ranges[j].Start
	})
}

func (rs *Ranges) Add(i Range) {
	rs.Ranges = append(rs.Ranges, i)
}

func (rs *Ranges) Compact() {
	rs2 := []Range{}
	start := rs.Ranges[0].Start
	end := rs.Ranges[0].End
	for _, r := range rs.Ranges {
		if r.Start-1 <= end {
			start = math.MinInt(start, r.Start)
			end = math.MaxInt(end, r.End)
		} else {
			r := Range{Start: start, End: end}
			rs2 = append(rs2, r)
			start = r.Start
			end = r.End
		}
	}
	r := Range{Start: start, End: end}
	rs2 = append(rs2, r)
	rs.Ranges = rs2
}

func (rs *Ranges) Sum() int {
	sum := 0
	for _, r := range rs.Ranges {
		sum += (r.End - r.Start)
	}
	return sum
}

func (rs *Ranges) NumSections() int {
	return len(rs.Ranges)
}
