package adventofcode2022

const (
	North = 0 + 1i
	South = 0 - 1i
	East  = 1 + 0i
	West  = -1 + 0i

	NorthEast = North + East
	NorthWest = North + West
	SouthEast = South + East
	SouthWest = South + West
)

func R2c(x, y int) complex128 {
	return complex(float64(x), float64(y))
}
