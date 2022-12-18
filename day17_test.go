package adventofcode2022

import (
	"os"
	"strings"
	"testing"
)

func TestDay17Example(t *testing.T) {
	const (
		rocks = 2022
		want  = 3068
	)
	buf, err := os.ReadFile(exampleFilename(17))
	if err != nil {
		t.Fatal(err)
	}
	got := Day17(strings.TrimSpace(string(buf)), rocks)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Part1(t *testing.T) {
	const (
		rocks = 2022
		want  = 3200
	)
	// example has only one line, puzzle input has multiple lines
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		t.Fatal(err)
	}
	got := Day17(strings.Join(lines, ""), rocks)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay17(b *testing.B) {
	const (
		rocks = 2022
	)
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		b.Fatal(err)
	}
	input := strings.Join(lines, "")
	for i := 0; i < b.N; i++ {
		Day17(input, rocks)
	}
}

/*
func TestDay17Part2(t *testing.T) {
	// looking for cycle
	// len(shapes) * len(pattern):
	// 50455, height 79972.000000  ->
	// 100910, height 159945.000000
	const (
		rocks = 1000000000000
		// want  = 1514285714288
		want = 1585016351204 // too high
	)
	// example has only one line, puzzle input has multiple lines
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		t.Fatal(err)
	}
	got := Day17(strings.Join(lines, ""), rocks)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/

/*
1000 rocks, height 1577.000000
2000 rocks, height 3166.000000
3000 rocks, height 4752.000000
4000 rocks, height 6337.000000
5000 rocks, height 7932.000000
6000 rocks, height 9511.000000
7000 rocks, height 11087.000000
8000 rocks, height 12672.000000
9000 rocks, height 14266.000000
10000 rocks, height 15843.000000
11000 rocks, height 17427.000000
12000 rocks, height 19030.000000
13000 rocks, height 20610.000000
14000 rocks, height 22194.000000
15000 rocks, height 23770.000000
16000 rocks, height 25367.000000
17000 rocks, height 26950.000000
18000 rocks, height 28531.000000
19000 rocks, height 30118.000000
20000 rocks, height 31693.000000
21000 rocks, height 33276.000000
22000 rocks, height 34860.000000
23000 rocks, height 36445.000000
24000 rocks, height 38050.000000
25000 rocks, height 39635.000000
26000 rocks, height 41203.000000
27000 rocks, height 42787.000000
28000 rocks, height 44382.000000
29000 rocks, height 45959.000000
30000 rocks, height 47544.000000
31000 rocks, height 49139.000000
32000 rocks, height 50722.000000
33000 rocks, height 52297.000000
34000 rocks, height 53890.000000
35000 rocks, height 55470.000000
36000 rocks, height 57056.000000
37000 rocks, height 58641.000000
38000 rocks, height 60229.000000
39000 rocks, height 61802.000000
40000 rocks, height 63391.000000
41000 rocks, height 64973.000000
42000 rocks, height 66558.000000
43000 rocks, height 68168.000000
44000 rocks, height 69747.000000
45000 rocks, height 71327.000000
46000 rocks, height 72894.000000
47000 rocks, height 74491.000000
48000 rocks, height 76080.000000
49000 rocks, height 77659.000000
50000 rocks, height 79257.000000
51000 rocks, height 80831.000000
52000 rocks, height 82413.000000
53000 rocks, height 84006.000000
54000 rocks, height 85588.000000
55000 rocks, height 87181.000000
56000 rocks, height 88761.000000
57000 rocks, height 90339.000000
58000 rocks, height 91923.000000
59000 rocks, height 93508.000000
60000 rocks, height 95083.000000
61000 rocks, height 96669.000000
62000 rocks, height 98280.000000
63000 rocks, height 99857.000000
64000 rocks, height 101444.000000
65000 rocks, height 103018.000000
66000 rocks, height 104613.000000
67000 rocks, height 106189.000000
68000 rocks, height 107776.000000
69000 rocks, height 109369.000000
70000 rocks, height 110937.000000
71000 rocks, height 112526.000000
72000 rocks, height 114112.000000
73000 rocks, height 115697.000000
74000 rocks, height 117292.000000
75000 rocks, height 118871.000000
76000 rocks, height 120447.000000
*/
