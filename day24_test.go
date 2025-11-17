package adventofcode2022

import "testing"

func TestDay24Part1Example(t *testing.T) {
	const want uint = 18
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	const want uint = 299
	lines, err := linesFromFilename(filename(24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// TestDay24Part1Minute0 verifies the initial blizzard configuration
func TestDay24Part1Minute0(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	v := NewValley(lines)

	// Expected initial grid
	expected := []string{
		"#.######",
		"#>>.<^<#",
		"#.<..<<#",
		"#>v.><>#",
		"#<^v^^>#",
		"######.#",
	}

	actual := v.RenderGrid(0)

	// Compare
	for y := 0; y < len(expected); y++ {
		if actual[y] != expected[y] {
			t.Errorf("Minute 0, row %d:\n  want: %s\n  got:  %s", y, expected[y], actual[y])
		}
	}

	if t.Failed() {
		t.Log("Full grid at minute 0:")
		for _, line := range actual {
			t.Log(line)
		}
	}
}

// TestDay24Part1Minute1 verifies the blizzard configuration at minute 1
func TestDay24Part1Minute1(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	v := NewValley(lines)

	// Expected grid at Minute 1 (without E marker)
	expected := []string{
		"#.######",
		"#.>3.<.#",
		"#<..<<.#",
		"#>2.22.#",
		"#>v..^<#",
		"######.#",
	}

	actual := v.RenderGrid(1)

	// Compare
	for y := 0; y < len(expected); y++ {
		if actual[y] != expected[y] {
			t.Errorf("Minute 1, row %d:\n  want: %s\n  got:  %s", y, expected[y], actual[y])
		}
	}

	if t.Failed() {
		t.Log("Full grid at minute 1:")
		for _, line := range actual {
			t.Log(line)
		}
	}
}

// TestDay24Part1Minute2 verifies the blizzard configuration at minute 2
func TestDay24Part1Minute2(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	v := NewValley(lines)

	expected := []string{
		"#.######",
		"#.2>2..#",
		"#.^22^<#",
		"#.>2.^>#",
		"#.>..<.#",
		"######.#",
	}

	actual := v.RenderGrid(2)

	for y := 0; y < len(expected); y++ {
		if actual[y] != expected[y] {
			t.Errorf("Minute 2, row %d:\n  want: %s\n  got:  %s", y, expected[y], actual[y])
		}
	}

	if t.Failed() {
		t.Log("Full grid at minute 2:")
		for _, line := range actual {
			t.Log(line)
		}
	}
}

// TestDay24Part1Minute3 verifies the blizzard configuration at minute 3
func TestDay24Part1Minute3(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	v := NewValley(lines)

	expected := []string{
		"#.######",
		"#<^<22.#",
		"#.2<.2.#",
		"#><2>..#",
		"#..><..#",
		"######.#",
	}

	actual := v.RenderGrid(3)

	for y := 0; y < len(expected); y++ {
		if actual[y] != expected[y] {
			t.Errorf("Minute 3, row %d:\n  want: %s\n  got:  %s", y, expected[y], actual[y])
		}
	}

	if t.Failed() {
		t.Log("Full grid at minute 3:")
		for _, line := range actual {
			t.Log(line)
		}
	}
}

// TestDay24Part1Minute12 verifies blizzards return to original positions
func TestDay24Part1Minute12(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	v := NewValley(lines)

	// At minute 12, blizzards should be back to original positions
	expected := []string{
		"#.######",
		"#>>.<^<#",
		"#.<..<<#",
		"#>v.><>#",
		"#<^v^^>#",
		"######.#",
	}

	actual := v.RenderGrid(12)

	// Compare
	for y := 0; y < len(expected); y++ {
		if actual[y] != expected[y] {
			t.Errorf("Minute 12, row %d:\n  want: %s\n  got:  %s", y, expected[y], actual[y])
		}
	}

	if t.Failed() {
		t.Log("Full grid at minute 12:")
		for _, line := range actual {
			t.Log(line)
		}
	}
}

// TestDay24Part1Minute11 verifies the blizzard configuration at minute 11
func TestDay24Part1Minute11(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	v := NewValley(lines)

	// Expected grid at Minute 11 (without E marker)
	// The rundown shows #2^E^2># where E hides what's underneath
	// Our simulation determines what's actually there
	expected := []string{
		"#.######",
		"#2^.^2>#",  // Position 3: our simulation shows no blizzard
		"#<v<.^<#",
		"#..2.>2#",
		"#.<..>.#",
		"######.#",
	}

	actual := v.RenderGrid(11)

	// Compare
	for y := 0; y < len(expected); y++ {
		if actual[y] != expected[y] {
			t.Errorf("Minute 11, row %d:\n  want: %s\n  got:  %s", y, expected[y], actual[y])
		}
	}

	if t.Failed() {
		t.Log("Full grid at minute 11:")
		for _, line := range actual {
			t.Log(line)
		}
	}
}
