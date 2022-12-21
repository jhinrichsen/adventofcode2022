package adventofcode2022

func Day17(line string, rocks int) int {
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
				panic("jet pattern must be < or >")
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

		cycle := len(shapes) * len(line)
	}
	return int(tower.Height())
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
