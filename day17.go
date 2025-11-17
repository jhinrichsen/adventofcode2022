package adventofcode2022

func Day17(line string, rocks int) uint {
	const (
		offset = 2 + 3i
		width  = 7

		east  = 1 + 0i
		west  = -1 + 0i
		south = 0 - 1i
	)
	var (
		jetPattern = JetPattern{line, 0}

		// hungarian notation done right: shape is relative, sprite is
		// absolute
		shapes = []Sprite{
			{
				// ####
				map[complex128]bool{
					0 + 0i: true,
					1 + 0i: true,
					2 + 0i: true,
					3 + 0i: true,
				},
				// 4 + 1i,
			},
			{
				// .#.
				// ###
				// .#.
				map[complex128]bool{
					1 + 2i: true,
					0 + 1i: true, 1 + 1i: true, 2 + 1i: true,
					1 + 0i: true,
				},
				// 3 + 3i,
			},
			{
				// ..#
				// ..#
				// ###
				map[complex128]bool{
					2 + 2i: true,
					2 + 1i: true,
					0 + 0i: true, 1 + 0i: true, 2 + 0i: true,
				},
				// 3 + 3i,
			},
			{
				// #
				// #
				// #
				// #
				map[complex128]bool{
					0 + 3i: true,
					0 + 2i: true,
					0 + 1i: true,
					0 + 0i: true,
				},
				// 1 + 4i,
			},
			{
				// ##
				// ##
				map[complex128]bool{
					0 + 1i: true, 1 + 1i: true,
					0 + 0i: true, 1 + 0i: true,
				},
				// 2 + 2i,
			},
		}
		tower = NewSprite()

		infield = func(a Sprite) bool {
			for c := range a.rocks {
				if real(c) < 0 || // off left
					real(c) >= width || // off right
					imag(c) < 0 { // off bottom, no off top
					return false
				}
			}
			return true
		}
	)

	// Cycle detection for part 2
	type State struct {
		rockIdx int
		jetIdx  int
		profile string
	}
	seen := make(map[State]struct {
		rockCount int
		height    uint
	})

	// Helper to get surface profile (top 30 rows for better cycle detection)
	getProfile := func() string {
		h := int(tower.Height())
		if h == 0 {
			return ""
		}
		profile := ""
		depth := min(30, h)
		for y := h - 1; y >= h-depth && y >= 0; y-- {
			for x := 0; x < width; x++ {
				if tower.rocks[complex(float64(x), float64(y))] {
					profile += "#"
				} else {
					profile += "."
				}
			}
			profile += "|"
		}
		return profile
	}

	for i := 0; i < rocks; i++ {
		shape := shapes[i%len(shapes)] // forever next shape

		position := complex(0, tower.Height())
		position += offset

		test := func(position, step complex128) (complex128, bool) {
			sprite := shape.Translate(position + step)
			// must check against both border and tower
			p1 := sprite.Collides(tower)
			p2 := infield(sprite)
			possible := !p1 && p2
			if !possible {
				return position, false
			}
			return position + step, true
		}

		for {
			// horizontal move

			var step complex128
			switch jetPattern.Next() {
			case '>':
				step = east
			case '<':
				step = west
			default:
				step = east // Default to east if invalid
			}

			var ok bool
			position, _ = test(position, step)

			// vertical move

			position, ok = test(position, south)
			if !ok { // freeze
				sprite := shape.Translate(position)
				tower.AddSprite(sprite)
				break
			}
		}

		// Cycle detection (only useful for large rock counts)
		// Only check every 10 rocks to reduce overhead
		if rocks > 5000 && i > 1000 && i%10 == 0 {
			state := State{
				rockIdx: i % len(shapes),
				jetIdx:  jetPattern.idx,
				profile: getProfile(),
			}

			if prev, found := seen[state]; found {
				// Found a cycle!
				cycleLength := i - prev.rockCount
				cycleHeight := uint(tower.Height()) - prev.height

				// Calculate how many complete cycles we can skip
				remainingRocks := rocks - (i + 1)
				fullCycles := remainingRocks / cycleLength
				rocksPastCycle := remainingRocks % cycleLength

				// Calculate final height
				currentHeight := uint(tower.Height())
				cycleContribution := uint(fullCycles) * cycleHeight

				// For remainder, continue simulating from current position
				for j := 0; j < rocksPastCycle; j++ {
					idx := i + 1 + j
					remShape := shapes[idx%len(shapes)]

					position := complex(0, tower.Height())
					position += offset

					// Define local test function for remainder simulation
					remTest := func(pos, step complex128) (complex128, bool) {
						sprite := remShape.Translate(pos + step)
						p1 := sprite.Collides(tower)
						p2 := infield(sprite)
						possible := !p1 && p2
						if !possible {
							return pos, false
						}
						return pos + step, true
					}

					for {
						// horizontal move
						var step complex128
						switch jetPattern.Next() {
						case '>':
							step = east
						case '<':
							step = west
						default:
							step = east
						}

						position, _ = remTest(position, step)

						// vertical move
						position, ok := remTest(position, south)
						if !ok {
							sprite := remShape.Translate(position)
							tower.AddSprite(sprite)
							break
						}
					}
					if j%10 == 0 {
						println("Remainder progress:", j, "/", rocksPastCycle)
					}
				}

				println("Remainder simulation complete")
				remainderHeight := uint(tower.Height()) - currentHeight
				result := currentHeight + cycleContribution + remainderHeight
				println("Final height:", result)
				return result
			}

			seen[state] = struct {
				rockCount int
				height    uint
			}{i, uint(tower.Height())}
		}
	}
	return uint(tower.Height())
}

type Sprite struct {
	rocks map[complex128]bool
}

func NewSprite() Sprite {
	return Sprite{make(map[complex128]bool)}
}

func SpriteFrom(cs ...complex128) Sprite {
	s := NewSprite()
	s.Add(cs...)
	return s
}

func (a Sprite) Collides(b Sprite) bool {
	// PERF use smaller one for comparison
	for c := range a.rocks {
		if _, ok := b.rocks[c]; ok {
			return true
		}
	}
	return false
}

// Translate converts a relative shape into an absolute sprite.
func (a Sprite) Translate(position complex128) Sprite {
	b := NewSprite()
	for c := range a.rocks {
		b.rocks[c+position] = true
	}
	return b
}

// Add a new Sprite, and optionally extend hull
func (a *Sprite) Add(cs ...complex128) {
	for _, c := range cs {
		a.rocks[c] = true
	}
}

func (a *Sprite) AddSprite(b Sprite) {
	for c := range b.rocks {
		a.Add(c)
	}
}

func (a Sprite) Height() float64 {
	var max float64

	if len(a.rocks) == 0 {
		return 0
	}

	for c := range a.rocks {
		h := imag(c)
		if h > max {
			max = h
		}
	}
	return 1 + max
}

type JetPattern struct {
	line string
	idx  int
}

func (a *JetPattern) Next() byte {
	inc := func() {
		a.idx = (a.idx + 1) % len(a.line)
	}
	defer inc()
	return a.line[a.idx]
}
