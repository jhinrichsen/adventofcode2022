package adventofcode2022

func Day22(lines []string, part1 bool) uint {
	if part1 {
		return day22Part1(lines)
	}
	return day22Part2(lines)
}

func day22Part1(lines []string) uint {
	const (
		north = 0 + 1i
		south = 0 - 1i
		east  = 1 + 0i
		west  = -1 + 0i
	)
	var (
		m = make(map[complex128]bool) // 1-based

		walk = func(c complex128) bool {
			// entry and it is true
			b, ok := m[c]
			return ok && b
		}
		wall = func(c complex128) bool {
			// entry and it is false
			b, ok := m[c]
			return ok && !b
		}
		wrap = func(c complex128) bool {
			// no entry
			_, ok := m[c]
			return !ok
		}

		parseMap = func() complex128 {
			var maxX, maxY int
			for y := 0; y < len(lines); y++ {
				if lines[y] == "" {
					break
				}

				// lines are not right padded
				thisX := len(lines[y])
				if thisX > maxX {
					maxX = thisX
				}
				maxY = y

				for x := 0; x < thisX; x++ {
					c := complex(float64(x+1), float64(y+1))
					switch lines[y][x] {
					case '.': // exists and can walk
						m[c] = true
					case '#': // exists and cannot walk
						m[c] = false
					default:
					}
				}
			}
			return complex(float64(maxX+1), float64(maxY+1))
		}

		facing = east
	)

	parseMap()
	position := func() complex128 {
		c := 1 + 1i
		for {
			if walk(c) {
				break
			}
			if wall(c) {
				break
			}
			c += east
		}
		return c
	}()

	for _, cmd := range NewCommands(lines[len(lines)-1]) {
		// move
		for j := 0; j < cmd.count; {
			p2 := position + facing
			if wrap(p2) {
				// we stepped on a wrap field, carriage return into
				// opposite direction until wrap (or border, which is a
				// wrap)
				opposite := facing * -1
				for walk(p2+opposite) || wall(p2+opposite) {
					p2 += opposite
				}
			}
			if wall(p2) {
				break
			}
			if walk(p2) {
				position = p2
				j++
			}
		}
		facing *= cmd.turn
	}
	result := 1000 * uint(imag(position))
	result += 4 * uint(real(position))
	switch facing {
	case east:
		// NOP
	case south:
		result += 1
	case west:
		result += 2
	case north:
		result += 3
	}
	return result
}

type face struct {
	x, y int // top-left corner in face coordinates (0-indexed)
}

func day22Part2(lines []string) uint {
	// Parse map
	var mapLines []string
	for _, line := range lines {
		if line == "" {
			break
		}
		mapLines = append(mapLines, line)
	}

	// Detect face size
	var maxLen int
	for _, line := range mapLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	faceSize := 50
	if maxLen <= 16 {
		faceSize = 4
	}

	// Build grid and face map
	grid := make([][]byte, len(mapLines))
	for i := range grid {
		grid[i] = make([]byte, maxLen)
		for j := range grid[i] {
			if j < len(mapLines[i]) {
				grid[i][j] = mapLines[i][j]
			} else {
				grid[i][j] = ' '
			}
		}
	}

	// Define cube topology based on layout
	// Returns (newX, newY, newFacing) given (x, y, facing) attempting to leave a face
	var wrapCube func(x, y, facing int) (int, int, int)

	if faceSize == 4 {
		// Example layout (.=space, X=face):
		//   .XXX
		//   .XX.
		wrapCube = wrapCubeExample
	} else {
		// Actual input layout:
		//   .XX
		//   .X.
		//   XX.
		//   X..
		wrapCube = wrapCubeActual
	}

	// Find starting position
	x, y := 0, 0
	for x = 0; x < len(grid[0]); x++ {
		if grid[0][x] == '.' {
			break
		}
	}

	facing := 0 // 0=right, 1=down, 2=left, 3=up

	// Process commands
	for _, cmd := range NewCommands(lines[len(lines)-1]) {
		// Move
		for step := 0; step < cmd.count; step++ {
			nx, ny, nf := x, y, facing
			switch facing {
			case 0: // right
				nx++
			case 1: // down
				ny++
			case 2: // left
				nx--
			case 3: // up
				ny--
			}

			// Check if we need to wrap
			if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) || grid[ny][nx] == ' ' {
				nx, ny, nf = wrapCube(x, y, facing)
			}

			// Check for wall
			if grid[ny][nx] == '#' {
				break
			}

			x, y, facing = nx, ny, nf
		}

		// Turn
		switch cmd.turn {
		case 0 - 1i: // L
			facing = (facing + 3) % 4
		case 0 + 1i: // R
			facing = (facing + 1) % 4
		}
	}

	return uint(1000*(y+1) + 4*(x+1) + facing)
}

