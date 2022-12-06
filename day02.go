package adventofcode2021

const (
	rock     = 1
	paper    = 2
	scissors = 3

	draw = 3
	lose = 0
	win  = 6
)

var scores = map[string]int{
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

func Day02(lines []string) int {
	sum := 0
	for _, line := range lines {
		score := scores[line]
		sum += score
	}
	return sum
}
