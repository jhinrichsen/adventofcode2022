package adventofcode2022

import (
	"bufio"
	"os"
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	const want = "CMZ"
	r, err := os.Open(exampleFilename(5))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	got, err := Day05(r, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = "CFFHVVHNC"
	r, err := os.Open(filename(5))
	die(err, t)
	defer r.Close()
	got, err := Day05(r, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay05Part2Example(t *testing.T) {
	const want = "MCD"
	r, err := os.Open(exampleFilename(5))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	got, err := Day05(r, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = "FSZWBPTBG"
	r, err := os.Open(filename(5))
	die(err, t)
	defer r.Close()
	got, err := Day05(r, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	bench05(b, true)
}

func BenchmarkDay05Part2(b *testing.B) {
	bench05(b, false)
}

func bench05(b *testing.B, part1 bool) {
	name := filename(5)
	st, _ := os.Stat(name)
	f, _ := os.Open(name)
	defer f.Close()
	r := bufio.NewReaderSize(f, int(st.Size()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day05(r, part1)
	}
}
