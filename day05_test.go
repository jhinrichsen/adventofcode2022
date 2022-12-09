package adventofcode2021

import (
	"bufio"
	"os"
	"testing"
)

func TestDay05Example(t *testing.T) {
	const want = "CMZ"
	r, err := os.Open(exampleFilename(5))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	got, err := Day05(r)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay05(t *testing.T) {
	const want = "RTGWZTHLD"
	r, err := os.Open(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	got, err := Day05(r)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay05(b *testing.B) {
	die := func(err error) {
		if err != nil {
			b.Fatal(err)
		}
	}
	name := filename(5)
	st, err := os.Stat(name)
	die(err)

	f, err := os.Open(name)
	die(err)
	defer f.Close()

	r := bufio.NewReaderSize(f, int(st.Size()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day05(r)
	}
}
