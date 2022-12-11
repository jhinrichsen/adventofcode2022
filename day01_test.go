package adventofcode2022

import "testing"

func TestDay01Part1Example(t *testing.T) {
	const want = 24000
	ns, err := linesFromFilename(exampleFilename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(ns, 1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01(t *testing.T) {
	const want = 69177
	ns, err := linesFromFilename(filename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(ns, 1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2Example(t *testing.T) {
	const want = 45000
	ns, err := linesFromFilename(exampleFilename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(ns, 3)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 207456
	ns, err := linesFromFilename(filename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(ns, 3)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	ns, err := linesFromFilename(filename(1))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day01(ns, 3)
	}
}
