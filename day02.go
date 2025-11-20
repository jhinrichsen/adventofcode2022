package adventofcode2022

// Day02 calculates Rock Paper Scissors tournament score.
// Part1: X/Y/Z are your moves (rock/paper/scissors)
// Part2: X/Y/Z are desired outcomes (lose/draw/win)
func Day02(lines []string, part1 bool) uint {
	var sum uint

	// Lookup tables indexed by [opponent][you] where:
	// opponent: A=0, B=1, C=2 (rock, paper, scissors)
	// you: X=0, Y=1, Z=2
	var scores [3][3]uint

	if part1 {
		// Part 1: X/Y/Z = rock/paper/scissors
		scores = [3][3]uint{
			// A (rock):     X(rock)  Y(paper) Z(scissors)
			{1 + 3, 2 + 6, 3 + 0}, // draw, win, lose
			// B (paper):    X(rock)  Y(paper) Z(scissors)
			{1 + 0, 2 + 3, 3 + 6}, // lose, draw, win
			// C (scissors): X(rock)  Y(paper) Z(scissors)
			{1 + 6, 2 + 0, 3 + 3}, // win, lose, draw
		}
	} else {
		// Part 2: X/Y/Z = lose/draw/win
		scores = [3][3]uint{
			// A (rock):     X(lose)  Y(draw) Z(win)
			{0 + 3, 3 + 1, 6 + 2}, // scissors, rock, paper
			// B (paper):    X(lose)  Y(draw) Z(win)
			{0 + 1, 3 + 2, 6 + 3}, // rock, paper, scissors
			// C (scissors): X(lose)  Y(draw) Z(win)
			{0 + 2, 3 + 3, 6 + 1}, // paper, scissors, rock
		}
	}

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		// Parse "A X" format: opponent at [0], you at [2]
		opponent := line[0] - 'A' // 0, 1, or 2
		you := line[2] - 'X'      // 0, 1, or 2
		sum += scores[opponent][you]
	}

	return sum
}
