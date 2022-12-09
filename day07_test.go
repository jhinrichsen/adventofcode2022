package adventofcode2022

import "testing"

func TestDay07ExamplePart1(t *testing.T) {
	const want = 95437
	lines, err := linesFromFilename(exampleFilename(7))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 1243729
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07ExamplePart2(t *testing.T) {
	const want = 24933642
	lines, err := linesFromFilename(exampleFilename(7))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 4443914
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
