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
	MaxOre             int // Max ore needed per turn
	MaxClay            int // Max clay needed per turn
	MaxObsidian        int // Max obsidian needed per turn
}

type State struct {
	Minute             int
	Ore                int
	Clay               int
	Obsidian           int
	Geodes             int
	OreRobots          int
	ClayRobots         int
	ObsidianRobots     int
	GeodeRobots        int
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
	var dfs func(s State)

	dfs = func(s State) {
		if s.Minute == minutes {
			if s.Geodes > maxGeodes {
				maxGeodes = s.Geodes
			}
			return
		}

		// Pruning: if we built a geode robot every remaining minute, could we beat max?
		remaining := minutes - s.Minute
		maxPossible := s.Geodes + s.GeodeRobots*remaining + remaining*(remaining-1)/2
		if maxPossible <= maxGeodes {
			return
		}

		// Try building each robot type (or do nothing)
		canBuildGeodeRobot := s.Ore >= bp.GeodeRobotOre && s.Obsidian >= bp.GeodeRobotObsidian
		canBuildObsidianRobot := s.Ore >= bp.ObsidianRobotOre && s.Clay >= bp.ObsidianRobotClay
		canBuildClayRobot := s.Ore >= bp.ClayRobotOre
		canBuildOreRobot := s.Ore >= bp.OreRobotOre

		// Always build geode robot if possible
		if canBuildGeodeRobot {
			next := s
			next.Ore -= bp.GeodeRobotOre
			next.Obsidian -= bp.GeodeRobotObsidian
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.GeodeRobots++
			next.Minute++
			dfs(next)
			return
		}

		// Try building obsidian robot (if we need more obsidian production)
		if canBuildObsidianRobot && s.ObsidianRobots < bp.MaxObsidian {
			next := s
			next.Ore -= bp.ObsidianRobotOre
			next.Clay -= bp.ObsidianRobotClay
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.ObsidianRobots++
			next.Minute++
			dfs(next)
		}

		// Try building clay robot (if we need more clay production)
		if canBuildClayRobot && s.ClayRobots < bp.MaxClay {
			next := s
			next.Ore -= bp.ClayRobotOre
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.ClayRobots++
			next.Minute++
			dfs(next)
		}

		// Try building ore robot (if we need more ore production)
		if canBuildOreRobot && s.OreRobots < bp.MaxOre {
			next := s
			next.Ore -= bp.OreRobotOre
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.OreRobots++
			next.Minute++
			dfs(next)
		}

		// Try doing nothing (just collect resources)
		{
			next := s
			next.Ore += s.OreRobots
			next.Clay += s.ClayRobots
			next.Obsidian += s.ObsidianRobots
			next.Geodes += s.GeodeRobots
			next.Minute++
			dfs(next)
		}
	}

	dfs(initialState)
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
