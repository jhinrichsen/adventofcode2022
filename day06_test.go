package adventofcode2022

import "testing"

var day06Tests = []struct {
	stream string
	part1  int
	part2  int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
}

func TestDay06ExamplesPart1(t *testing.T) {
	for _, tt := range day06Tests {
		t.Run(tt.stream, func(t *testing.T) {
			got, _ := Day06([]byte(tt.stream), true)
			if tt.part1 != got {
				t.Fatalf("want %d but got %d", tt.part1, got)
			}
		})
	}
}

func TestDay06ExamplesPart2(t *testing.T) {
	for _, tt := range day06Tests {
		t.Run(tt.stream, func(t *testing.T) {
			got, _ := Day06([]byte(tt.stream), false)
			if tt.part2 != got {
				t.Fatalf("want %d but got %d", tt.part2, got)
			}
		})
	}
}

func TestDay06Part1(t *testing.T) {
	testSolver(t, 6, filename, true, Day06, 1876)
}

func TestDay06Part2(t *testing.T) {
	testSolver(t, 6, filename, false, Day06, 2202)
}

func BenchmarkDay06Part1(b *testing.B) {
	buf := file(b, 6)
	for b.Loop() {
		_, _ = Day06(buf, true)
	}
}

func BenchmarkDay06Part2(b *testing.B) {
	buf := file(b, 6)
	for b.Loop() {
		_, _ = Day06(buf, false)
	}
}
