package adventofcode2022

import "testing"

func TestDay19Part1Example(t *testing.T) {
	testLines(t, 19, exampleFilename, true, Day19, uint(33))
}

func TestDay19Part1(t *testing.T) {
	testLines(t, 19, filename, true, Day19, uint(1356))
}

func TestDay19Part2ExampleBlueprints(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(19))

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
	testLines(t, 19, exampleFilename, false, Day19, uint(56*62))
}

func TestDay19Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(19))
	got := Day19(lines, false)
	if got == 0 {
		t.Fatal("got 0, expected non-zero result")
	}
	t.Logf("Part 2 result: %d", got)
}

func BenchmarkDay19Part1(b *testing.B) {
	benchLines(b, 19, true, Day19)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchLines(b, 19, false, Day19)
}