func wrapCubeExample(x, y, facing int) (int, int, int) {
	// Example cube net (face size 4):
	//     [1]
	// [2][3][4]
	//     [5][6]

	faceSize := 4
	fx, fy := x/faceSize, y/faceSize  // which face (in face coords)
	lx, ly := x%faceSize, y%faceSize  // local position within face

	// Determine current face
	var faceID int
	if fy == 0 && fx == 2 {
		faceID = 1
	} else if fy == 1 && fx == 0 {
		faceID = 2
	} else if fy == 1 && fx == 1 {
		faceID = 3
	} else if fy == 1 && fx == 2 {
		faceID = 4
	} else if fy == 2 && fx == 2 {
		faceID = 5
	} else if fy == 2 && fx == 3 {
		faceID = 6
	}

	// Define wrapping rules
	var targetFaceX, targetFaceY, newLX, newLY, newFacing int

	switch faceID {
	case 1:
		switch facing {
		case 0: // right -> face 6 left edge, flipped
			targetFaceX, targetFaceY = 3, 2
			newLX, newLY = faceSize-1, faceSize-1-ly
			newFacing = 2
		case 1: // down -> face 4 top edge
			targetFaceX, targetFaceY = 2, 1
			newLX, newLY = lx, 0
			newFacing = 1
		case 2: // left -> face 3 top edge
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = ly, 0
			newFacing = 1
		case 3: // up -> face 2 top edge, flipped
			targetFaceX, targetFaceY = 0, 1
			newLX, newLY = faceSize-1-lx, 0
			newFacing = 1
		}
	case 2:
		switch facing {
		case 0: // right -> face 3 left edge
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = 0, ly
			newFacing = 0
		case 1: // down -> face 5 bottom edge, flipped
			targetFaceX, targetFaceY = 2, 2
			newLX, newLY = faceSize-1-lx, faceSize-1
			newFacing = 3
		case 2: // left -> face 6 bottom edge, flipped
			targetFaceX, targetFaceY = 3, 2
			newLX, newLY = faceSize-1-ly, faceSize-1
			newFacing = 3
		case 3: // up -> face 1 top edge, flipped
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = faceSize-1-lx, 0
			newFacing = 1
		}
	case 3:
		switch facing {
		case 0: // right -> face 4 left edge
			targetFaceX, targetFaceY = 2, 1
			newLX, newLY = 0, ly
			newFacing = 0
		case 1: // down -> face 5 left edge, flipped
			targetFaceX, targetFaceY = 2, 2
			newLX, newLY = 0, faceSize-1-lx
			newFacing = 0
		case 2: // left -> face 2 right edge
			targetFaceX, targetFaceY = 0, 1
			newLX, newLY = faceSize-1, ly
			newFacing = 2
		case 3: // up -> face 1 left edge
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = lx, 0
			newFacing = 1
		}
	case 4:
		switch facing {
		case 0: // right -> face 6 top edge, flipped
			targetFaceX, targetFaceY = 3, 2
			newLX, newLY = faceSize-1-ly, 0
			newFacing = 1
		case 1: // down -> face 5 top edge
			targetFaceX, targetFaceY = 2, 2
			newLX, newLY = lx, 0
			newFacing = 1
		case 2: // left -> face 3 right edge
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = faceSize-1, ly
			newFacing = 2
		case 3: // up -> face 1 bottom edge
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = lx, faceSize-1
			newFacing = 3
		}
	case 5:
		switch facing {
		case 0: // right -> face 6 left edge
			targetFaceX, targetFaceY = 3, 2
			newLX, newLY = 0, ly
			newFacing = 0
		case 1: // down -> face 2 bottom edge, flipped
			targetFaceX, targetFaceY = 0, 1
			newLX, newLY = faceSize-1-lx, faceSize-1
			newFacing = 3
		case 2: // left -> face 3 bottom edge, flipped
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = faceSize-1-ly, faceSize-1
			newFacing = 3
		case 3: // up -> face 4 bottom edge
			targetFaceX, targetFaceY = 2, 1
			newLX, newLY = lx, faceSize-1
			newFacing = 3
		}
	case 6:
		switch facing {
		case 0: // right -> face 1 right edge, flipped
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = faceSize-1, faceSize-1-ly
			newFacing = 2
		case 1: // down -> face 2 left edge, flipped
			targetFaceX, targetFaceY = 0, 1
			newLX, newLY = 0, faceSize-1-lx
			newFacing = 0
		case 2: // left -> face 5 right edge
			targetFaceX, targetFaceY = 2, 2
			newLX, newLY = faceSize-1, ly
			newFacing = 2
		case 3: // up -> face 4 right edge, flipped
			targetFaceX, targetFaceY = 2, 1
			newLX, newLY = faceSize-1, faceSize-1-lx
			newFacing = 2
		}
	}

	return targetFaceX*faceSize + newLX, targetFaceY*faceSize + newLY, newFacing
}

