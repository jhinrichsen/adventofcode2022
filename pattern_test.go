package adventofcode2022

import "testing"

// testLines is a generic test helper for day part tests that work directly with []string lines.
func testLines[R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	solver func([]string, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	got := solver(lines, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// benchLines is a generic benchmark helper for day part benchmarks that work directly with []string lines.
func benchLines[R any](
	b *testing.B,
	day uint8,
	part1 bool,
	solver func([]string, bool) R,
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	for b.Loop() {
		_ = solver(lines, part1)
	}
}

// testLinesErr is a test helper for solvers that return (R, error).
func testLinesErr[R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	solver func([]string, bool) (R, error),
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	got, err := solver(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// benchLinesErr is a benchmark helper for solvers that return (R, error).
func benchLinesErr[R any](
	b *testing.B,
	day uint8,
	part1 bool,
	solver func([]string, bool) (R, error),
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	for b.Loop() {
		_, _ = solver(lines, part1)
	}
}
