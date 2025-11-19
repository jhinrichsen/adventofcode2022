package adventofcode2022

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func TestPBM(t *testing.T) {
	var (
		// From PBM documentation
		want = []string{
			"P1 24 7",
			"0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
			"0 1 1 1 1 0 0 1 1 1 1 0 0 1 1 1 1 0 0 1 1 1 1 0",
			"0 1 0 0 0 0 0 1 0 0 0 0 0 1 0 0 0 0 0 1 0 0 1 0",
			"0 1 1 1 0 0 0 1 1 1 0 0 0 1 1 1 0 0 0 1 1 1 1 0",
			"0 1 0 0 0 0 0 1 0 0 0 0 0 1 0 0 0 0 0 1 0 0 0 0",
			"0 1 0 0 0 0 0 1 1 1 1 0 0 1 1 1 1 0 0 1 0 0 0 0",
			"0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
		}

		buf bytes.Buffer
	)

	err := WritePBM(&buf, 24, 7, func(x, y int) bool {
		return want[y+1][x*2] == '1'
	})
	if err != nil {
		t.Fatal(err)
	}

	var got []string
	sc := bufio.NewScanner(&buf)
	for sc.Scan() {
		line := sc.Text()
		got = append(got, line)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v\n", want, got)
	}
}