func wrapCubeActual(x, y, facing int) (int, int, int) {
	// Actual cube net (face size 50):
	//     [1][2]
	//     [3]
	// [4][5]
	// [6]

	faceSize := 50
	fx, fy := x/faceSize, y/faceSize
	lx, ly := x%faceSize, y%faceSize

	// Determine current face
	var faceID int
	if fy == 0 && fx == 1 {
		faceID = 1
	} else if fy == 0 && fx == 2 {
		faceID = 2
	} else if fy == 1 && fx == 1 {
		faceID = 3
	} else if fy == 2 && fx == 0 {
		faceID = 4
	} else if fy == 2 && fx == 1 {
		faceID = 5
	} else if fy == 3 && fx == 0 {
		faceID = 6
	}

	var targetFaceX, targetFaceY, newLX, newLY, newFacing int

	switch faceID {
	case 1:
		switch facing {
		case 0: // right -> face 2
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = 0, ly
			newFacing = 0
		case 1: // down -> face 3
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = lx, 0
			newFacing = 1
		case 2: // left -> face 4 left edge, flipped
			targetFaceX, targetFaceY = 0, 2
			newLX, newLY = 0, faceSize-1-ly
			newFacing = 0
		case 3: // up -> face 6 left edge
			targetFaceX, targetFaceY = 0, 3
			newLX, newLY = 0, lx
			newFacing = 0
		}
	case 2:
		switch facing {
		case 0: // right -> face 5 right edge, flipped
			targetFaceX, targetFaceY = 1, 2
			newLX, newLY = faceSize-1, faceSize-1-ly
			newFacing = 2
		case 1: // down -> face 3 right edge
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = faceSize-1, lx
			newFacing = 2
		case 2: // left -> face 1
			targetFaceX, targetFaceY = 1, 0
			newLX, newLY = faceSize-1, ly
			newFacing = 2
		case 3: // up -> face 6 bottom edge
			targetFaceX, targetFaceY = 0, 3
			newLX, newLY = lx, faceSize-1
			newFacing = 3
		}
	case 3:
		switch facing {
		case 0: // right -> face 2 bottom edge
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = ly, faceSize-1
			newFacing = 3
		case 1: // down -> face 5
			targetFaceX, targetFaceY = 1, 2
			newLX, newLY = lx, 0
			newFacing = 1
		case 2: // left -> face 4 top edge
			targetFaceX, targetFaceY = 0, 2
			newLX, newLY = ly, 0
			newFacing = 1
		case 3: // up -> face 1
			targetFaceX, targetFaceY = 1, 0
			newLX, newLY = lx, faceSize-1
			newFacing = 3
		}
	case 4:
		switch facing {
		case 0: // right -> face 5
			targetFaceX, targetFaceY = 1, 2
			newLX, newLY = 0, ly
			newFacing = 0
		case 1: // down -> face 6
			targetFaceX, targetFaceY = 0, 3
			newLX, newLY = lx, 0
			newFacing = 1
		case 2: // left -> face 1 left edge, flipped
			targetFaceX, targetFaceY = 1, 0
			newLX, newLY = 0, faceSize-1-ly
			newFacing = 0
		case 3: // up -> face 3 left edge
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = 0, lx
			newFacing = 0
		}
	case 5:
		switch facing {
		case 0: // right -> face 2 right edge, flipped
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = faceSize-1, faceSize-1-ly
			newFacing = 2
		case 1: // down -> face 6 right edge
			targetFaceX, targetFaceY = 0, 3
			newLX, newLY = faceSize-1, lx
			newFacing = 2
		case 2: // left -> face 4
			targetFaceX, targetFaceY = 0, 2
			newLX, newLY = faceSize-1, ly
			newFacing = 2
		case 3: // up -> face 3
			targetFaceX, targetFaceY = 1, 1
			newLX, newLY = lx, faceSize-1
			newFacing = 3
		}
	case 6:
		switch facing {
		case 0: // right -> face 5 bottom edge
			targetFaceX, targetFaceY = 1, 2
			newLX, newLY = ly, faceSize-1
			newFacing = 3
		case 1: // down -> face 2 top edge
			targetFaceX, targetFaceY = 2, 0
			newLX, newLY = lx, 0
			newFacing = 1
		case 2: // left -> face 1 top edge
			targetFaceX, targetFaceY = 1, 0
			newLX, newLY = ly, 0
			newFacing = 1
		case 3: // up -> face 4
			targetFaceX, targetFaceY = 0, 2
			newLX, newLY = lx, faceSize-1
			newFacing = 3
		}
	}

	return targetFaceX*faceSize + newLX, targetFaceY*faceSize + newLY, newFacing
}

type Command struct {
	count int
	turn  complex128
}

func NewCommands(s string) (cs []Command) {
	var (
		n      int
		noturn complex128
	)
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n *= 10
			n += int(c - '0')
		} else {
			var turn complex128
			// y extends south, flip
			switch c {
			case 'L':
				turn = 0 - 1i
			case 'R':
				turn = 0 + 1i
			default:
				turn = noturn // Unknown turn, use no turn
			}
			cs = append(cs, Command{n, turn})
			n = 0
		}
	}
	// railing nomber
	cs = append(cs, Command{n, noturn})
	return
}
