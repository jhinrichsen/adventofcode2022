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

		// Shapes as slices for faster iteration (no map allocations)
		shapes = [][]complex128{
			{ // ####
				0 + 0i,
				1 + 0i,
				2 + 0i,
				3 + 0i,
			},
			{ // .#. / ### / .#.
				1 + 2i,
				0 + 1i, 1 + 1i, 2 + 1i,
				1 + 0i,
			},
			{ // ..# / ..# / ###
				2 + 2i,
				2 + 1i,
				0 + 0i, 1 + 0i, 2 + 0i,
			},
			{ // # / # / # / #
				0 + 3i,
				0 + 2i,
				0 + 1i,
				0 + 0i,
			},
			{ // ## / ##
				0 + 1i, 1 + 1i,
				0 + 0i, 1 + 0i,
			},
		}
		tower = NewSprite()

		// Check if position is valid (in field and no collision)
		isValid = func(shape []complex128, position complex128) bool {
			for i := range shape {
				c := shape[i] + position
				if real(c) < 0 || real(c) >= width || imag(c) < 0 {
					return false
				}
				if tower.rocks[c] {
					return false
				}
			}
			return true
		}
	)

	// Track height incrementally to avoid recalculation
	var currentHeight uint

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

	// Helper to get surface profile (top 50 rows for better cycle detection)
	// Pre-allocate to avoid repeated allocations
	profileBuf := make([]byte, 0, 50*(width+1))
	getProfile := func() string {
		h := int(currentHeight)
		if h == 0 {
			return ""
		}
		profileBuf = profileBuf[:0]
		depth := min(50, h)
		for y := h - 1; y >= h-depth && y >= 0; y-- {
			for x := 0; x < width; x++ {
				if tower.rocks[complex(float64(x), float64(y))] {
					profileBuf = append(profileBuf, '#')
				} else {
					profileBuf = append(profileBuf, '.')
				}
			}
			profileBuf = append(profileBuf, '|')
		}
		return string(profileBuf)
	}

	// Track cycle detection state
	var cycleDetected bool
	var cycleEndHeight, cycleHeightGain uint
	var targetRockCount int

	for i := 0; i < rocks; i++ {
		shape := shapes[i%len(shapes)] // forever next shape

		position := complex(0, float64(currentHeight))
		position += offset

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

			if isValid(shape, position+step) {
				position += step
			}

			// vertical move
			if isValid(shape, position+south) {
				position += south
			} else {
				// freeze - add rock positions to tower and update height
				var maxHeight uint
				for j := range shape {
					c := shape[j] + position
					tower.rocks[c] = true
					h := uint(imag(c)) + 1
					if h > maxHeight {
						maxHeight = h
					}
				}
				if maxHeight > currentHeight {
					currentHeight = maxHeight
				}
				break
			}
		}

		// Check if we've reached the target rock count after cycle detection
		if cycleDetected && i == targetRockCount {
			remainderHeight := currentHeight - cycleEndHeight
			finalHeight := cycleEndHeight + cycleHeightGain + remainderHeight
			return finalHeight
		}

		// Cycle detection (only useful for large rock counts)
		if !cycleDetected && rocks > 5000 && i > 1000 && i%10 == 0 {
			state := State{
				rockIdx: i % len(shapes),
				jetIdx:  jetPattern.idx,
				profile: getProfile(),
			}

			if prev, found := seen[state]; found {
				// Found a cycle!
				cycleLength := i - prev.rockCount
				cycleHeight := currentHeight - prev.height

				// Calculate how many complete cycles we can skip
				remainingRocks := rocks - i
				fullCycles := remainingRocks / cycleLength
				remainder := remainingRocks % cycleLength

				// Set up to continue simulation for remainder rocks only
				cycleDetected = true
				cycleEndHeight = currentHeight
				cycleHeightGain = uint(fullCycles) * cycleHeight
				targetRockCount = i + remainder - 1 // -1 because we check after placing

				// If remainder is 0, we can return immediately
				if remainder == 0 {
					return cycleEndHeight + cycleHeightGain
				}
			} else {
				seen[state] = struct {
					rockCount int
					height    uint
				}{i, currentHeight}
			}
		}
	}
	return currentHeight
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
