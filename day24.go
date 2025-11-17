package adventofcode2022

func NewDay24(lines []string) day24 {
	var (
		world day24
		X     = len(lines[0])
		Y     = len(lines)
	)
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			// invert y when converting from lines
			// to our complex map
			c := R2c(x, len(lines)-1-y)

			var facing complex128
			switch lines[y][x] {
			case '#': // ignore
			case '.':
				// check for entrance or exit
				if y == 0 {
					world.Start = c
				} else if y == len(lines)-1 {
					world.Finish = c
				}
			case '^':
				facing = North
			case '>':
				facing = East
			case 'v':
				facing = South
			case '<':
				facing = West
			default:
				facing = 0 // Unknown element, no facing
			}
			world.blizzards = append(world.blizzards,
				Blizzard{Dim: R2c(X, Y), Facing: facing})
		}
	}
	return world
}

type day24 struct {
	// m             map[complex128]int
	Dim           complex128
	Start, Finish complex128
	blizzards     []Blizzard
}

type Blizzard struct {
	Dim, Facing, Position complex128
	// x, y                  float64
}

func (a *Blizzard) Step() {
	a.Position += a.Facing

	// wrap
	switch a.Facing {
	case North:
		if imag(a.Position) == imag(a.Dim) {
			a.Position = complex(real(a.Position), 1)
		}
	case East:
	case South:
	case West:
	default:
		// Unknown direction, no step
	}
}
