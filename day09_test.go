package adventofcode2022

import (
	"fmt"
	"testing"
)

func testDay09(t *testing.T, example int, tails int, want int) {
	var f = func() func(int) string {
		if example == 0 {
			return filename
		}
		return func(_ int) string {
			return fmt.Sprintf("testdata/day09_example%d.txt",
				example)
		}
	}()
	lines, err := linesFromFilename(f(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09(lines, make([]complex128, tails))
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay09Part1Example(t *testing.T) {
	testDay09(t, 1, 2, 13)
}

func TestDay09Part1(t *testing.T) {
	testDay09(t, 0, 2, 6044)
}

func BenchmarkDay09Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(9))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day09(lines, []complex128{0 + 0i})
	}
}

func TestDay09Part2Example1(t *testing.T) {
	testDay09(t, 1, 10, 1)
}

func TestDay09Part2Example2(t *testing.T) {
	testDay09(t, 2, 10, 36)
}

func TestDay09Part2(t *testing.T) {
	testDay09(t, 0, 10, 2384)
}

func BenchmarkDay09Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(9))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day09(lines, []complex128{
			0 + 0i, // 1
			0 + 0i, // 2
			0 + 0i, // 3
			0 + 0i, // 4
			0 + 0i, // 5
			0 + 0i, // 6
			0 + 0i, // 7
			0 + 0i, // 8
			0 + 0i, // 9
			0 + 0i, // 10
		})
	}
}
