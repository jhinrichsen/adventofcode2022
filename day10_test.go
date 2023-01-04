package adventofcode2022

import (
	"bytes"
	"io"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/otiai10/gosseract/v2"
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
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want:\n%s\ngot:\n%s\n", want, got)
	}
}
*/

func TestDay10Part2HumanIntervention(t *testing.T) {
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

func day10Crt(t *testing.T) bytes.Buffer {
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	_, crt := Day10(lines, false)

	// simulate a virtual border of dimension 1 around crt
	height := len(crt) + 2
	width := len(crt[0]) + 2
	f := func(x, y int) bool {
		y = height - 1 - y // australian

		border := x <= 0 || x >= width-1 || y <= 0 || y >= height-1
		if border {
			return false
		}
		return crt[y-1][x-1] == '#'
	}

	var pbm bytes.Buffer
	if err := WritePBM(&pbm, width, height, f); err != nil {
		t.Fatal(err)
	}
	return pbm
}

// TestDay10OCR will apply optical character recognition to the resulting CRT
// output.
// The expected result can therefore be compared against text, not against an
// image.
// OCR requires tesseract as a prerequisite. If not available, test will be
// skipped.
func TestDay10Part2Gosseract(t *testing.T) {
	t.Skip("https://github.com/otiai10/gosseract/issues/266")

	const want = "BPJAZGAP"

	pbm := day10Crt(t)
	ocr := gosseract.NewClient()
	defer ocr.Close()
	ocr.SetPageSegMode(gosseract.PSM_RAW_LINE) // --psm 13
	ocr.SetImageFromBytes(pbm.Bytes())
	got, err := ocr.Text()
	if err != nil {
		// tesseract not installed -> skip
		t.Skip("external dependency 'tesseract' not found, skipping")
	}

	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay10Part2Tesseract(t *testing.T) {
	const want = "BPJAZGAP"

	path, err := exec.LookPath("tesseract")
	if err != nil {
		t.Skip("external command 'tesseract' not installed")
	}
	cmd := exec.Command(path, "-", "-", "--psm", "13")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		pbm := day10Crt(t)
		stdin.Write(pbm.Bytes())
	}()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		t.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	buf, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	errmsg, err := io.ReadAll(stderr)
	if err != nil {
		t.Fatal(err)
	}
	if len(errmsg) > 0 {
		t.Fatalf("tesseract reports error: %q\n", errmsg)
	}
	if err := cmd.Wait(); err != nil {
		t.Fatal(err)
	}

	got := strings.TrimSpace(string(buf))
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
