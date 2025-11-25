package adventofcode2022

import "testing"

func TestDay05Part1Example(t *testing.T) {
	testSolver(t, 5, exampleFilename, true, Day05, "CMZ")
}

func TestDay05Part1(t *testing.T) {
	testSolver(t, 5, filename, true, Day05, "CFFHVVHNC")
}

func TestDay05Part2Example(t *testing.T) {
	testSolver(t, 5, exampleFilename, false, Day05, "MCD")
}

func TestDay05Part2(t *testing.T) {
	testSolver(t, 5, filename, false, Day05, "FSZWBPTBG")
}

func BenchmarkDay05Part1(b *testing.B) {
	buf := file(b, 5)
	for b.Loop() {
		_, _ = Day05(buf, true)
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	buf := file(b, 5)
	for b.Loop() {
		_, _ = Day05(buf, false)
	}
}
