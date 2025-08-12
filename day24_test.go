package adventofcode2022

import (
	"bytes"
	"os"
	"testing"
)

func TestDay24Rundown(t *testing.T) {
	const filename = "testdata/day24_rundown.txt"
	want, err := os.Open(filename)
	die(err, t)

	fi, err := os.Stat(filename)
	die(err, t)

	got := bytes.NewBuffer(make([]byte, fi.Size()))
	compare(want, got)
}

/*
func TestDay24Example(t *testing.T) {
	const want = 18
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/
