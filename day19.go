package adventofcode2022

import (
	"strconv"
	"strings"
)

type Blueprint struct {
	ID                 int
	OreRobotOre        int
	ClayRobotOre       int
	ObsidianRobotOre   int
	ObsidianRobotClay  int
	GeodeRobotOre      int
	GeodeRobotObsidian int
	MaxOre             int
	MaxClay            int
	MaxObsidian        int
}

type State struct {
	Minute         uint8
	Ore            uint8
	Clay           uint8
	Obsidian       uint8
	Geodes         uint8
	OreRobots      uint8
	ClayRobots     uint8
	ObsidianRobots uint8
	GeodeRobots    uint8
}

func NewBlueprint(line string) Blueprint {
	line = strings.ReplaceAll(line, ":", "")
	line = strings.ReplaceAll(line, ".", "")
	fs := strings.Fields(line)

	num := func(idx int) int {
		if idx >= len(fs) {
			return 0
		}
		n, err := strconv.Atoi(fs[idx])
		if err != nil {
			return 0
		}
		return n
	}

	bp := Blueprint{
		ID:                 num(1),
		OreRobotOre:        num(6),
		ClayRobotOre:       num(12),
		ObsidianRobotOre:   num(18),
		ObsidianRobotClay:  num(21),
		GeodeRobotOre:      num(27),
		GeodeRobotObsidian: num(30),
	}

	bp.MaxOre = max(bp.OreRobotOre, bp.ClayRobotOre, bp.ObsidianRobotOre, bp.GeodeRobotOre)
	bp.MaxClay = bp.ObsidianRobotClay
	bp.MaxObsidian = bp.GeodeRobotObsidian

	return bp
}

func (bp Blueprint) maxGeodes(minutes int) int {
	initialState := State{
		Minute:    0,
		OreRobots: 1,
	}

	maxGeodes := 0
	stack := []State{initialState}

	// Use smaller map with more aggressive pruning
	seen := make(map[State]struct{}, 100000)

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if int(s.Minute) == minutes {
			if int(s.Geodes) > maxGeodes {
				maxGeodes = int(s.Geodes)
			}
			continue
		}

		// Pruning: if we built a geode robot every remaining minute, could we beat max?
		remaining := minutes - int(s.Minute)
		maxPossible := int(s.Geodes) + int(s.GeodeRobots)*remaining + remaining*(remaining-1)/2
		if maxPossible <= maxGeodes {
			continue
		}

		// Aggressive resource capping
		maxOreNeeded := uint8(min(255, bp.MaxOre*remaining))
		maxClayNeeded := uint8(min(255, bp.MaxClay*remaining))
		maxObsidianNeeded := uint8(min(255, bp.MaxObsidian*remaining))

		if s.Ore > maxOreNeeded {
			s.Ore = maxOreNeeded
		}
		if s.Clay > maxClayNeeded {
			s.Clay = maxClayNeeded
		}
		if s.Obsidian > maxObsidianNeeded {
			s.Obsidian = maxObsidianNeeded
		}

		// Check if we've seen this state
		if _, exists := seen[s]; exists {
			continue
		}
		seen[s] = struct{}{}

		canBuildGeodeRobot := int(s.Ore) >= bp.GeodeRobotOre && int(s.Obsidian) >= bp.GeodeRobotObsidian
		canBuildObsidianRobot := int(s.Ore) >= bp.ObsidianRobotOre && int(s.Clay) >= bp.ObsidianRobotClay
		canBuildClayRobot := int(s.Ore) >= bp.ClayRobotOre
		canBuildOreRobot := int(s.Ore) >= bp.OreRobotOre

		// Always build geode robot if possible (greedy)
		if canBuildGeodeRobot {
			next := s
			next.Ore -= uint8(bp.GeodeRobotOre)
			next.Obsidian -= uint8(bp.GeodeRobotObsidian)
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.GeodeRobots++
			next.Minute++
			stack = append(stack, next)
			continue // Don't explore other options if we can build geode robot
		}

		// Try building obsidian robot (if we need more obsidian production)
		if canBuildObsidianRobot && s.ObsidianRobots < uint8(bp.MaxObsidian) {
			next := s
			next.Ore -= uint8(bp.ObsidianRobotOre)
			next.Clay -= uint8(bp.ObsidianRobotClay)
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.ObsidianRobots++
			next.Minute++
			stack = append(stack, next)
		}

		// Try building clay robot (if we need more clay production)
		if canBuildClayRobot && s.ClayRobots < uint8(bp.MaxClay) {
			next := s
			next.Ore -= uint8(bp.ClayRobotOre)
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.ClayRobots++
			next.Minute++
			stack = append(stack, next)
		}

		// Try building ore robot (if we need more ore production)
		if canBuildOreRobot && s.OreRobots < uint8(bp.MaxOre) {
			next := s
			next.Ore -= uint8(bp.OreRobotOre)
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.OreRobots++
			next.Minute++
			stack = append(stack, next)
		}

		// Try doing nothing (wait)
		{
			next := s
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.Minute++
			stack = append(stack, next)
		}
	}

	return maxGeodes
}

func (bp Blueprint) QualityLevel(minutes int) int {
	return bp.ID * bp.maxGeodes(minutes)
}

func Day19(lines []string, part1 bool) uint {
	var bps []Blueprint
	for _, line := range lines {
		if line == "" {
			continue
		}
		bp := NewBlueprint(line)
		bps = append(bps, bp)
	}

	if part1 {
		sum := 0
		for i := range bps {
			sum += bps[i].QualityLevel(24)
		}
		return uint(sum)
	}

	// Part 2: first 3 blueprints, 32 minutes, multiply results
	product := 1
	for i := range min(3, len(bps)) {
		product *= bps[i].maxGeodes(32)
	}
	return uint(product)
}
