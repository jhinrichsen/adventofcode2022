package adventofcode2022

import (
	"reflect"
	"testing"
)

func TestDay10Part1Example(t *testing.T) {
	const want = 13140
	lines, err := linesFromFilename(exampleFilename(10))
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day10(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 15140
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day10(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day10(lines, true)
	}
}

/*
func TestDay10Part2Example(t *testing.T) {
	var want = []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}
	lines, err := linesFromFilename(exampleFilename(10))
	if err != nil {
		t.Fatal(err)
	}
	_, got := Day10(lines, false)
	fmt.Println(strings.Join(got, "\n"))
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want:\n%s\ngot:\n%s\n", want, got)
	}
}
*/

func TestDay10Part2(t *testing.T) {
	var want = []string{
		"###..###....##..##..####..##...##..###..",
		"#..#.#..#....#.#..#....#.#..#.#..#.#..#.",
		"###..#..#....#.#..#...#..#....#..#.#..#.",
		"#..#.###.....#.####..#...#.##.####.###..",
		"#..#.#....#..#.#..#.#....#..#.#..#.#....",
		"###..#.....##..#..#.####..###.#..#.#....",
	}
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	_, got := Day10(lines, false)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want:\n%s\ngot:\n%s\n", want, got)
	}

	/*
	  If you, like me, cannot read the garbled hashes, uncomment this
	  replacer:
	  fmt.Println(strings.Replace(strings.Join(want, "\n"),
	   	    "#", "█", -1))
	  It will then show

	  ███..███....██..██..████..██...██..███..
	  █..█.█..█....█.█..█....█.█..█.█..█.█..█.
	  ███..█..█....█.█..█...█..█....█..█.█..█.
	  █..█.███.....█.████..█...█.██.████.███..
	  █..█.█....█..█.█..█.█....█..█.█..█.█....
	  ███..█.....██..█..█.████..███.█..█.█....

	*/
}
