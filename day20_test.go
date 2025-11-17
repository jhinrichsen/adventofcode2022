package adventofcode2022

import (
	"testing"
)

func TestDay20Part1Example(t *testing.T) {
	const want = 3
	ns, err := numbersFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(ns, 1, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part1(t *testing.T) {
	const want = 11616
	ns, err := numbersFromFilename(filename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(ns, 1, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

/*
func TestDay20Part2Example(t *testing.T) {
	const want = 58
	lines, err := linesFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part2(t *testing.T) {
	const want = 3304 // too high
	lines, err := linesFromFilename(filename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/
