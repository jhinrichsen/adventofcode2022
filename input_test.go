package adventofcode2022

import (
	"errors"
	"strings"
	"testing"
)

func TestCompareEqual(t *testing.T) {
	err := compare(
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"line 2",
		}, "\n")),
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"line 2",
		}, "\n")),
	)
	die(err, t)
}

func TestCompareDifferent(t *testing.T) {
	err := compare(
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"line 2",
		}, "\n")),
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"Aller guten Dinge sind drei",
		}, "\n")))
	if err == nil {
		t.Fatalf("want error but got nil")
	}
	var ce *CompareError
	if errors.As(err, &ce) {
		if ce.line != 3 {
			t.Fatalf("want %d but got %d", 3, ce.line)
		}
	} else {
		t.Fatalf("cannot As into CompareError: %+v", err)
	}
}
