package adventofcode2022

import (
	"strconv"
	"strings"
)

// Day16 solves day 16: Proboscidea Volcanium
func Day16(lines []string, part1 bool) uint {
	valves := parseValves(lines)

	if part1 {
		return day16Part1(valves)
	}
	return 0 // Part 2 placeholder
}

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
}

func parseValves(lines []string) map[string]*Valve {
	valves := make(map[string]*Valve)

	for _, line := range lines {
		// Parse: "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB"
		parts := strings.Split(line, "; ")
		if len(parts) != 2 {
			continue
		}

		// Parse valve name and flow rate
		leftParts := strings.Split(parts[0], " has flow rate=")
		if len(leftParts) != 2 {
			continue
		}

		name := strings.TrimPrefix(leftParts[0], "Valve ")
		flowRate, err := strconv.Atoi(leftParts[1])
		if err != nil {
			continue
		}

		// Parse tunnels
		tunnelsPart := parts[1]
		tunnelsPart = strings.TrimPrefix(tunnelsPart, "tunnels lead to valves ")
		tunnelsPart = strings.TrimPrefix(tunnelsPart, "tunnel leads to valve ")
		tunnels := strings.Split(tunnelsPart, ", ")

		valves[name] = &Valve{
			name:     name,
			flowRate: flowRate,
			tunnels:  tunnels,
		}
	}

	return valves
}

type State struct {
	pos     string
	time    int
	opened  uint64 // Bitmask of opened valves
	pressure uint
}

func day16Part1(valves map[string]*Valve) uint {
	// Build valve index for bitmask
	valveIdx := make(map[string]int)
	idx := 0
	for name := range valves {
		if valves[name].flowRate > 0 {
			valveIdx[name] = idx
			idx++
		}
	}

	maxPressure := uint(0)

	var dfs func(pos string, time int, opened uint64, totalPressure uint)
	dfs = func(pos string, time int, opened uint64, totalPressure uint) {
		if time > 30 {
			return
		}

		// Update max pressure
		if totalPressure > maxPressure {
			maxPressure = totalPressure
		}

		// Try opening current valve if it has flow and isn't opened
		if vIdx, exists := valveIdx[pos]; exists {
			bit := uint64(1) << vIdx
			if opened&bit == 0 && time < 30 { // Not yet opened and time left
				newOpened := opened | bit
				flowRate := uint(valves[pos].flowRate)
				// Valve opens at time+1, releases pressure for remaining time
				remainingTime := 30 - time
				addedPressure := flowRate * uint(remainingTime)
				dfs(pos, time+1, newOpened, totalPressure+addedPressure)
			}
		}

		// Try moving to adjacent valves
		if time < 30 {
			for _, next := range valves[pos].tunnels {
				dfs(next, time+1, opened, totalPressure)
			}
		}
	}

	dfs("AA", 0, 0, 0)
	return maxPressure
}
