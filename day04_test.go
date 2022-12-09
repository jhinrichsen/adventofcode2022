package adventofcode2022

import (
	"fmt"
	"testing"
)

var containsTests = []struct {
	a1, a2, b1, b2 int
	want           bool
}{
	{1, 4, 2, 3, true},
	{1, 1, 1, 1, true},
	{1, 1, 1, 3, true},
	{1, 3, 1, 1, true},
	{1, 1, 2, 2, false},
	{2, 2, 1, 1, false},
	{0, 0, 1, 9, false},
	{1, 9, 0, 0, false},
}

func TestContains(t *testing.T) {
	for _, tt := range containsTests {
		id := fmt.Sprintf("[%d-%d] <-> [%d-%d]", tt.a1, tt.a2,
			tt.b1, tt.b2)
		t.Run(id, func(t *testing.T) {
			got := Contains(tt.a1, tt.a2, tt.b1, tt.b2)
			if tt.want != got {
				t.Fatalf("want %t but got %t", tt.want, got)
			}
		})
	}
}

func TestDay04Example(t *testing.T) {
	const want = 2
	lines, err := linesFromFilename(exampleFilename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04(t *testing.T) {
	const want = 500
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
