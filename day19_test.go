package adventofcode2022

import "testing"

func TestDay19Part1Example(t *testing.T) {
	const want uint = 33
	lines, err := linesFromFilename(exampleFilename(19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part1(t *testing.T) {
	const want uint = 1356
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2ExampleBlueprints(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(19))
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		blueprintIdx int
		want         int
	}{
		{0, 56}, // Blueprint 1: 56 geodes in 32 minutes
		{1, 62}, // Blueprint 2: 62 geodes in 32 minutes
	}

	for _, tt := range tests {
		t.Run("Blueprint_"+string(rune('1'+tt.blueprintIdx)), func(t *testing.T) {
			bp := NewBlueprint(lines[tt.blueprintIdx])
			got := bp.maxGeodes(32)
			if got != tt.want {
				t.Errorf("Blueprint %d: want %d geodes, got %d", tt.blueprintIdx+1, tt.want, got)
			}
		})
	}
}

func TestDay19Part2Example(t *testing.T) {
	const want uint = 56 * 62 // Product of first 2 blueprints (only 2 in example)
	lines, err := linesFromFilename(exampleFilename(19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2(t *testing.T) {
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(lines, false)
	if got == 0 {
		t.Fatal("got 0, expected non-zero result")
	}
	t.Logf("Part 2 result: %d", got)
}

func BenchmarkDay19Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day19(lines, true)
	}
}

func BenchmarkDay19Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day19(lines, false)
	}
}
