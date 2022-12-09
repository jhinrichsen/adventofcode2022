package adventofcode2021

import (
	"os"
	"testing"
)

var day06Tests = []struct {
	stream string
	marker int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
}

func TestDay06Examples(t *testing.T) {
	for _, tt := range day06Tests {
		id := tt.stream
		t.Run(id, func(t *testing.T) {
			want := tt.marker
			got := Day06(tt.stream)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay06(t *testing.T) {
	const want = 1625
	buf, err := os.ReadFile(filename(6))
	if err != nil {
		t.Fatal(err)
	}
	got := Day06(string(buf))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay06(b *testing.B) {
	buf, err := os.ReadFile(filename(6))
	if err != nil {
		b.Fatal(err)
	}
	s := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day06(s)
	}
}
