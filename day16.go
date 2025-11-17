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

// BFS to find shortest distance between two valves
func bfs(valves map[string]*Valve, start, end string) int {
	if start == end {
		return 0
	}

	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	dist := make(map[string]int)
	dist[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return dist[end]
		}

		for _, next := range valves[current].tunnels {
			if !visited[next] {
				visited[next] = true
				dist[next] = dist[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return 1000 // Unreachable
}

func day16Part1(valves map[string]*Valve) uint {
	// Build list of valves with non-zero flow
	var important []string
	valveIdx := make(map[string]int)
	idx := 0

	for name, valve := range valves {
		if valve.flowRate > 0 {
			important = append(important, name)
			valveIdx[name] = idx
			idx++
		}
	}

	// Precompute distances between all important valves and from AA
	dist := make(map[string]map[string]int)
	allNodes := append([]string{"AA"}, important...)

	for _, from := range allNodes {
		dist[from] = make(map[string]int)
		for _, to := range important {
			if from != to {
				dist[from][to] = bfs(valves, from, to)
			}
		}
	}

	// DFS with memoization
	cache := make(map[string]uint)

	var dfs func(pos string, time int, opened uint64) uint
	dfs = func(pos string, time int, opened uint64) uint {
		// Create state key
		key := pos + "," + strconv.Itoa(time) + "," + strconv.FormatUint(opened, 10)
		if cached, ok := cache[key]; ok {
			return cached
		}

		best := uint(0)

		// Try opening each unopened valve
		for _, next := range important {
			bit := uint64(1) << valveIdx[next]
			if opened&bit == 0 { // Not opened yet
				travelTime := dist[pos][next]
				newTime := time + travelTime + 1 // Travel + open valve

				if newTime < 30 {
					remaining := 30 - newTime
					pressure := uint(valves[next].flowRate * remaining)
					newOpened := opened | bit

					result := pressure + dfs(next, newTime, newOpened)
					if result > best {
						best = result
					}
				}
			}
		}

		cache[key] = best
		return best
	}

	return dfs("AA", 0, 0)
}
