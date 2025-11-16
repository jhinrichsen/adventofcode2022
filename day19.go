package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

func Day19(lines []string, part1 bool) uint {
	var bps []Blueprint
	for _, line := range lines {
		bp := NewBlueprint(line)
		bps = append(bps, bp)
	}

	if len(bps) > 0 {
		bps[0].Step(24)
	}

	sum := 0
	for i := range bps {
		sum += bps[i].QualityLevel()
	}
	return uint(sum)
}

func NewBlueprint(line string) Blueprint {
	var bp Blueprint
	line = strings.Replace(line, ":", " ", 1)
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
	bp.ID = num(1)
	bp.OreRobot = num(6)
	bp.ClayRobot = num(12)
	bp.ObsidianRobotOre = num(18)
	bp.ObsidianRobotClay = num(21)
	bp.GeodeRobotOre = num(27)
	bp.GeodeRobotObsidian = num(30)

	bp.OreRobots = 1
	bp.ClayRobots = 0
	bp.ObsidianRobots = 0
	bp.GeodeRobots = 0

	return bp
}

type Blueprint struct {
	ID int

	OreRobot           int
	ClayRobot          int
	ObsidianRobotOre   int
	ObsidianRobotClay  int
	GeodeRobotOre      int
	GeodeRobotObsidian int

	OreRobots      int
	ClayRobots     int
	ObsidianRobots int
	GeodeRobots    int

	Ore       int
	Clay      int
	Obsidians int
	Geodes    int
}

func (a Blueprint) QualityLevel() int {
	return a.ID * a.Geodes
}

func (a Blueprint) Step(n int) {
	/*
		   geodeR := 3*ore + 12*obsidian
		   obsidianR := 3*ore + 8*clay
		   clayR := 3 * ore
		   oreR := 2

		   // =>

		   geodeR := 3*ore + 12*(3*ore+8*(3*ore))

		Each ore robot costs 2 ore.
		Each clay robot costs 3 ore.
		Each obsidian robot costs 3 ore and 8 clay.
		Each geode robot costs 3 ore and 12 obsidian.

		=>

		geode := ore : obsidian = 3 : 12
		obsidian := ore : clay = 3 : 8
		(geode := 9 : 96 = 1 : 3)
		clay := clay : ore = 2 : 1

	*/
	for i := 0; i < n; i++ {
		fmt.Printf("== Minute %d ==\n", i+1)

		// buy robots

		// produce
		a.Ore += a.OreRobots * 1
		fmt.Printf("%d ore-collecting robot collects %d ore; "+
			"you now have %d ore\n",
			a.OreRobots, a.OreRobots, a.Ore)

		// robots finished
	}
}
