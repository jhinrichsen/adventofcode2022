package adventofcode2022

const (
	rock     = 1
	paper    = 2
	scissors = 3

	lose = 0
	draw = 3
	win  = 6
)

var scoresPart1 = map[string]int{
	// opponent draws rock
	"A X": rock + draw,
	"A Y": paper + win,
	"A Z": scissors + lose,

	// opponent draws paper
	"B X": rock + lose,
	"B Y": paper + draw,
	"B Z": scissors + win,

	// opponent draws scissors
	"C X": rock + win,
	"C Y": paper + lose,
	"C Z": scissors + draw,
}

var scoresPart2 = map[string]int{
	// opponent draws rock
	"A X": lose + scissors,
	"A Y": draw + rock,
	"A Z": win + paper,

	// opponent draws paper
	"B X": lose + rock,
	"B Y": draw + paper,
	"B Z": win + scissors,

	// opponent draws scissors
	"C X": lose + paper,
	"C Y": draw + scissors,
	"C Z": win + rock,
}

func Day02(lines []string, part1 bool) int {
	scores := scoresPart2
	if part1 {
		scores = scoresPart1
	}
	sum := 0
	for _, line := range lines {
		score := scores[line]
		sum += score
	}
	return sum
}
