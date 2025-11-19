package adventofcode2022

import (
	"testing"
)

func TestDay11Part1Example(t *testing.T) {
	const want uint = 10605
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part1(t *testing.T) {
	const want uint = 78960
	lines := linesFromFilename(t, filename(11))
	got := Day11(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// Part 2 intermediate round tests
func TestDay11Part2Round1(t *testing.T) {
	want := []uint{2, 4, 3, 6}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 1)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 1: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round20(t *testing.T) {
	want := []uint{99, 97, 8, 103}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 20)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 20: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round1000(t *testing.T) {
	want := []uint{5204, 4792, 199, 5192}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 1000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 1000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round2000(t *testing.T) {
	want := []uint{10419, 9577, 392, 10391}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 2000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 2000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round3000(t *testing.T) {
	want := []uint{15638, 14358, 587, 15593}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 3000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 3000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round4000(t *testing.T) {
	want := []uint{20858, 19138, 780, 20797}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 4000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 4000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round5000(t *testing.T) {
	want := []uint{26075, 23921, 974, 26000}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 5000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 5000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round6000(t *testing.T) {
	want := []uint{31294, 28702, 1165, 31204}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 6000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 6000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round7000(t *testing.T) {
	want := []uint{36508, 33488, 1360, 36400}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 7000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 7000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round8000(t *testing.T) {
	want := []uint{41728, 38268, 1553, 41606}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 8000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 8000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round9000(t *testing.T) {
	want := []uint{46945, 43051, 1746, 46807}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 9000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 9000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Round10000(t *testing.T) {
	want := []uint{52166, 47830, 1938, 52013}
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11Inspections(lines, 10000)
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("round 10000: monkey %d: want %d but got %d", i, want[i], got[i])
		}
	}
}

func TestDay11Part2Example(t *testing.T) {
	const want uint = 2713310158
	lines := linesFromFilename(t, exampleFilename(11))
	got := Day11(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part2(t *testing.T) {
	const want uint = 14561971968
	lines := linesFromFilename(t, filename(11))
	got := Day11(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(11))
	for range b.N {
		_ = Day11(lines, true)
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(11))
	for range b.N {
		_ = Day11(lines, false)
	}
}
