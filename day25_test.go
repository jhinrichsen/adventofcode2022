package adventofcode2022

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

type SnafuDecTests []struct {
	snafu   Snafu
	decimal int
}

var day25Tests = SnafuDecTests{
	{"1", 1},
	{"2", 2},
	{"1=", 3},
	{"1-", 4},
	{"10", 5},
	{"11", 6},
	{"12", 7},
	{"2=", 8},
	{"2-", 9},
	{"20", 10},
	{"1=0", 15},
	{"1-0", 20},
	{"1=11-2", 2022},
	{"1-0---0", 12345},
	{"1121-1110-1=0", 314159265},
	{"2=-01", 976},

	{"1=-0-2", 1747},
	{"12111", 906},
	{"2=0=", 198},
	{"21", 11},
	{"2=01", 201},
	{"111", 31},
	{"20012", 1257},
	{"112", 32},
	{"1=-1=", 353},
	{"1-12", 107},
	{"12", 7},
	{"1=", 3},
	{"122", 37},
}

/*
func TestDay25AddSnafu(t *testing.T) {
	var tests = []struct {
		a, b, sum Snafu
	}{
		{"0", "0", "0"},
		{"0", "1", "1"},
		{"1", "1", "2"},
		{"2", "1", "1="},
		{"2", "2", "1-"},
		{"-", "-", "="},
		{"=", "-", "1="},
		{"=", "=", "1-"},
		{"1111", "1", "1112"},
	}
	for _, tt := range tests {
		id := fmt.Sprintf("(%s)+(%s)=(%s)", tt.a, tt.b, tt.sum)
		t.Run(id, func(t *testing.T) {
			want := tt.sum
			got := AddSnafu(tt.a, tt.b)
			if want != got {
				t.Fatalf("want (%s) + (%s) = (%s) but got (%s)",
					tt.a, tt.b, want, got)
			}
		})
	}
}

func TestDay25TestsSnafuToDec(t *testing.T) {
	for _, tt := range day25Tests {
		id := string(tt.snafu)
		t.Run(id, func(t *testing.T) {
			want := tt.decimal
			got := SnafuToDec(tt.snafu)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}
*/

func TestDayRandomAddSnafu(t *testing.T) {
	const times = 10
	for i := 0; i < times; i++ {
		j := 2 + rand.Intn(len(day25Tests)-2)
		var picks SnafuDecTests
		for k := 0; k < j; k++ {
			l := rand.Intn(len(day25Tests))
			picks = append(picks, day25Tests[l])
		}

		var ops []string
		for _, op := range picks {
			ops = append(ops, string(op.snafu))
		}
		id := fmt.Sprintf("#%2d: ", i) + strings.Join(ops, "+")

		t.Run(id, func(t *testing.T) {
			want := 0
			sum := Snafu("0")
			for _, pick := range picks {
				want += pick.decimal
				sum = AddSnafu(sum, pick.snafu)
			}
			got := SnafuToDec(sum)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay25ExampleDec(t *testing.T) {
	const want = 4890
	lines := linesFromFilename(t, exampleFilename(25))
	var got int
	for _, line := range lines {
		got += SnafuToDec(Snafu(line))
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay25Part1Example(t *testing.T) {
	const want = "2=-1=0"
	lines := linesFromFilename(t, exampleFilename(25))
	got := Day25(lines)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay25Part1(t *testing.T) {
	const want = "122-12==0-01=00-0=02"
	lines := linesFromFilename(t, filename(25))
	got := Day25(lines)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay25ReverseDec(t *testing.T) {
	lines := linesFromFilename(t, filename(25))
	var want int
	for _, line := range lines {
		want += SnafuToDec(Snafu(line))
	}
	got := SnafuToDec(Day25(lines))
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay25StraightAdd(b *testing.B) {
	lines := linesFromFilename(b, filename(25))
	for b.Loop() {
		_ = Day25(lines)
	}
}

func BenchmarkDay25SnafuToDec(b *testing.B) {
	lines := linesFromFilename(b, filename(25))
	for b.Loop() {
		var sum int
		for _, line := range lines {
			sum += SnafuToDec(Snafu(line))
		}
		_ = sum
	}
}
