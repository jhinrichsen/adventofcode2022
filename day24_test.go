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
	const want uint = 300
	lines, err := linesFromFilename(filename(24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// TestDay24IntermediateStates verifies each step of the example walkthrough
func TestDay24IntermediateStates(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	v := NewValley(lines)

	// Track the expedition's path as shown in the rundown
	// Format: minute -> (x, y) position
	path := map[int]complex128{
		0:  complex(1, 0),  // Initial: start position (outside entrance)
		1:  complex(1, 1),  // Minute 1: move down to (1,1)
		2:  complex(1, 2),  // Minute 2: move down to (1,2)
		3:  complex(1, 2),  // Minute 3: wait at (1,2)
		4:  complex(1, 1),  // Minute 4: move up to (1,1)
		5:  complex(2, 1),  // Minute 5: move right to (2,1)
		6:  complex(3, 1),  // Minute 6: move right to (3,1)
		7:  complex(3, 2),  // Minute 7: move down to (3,2)
		8:  complex(2, 2),  // Minute 8: move left to (2,2)
		9:  complex(2, 1),  // Minute 9: move up to (2,1)
		10: complex(3, 1),  // Minute 10: move right to (3,1)
		11: complex(3, 1),  // Minute 11: wait at (3,1)
		12: complex(3, 2),  // Minute 12: move down to (3,2)
		13: complex(3, 3),  // Minute 13: move down to (3,3)
		14: complex(4, 3),  // Minute 14: move right to (4,3)
		15: complex(5, 3),  // Minute 15: move right to (5,3)
		16: complex(6, 3),  // Minute 16: move right to (6,3)
		17: complex(6, 4),  // Minute 17: move down to (6,4)
		18: complex(6, 5),  // Minute 18: move down to exit (6,5)
	}

	// Verify each position is valid at its corresponding time
	for minute, pos := range path {
		if !v.isValid(pos, minute) {
			t.Errorf("Minute %d: position (%v) should be valid but isn't", minute, pos)
		}
	}
}
