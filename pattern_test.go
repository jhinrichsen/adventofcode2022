package adventofcode2022

import "testing"

// testWithParser is a generic test helper for parser/solver pattern (parser returns error).
func testWithParser[P any, R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	puzzle, err := parser(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// testWithParserNoErr is a generic test helper for parser/solver pattern (parser doesn't return error).
func testWithParserNoErr[P any, R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	parser func([]string) P,
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	puzzle := parser(lines)
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// benchWithParser is a generic benchmark helper for parser/solver pattern (parser returns error).
func benchWithParser[P any, R any](
	b *testing.B,
	day uint8,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	for b.Loop() {
		puzzle, _ := parser(lines)
		_ = solver(puzzle, part1)
	}
}

// benchWithParserNoErr is a generic benchmark helper for parser/solver pattern (parser doesn't return error).
func benchWithParserNoErr[P any, R any](
	b *testing.B,
	day uint8,
	part1 bool,
	parser func([]string) P,
	solver func(P, bool) R,
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	for b.Loop() {
		puzzle := parser(lines)
		_ = solver(puzzle, part1)
	}
}

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
