package adventofcode2022

import (
	"strconv"
	"strings"
)

type Monkey struct {
	Items       []uint
	Operation   func(uint) uint
	DivisibleBy uint
	IfTrue      int
	IfFalse     int
	Inspections uint
}

func parseMonkey(lines []string) Monkey {
	var m Monkey

	// Parse starting items
	itemsStr := strings.Split(lines[1], ": ")[1]
	for _, s := range strings.Split(itemsStr, ", ") {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		m.Items = append(m.Items, uint(n))
	}

	// Parse operation
	opLine := strings.TrimSpace(lines[2][len("  Operation: new = "):])
	parts := strings.Fields(opLine)
	if len(parts) == 3 {
		op := parts[1]
		operand := parts[2]

		if operand == "old" {
			// old op old
			if op == "*" {
				m.Operation = func(old uint) uint { return old * old }
			} else if op == "+" {
				m.Operation = func(old uint) uint { return old + old }
			}
		} else {
			// old op number
			val, err := strconv.Atoi(operand)
			if err == nil {
				if op == "*" {
					m.Operation = func(old uint) uint { return old * uint(val) }
				} else if op == "+" {
					m.Operation = func(old uint) uint { return old + uint(val) }
				}
			}
		}
	}

	// Parse divisible by
	divisibleLine := strings.Fields(lines[3])
	if len(divisibleLine) >= 4 {
		n, err := strconv.Atoi(divisibleLine[3])
		if err == nil {
			m.DivisibleBy = uint(n)
		}
	}

	// Parse if true
	trueLine := strings.Fields(lines[4])
	if len(trueLine) >= 6 {
		n, err := strconv.Atoi(trueLine[5])
		if err == nil {
			m.IfTrue = n
		}
	}

	// Parse if false
	falseLine := strings.Fields(lines[5])
	if len(falseLine) >= 6 {
		n, err := strconv.Atoi(falseLine[5])
		if err == nil {
			m.IfFalse = n
		}
	}

	return m
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey
	for i := 0; i < len(lines); i += 7 {
		end := i + 6
		if end >= len(lines) {
			end = len(lines) - 1
		}
		monkeys = append(monkeys, parseMonkey(lines[i:end+1]))
	}
	return monkeys
}

func Day11(lines []string, part1 bool) uint {
	monkeys := parseMonkeys(lines)

	// Compute LCM (product of all divisors, assuming they're coprime)
	lcm := uint(1)
	for i := range monkeys {
		lcm *= monkeys[i].DivisibleBy
	}

	rounds := 20
	if !part1 {
		rounds = 10000
	}

	for range rounds {
		for i := range monkeys {
			for len(monkeys[i].Items) > 0 {
				monkeys[i].Inspections++

				// Get item and remove from current monkey
				worry := monkeys[i].Items[0]
				monkeys[i].Items = monkeys[i].Items[1:]

				// Apply operation
				worry = monkeys[i].Operation(worry)

				// Apply relief (Part 1 only)
				if part1 {
					worry /= 3
				} else {
					// Keep worry manageable using modular arithmetic
					worry %= lcm
				}

				// Test and throw
				var target int
				if worry%monkeys[i].DivisibleBy == 0 {
					target = monkeys[i].IfTrue
				} else {
					target = monkeys[i].IfFalse
				}
				monkeys[target].Items = append(monkeys[target].Items, worry)
			}
		}
	}

	// Find two most active monkeys
	max1, max2 := uint(0), uint(0)
	for i := range monkeys {
		if monkeys[i].Inspections > max1 {
			max2 = max1
			max1 = monkeys[i].Inspections
		} else if monkeys[i].Inspections > max2 {
			max2 = monkeys[i].Inspections
		}
	}

	return max1 * max2
}

// Day11Inspections runs the simulation and returns inspection counts after specified rounds
func Day11Inspections(lines []string, rounds int) []uint {
	monkeys := parseMonkeys(lines)

	// Compute LCM
	lcm := uint(1)
	for i := range monkeys {
		lcm *= monkeys[i].DivisibleBy
	}

	for range rounds {
		for i := range monkeys {
			for len(monkeys[i].Items) > 0 {
				monkeys[i].Inspections++
				worry := monkeys[i].Items[0]
				monkeys[i].Items = monkeys[i].Items[1:]
				worry = monkeys[i].Operation(worry)
				worry %= lcm // No division by 3 in Part 2

				var target int
				if worry%monkeys[i].DivisibleBy == 0 {
					target = monkeys[i].IfTrue
				} else {
					target = monkeys[i].IfFalse
				}
				monkeys[target].Items = append(monkeys[target].Items, worry)
			}
		}
	}

	inspections := make([]uint, len(monkeys))
	for i := range monkeys {
		inspections[i] = monkeys[i].Inspections
	}
	return inspections
}
