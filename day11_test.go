package adventofcode2022

import (
	"reflect"
	"testing"
)

func TestParseMonkey(t *testing.T) {
	want := Monkey{ID: 2,
		Items: []float64{79, 60, 97},
		// Operation:   "new = old * old",
		Operation:   "old * old",
		DivisibleBy: 13,
		IfTrue:      1,
		IfFalse:     3,
	}

	sample := []string{
		"Monkey 2:",
		"  Starting items: 79, 60, 97",
		"  Operation: new = old * old",
		"  Test: divisible by 13",
		"    If true: throw to monkey 1",
		"    If false: throw to monkey 3",
	}

	got, err := NewMonkey(sample)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestFormula(t *testing.T) {
	const (
		key     = "x"
		formula = key + " * " + key
		value   = 13.0
		want    = value * value
	)
	m := map[string]float64{
		key: value,
	}
	got := Eval(formula, m)
	if want != got {
		t.Fatalf("want %f but got %f", want, got)
	}
}

func TestDay11Part1Example(t *testing.T) {
	const want = 10605
	lines, err := linesFromFilename(exampleFilename(11))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day11(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part1(t *testing.T) {
	const want = 78960
	lines, err := linesFromFilename(filename(11))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day11(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay11(b *testing.B) {
	lines, err := linesFromFilename(filename(11))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day11(lines, true)
	}
}

/*
func TestDay11Part2Example(t *testing.T) {
	const want = 2499999996 // too low
	lines, err := linesFromFilename(exampleFilename(11))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day11(lines, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/
