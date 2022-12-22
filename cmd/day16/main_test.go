package main

import (
	"testing"
)

func BenchmarkDay16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day16("../../inputs/day16.txt")
	}
}
